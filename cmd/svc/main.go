package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"

	"github.com/MemeLabs/go-ppspp/internal/dao"
	"github.com/MemeLabs/go-ppspp/internal/frontend"
	"github.com/MemeLabs/go-ppspp/internal/network"
	"github.com/MemeLabs/go-ppspp/pkg/apis/type/key"
	"github.com/MemeLabs/go-ppspp/pkg/kv/bbolt"
	"github.com/MemeLabs/go-ppspp/pkg/vnic"
	"github.com/MemeLabs/go-ppspp/pkg/vpn"
	"github.com/MemeLabs/protobuf/pkg/rpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

var (
	profileDir    string
	addr          string
	metricsAddr   string
	rtmpAddr      string
	debugAddr     string
	webRTCPortMin uint
	webRTCPortMax uint
	vnicLabel     string
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("failed to locate home directory: %s", err)
	}

	flag.StringVar(&profileDir, "profile-dir", homeDir, "profile db location")
	flag.StringVar(&addr, "addr", ":8082", "bootstrap server listen address")
	flag.StringVar(&metricsAddr, "metrics-addr", ":1971", "metrics server listen address")
	flag.StringVar(&rtmpAddr, "rtmp-addr", ":1935", "rtmp server listen address")
	flag.StringVar(&debugAddr, "debug-addr", ":6060", "debug server listen address")
	flag.UintVar(&webRTCPortMin, "webrtc-port-min", 0, "webrtc ephemeral port range min")
	flag.UintVar(&webRTCPortMax, "webrtc-port-max", 0, "webrtc ephemeral port range max")
	flag.StringVar(&vnicLabel, "vnic-label", "", "vnic label")
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

	if debugAddr != "" {
		go func() {
			log.Println(http.ListenAndServe(debugAddr, nil))
		}()
	}

	if metricsAddr != "" {
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.ListenAndServe(metricsAddr, nil)
		}()
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	// t, err := newManager(logger)
	// if err != nil {
	// 	panic(err)
	// }
	// t.Run()
	// return

	store, err := initProfileStore()
	if err != nil {
		panic(err)
	}

	profile, err := dao.GetProfile(store)
	if err != nil {
		panic(err)
	}

	server := rpc.NewServer(logger, &noopDialer{})

	newVPN := func(key *key.Key) (*vpn.Host, error) {
		vnicHost, err := vnic.New(
			logger,
			key,
			vnic.WithLabel(vnicLabel),
			vnic.WithInterface(vnic.NewWSInterface(logger, addr)),
			vnic.WithInterface(vnic.NewWebRTCInterface(vnic.NewWebRTCDialer(
				logger,
				&vnic.WebRTCDialerOptions{
					PortMin: uint16(webRTCPortMin),
					PortMax: uint16(webRTCPortMax),
				},
			))),
		)
		if err != nil {
			return nil, err
		}
		return vpn.New(logger, vnicHost)
	}

	c := frontend.New(logger, server, newVPN, network.NewBroker(logger))
	if err := c.Init(context.Background(), profile, store); err != nil {
		logger.Fatal("frontend instance init failed", zap.Error(err))
	}

	select {}
}

func initProfileStore() (*dao.ProfileStore, error) {
	kv, err := bbolt.NewStore(path.Join(profileDir, ".strims"))
	if err != nil {
		log.Fatalf("failed to open db: %s", err)
	}
	ds, err := dao.NewMetadataStore(kv)
	if err != nil {
		panic(err)
	}

	profiles, err := dao.GetProfileSummaries(ds)
	if err != nil {
		return nil, err
	}

	name := "test"
	pw := "test"

	var found bool
	for _, p := range profiles {
		found = found || p.Name == name
	}
	if !found {
		_, profileStore, err := dao.CreateProfile(ds, name, pw)
		return profileStore, err
	}

	_, profileStore, err := dao.LoadProfile(ds, profiles[0].Id, pw)
	return profileStore, err
}

type noopDialer struct{}

func (d *noopDialer) Dial(context.Context, rpc.Dispatcher) (rpc.Transport, error) {
	return &noopTransport{}, nil
}

type noopTransport struct{}

func (t *noopTransport) Call(*rpc.CallOut, rpc.ResponseFunc) error {
	return nil
}

func (t *noopTransport) Listen() error {
	return nil
}
