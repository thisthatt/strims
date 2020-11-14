package network

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"path"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/dao"
	"github.com/MemeLabs/go-ppspp/pkg/event"
	"github.com/MemeLabs/go-ppspp/pkg/kv"
	"github.com/MemeLabs/go-ppspp/pkg/logutil"
	"github.com/MemeLabs/go-ppspp/pkg/pb"
	"github.com/MemeLabs/go-ppspp/pkg/protoutil"
	"github.com/MemeLabs/go-ppspp/pkg/vnic"
	"github.com/MemeLabs/go-ppspp/pkg/vpn"
	"go.uber.org/zap"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// errors ...
var (
	ErrDuplicateNetworkKey      = errors.New("duplicate network key")
	ErrNetworkNotFound          = errors.New("network not found")
	ErrPeerClosed               = errors.New("peer closed")
	ErrNetworkBindingsEmpty     = errors.New("network bindings empty")
	ErrDiscriminatorBounds      = errors.New("discriminator out of range")
	ErrNetworkOwnerMismatch     = errors.New("init and network certificate key mismatch")
	ErrNetworkAuthorityMismatch = errors.New("network ca mismatch")
	ErrNetworkIDBounds          = errors.New("network id out of range")
)

// Broker negotiates common networks with peers.
type Broker interface {
	SendKeys(c ReadWriteFlusher, keys [][]byte) error
	ReceiveKeys(c ReadWriteFlusher, keys [][]byte) ([][]byte, error)
}

// ReadWriteFlusher ...
type ReadWriteFlusher interface {
	io.ReadWriter
	Flush() error
}

// NewPeerHandler ...
func NewPeerHandler(logger *zap.Logger, broker Broker, host *vpn.Host, store *dao.ProfileStore, profile *pb.Profile) *PeerHandler {
	h := &PeerHandler{
		logger:  logger,
		broker:  broker,
		host:    host,
		store:   store,
		profile: profile,
	}

	host.VNIC().AddPeerHandler(h.handlePeer)

	go h.updateCertificates()

	return h
}

// PeerHandler ...
type PeerHandler struct {
	logger       *zap.Logger
	broker       Broker
	host         *vpn.Host
	store        *dao.ProfileStore
	profile      *pb.Profile
	certificates event.Observable
}

func (h *PeerHandler) handlePeer(p *vnic.Peer) {
	b := &bootstrap{
		logger:       h.logger,
		broker:       h.broker,
		host:         h.host,
		profile:      h.profile,
		peer:         p,
		links:        make(map[*vpn.Network]*bootstrapLink),
		handshakes:   make(chan *pb.NetworkHandshake),
		errors:       make(chan error),
		peerInit:     make(chan *pb.NetworkHandshake_Init, 1),
		bindings:     make(chan *pb.NetworkHandshake_NetworkBindings),
		certificates: &h.certificates,
		conn:         p.Channel(vnic.NetworkInitPort),
		brokerConn:   p.Channel(vnic.NetworkBrokerPort),
	}

	go b.run()
}

func (h *PeerHandler) updateCertificates() {
	certificates := make(chan *pb.Certificate, 1)
	h.certificates.Notify(certificates)

	for cert := range certificates {
		err := h.store.Update(func(tx kv.RWTx) error {
			network, err := dao.GetNetworkByKey(tx, dao.GetRootCert(cert).Key)
			if err != nil {
				return err
			}
			network.Certificate = cert
			return dao.UpsertNetwork(tx, network)
		})
		if err != nil {
			h.logger.Debug("error saving certificate", zap.Error(err))
		}
	}
}

type bootstrap struct {
	logger       *zap.Logger
	broker       Broker
	host         *vpn.Host
	peer         *vnic.Peer
	profile      *pb.Profile
	links        map[*vpn.Network]*bootstrapLink
	syncing      int32
	handshakes   chan *pb.NetworkHandshake
	errors       chan error
	peerInit     chan *pb.NetworkHandshake_Init
	bindings     chan *pb.NetworkHandshake_NetworkBindings
	certificates *event.Observable
	conn         *vnic.FrameReadWriter
	brokerConn   *vnic.FrameReadWriter
}

func (h *bootstrap) removeNetworkLinks() {
	for n := range h.links {
		h.logger.Debug(
			"removing peer from network",
			zap.Stringer("peer", h.peer.HostID()),
			logutil.ByteHex("network", n.Key()),
		)

		n.RemovePeer(h.peer.HostID())
	}
}

func (h *bootstrap) clientAndLinkForNetworkKey(networkKey []byte) (*vpn.Client, *bootstrapLink, error) {
	client, ok := h.host.Client(networkKey)
	if !ok {
		return nil, nil, ErrNetworkNotFound
	}

	return client, h.links[client.Network], nil
}

func (h *bootstrap) run() {
	networks := make(chan *vpn.Network, 1)
	h.host.NotifyNetwork(networks)

	certificates := make(chan *pb.Certificate, 1)
	h.certificates.Notify(certificates)

	defer func() {
		h.host.StopNotifyingNetwork(networks)
		h.certificates.StopNotifying(certificates)

		h.removeNetworkLinks()
		h.peer.Close()
	}()

	go h.sync()
	go h.readHandshakes()

	for {
		select {
		case <-networks:
			go h.sync()
		case cert := <-certificates:
			h.sendCertificateUpdate(cert)
		case handshake := <-h.handshakes:
			h.handleHandshake(handshake)
		case err := <-h.errors:
			h.logger.Error("failed to read handshake", zap.Error(err))
			return
		}
	}
}

func (h *bootstrap) readHandshakes() {
	for {
		handshake := &pb.NetworkHandshake{}
		if err := protoutil.ReadStream(h.conn, handshake); err != nil {
			h.errors <- err
			return
		}
		h.handshakes <- handshake
	}
}

func (h *bootstrap) handleHandshake(handshake *pb.NetworkHandshake) error {
	switch body := handshake.Body.(type) {
	case *pb.NetworkHandshake_Init_:
		h.peerInit <- body.Init
		go h.sync()
	case *pb.NetworkHandshake_NetworkBindings_:
		h.bindings <- body.NetworkBindings
	case *pb.NetworkHandshake_CertificateUpgradeOffer_:
		return h.handleCertificateUpgradeOffer(body.CertificateUpgradeOffer)
	case *pb.NetworkHandshake_CertificateUpgradeRequest_:
		return h.handleCertificateUpgradeRequest(body.CertificateUpgradeRequest)
	case *pb.NetworkHandshake_CertificateUpgradeResponse_:
		return h.handleCertificateUpgradeResponse(body.CertificateUpgradeResponse)
	case *pb.NetworkHandshake_CertificateUpdate_:
		return h.handleCertificateUpdate(body.CertificateUpdate)
	}
	return nil
}

func (h *bootstrap) handleCertificateUpgradeOffer(req *pb.NetworkHandshake_CertificateUpgradeOffer) error {
	csr, err := dao.NewCertificateRequest(
		h.profile.Key,
		pb.KeyUsage_KEY_USAGE_PEER|pb.KeyUsage_KEY_USAGE_SIGN,
		dao.WithSubject(h.profile.Name),
	)
	if err != nil {
		return err
	}

	client, _, err := h.clientAndLinkForNetworkKey(req.NetworkKey)
	if err != nil {
		return err
	}

	// TODO: do we need to upgrade this cert? (invite/expired)
	// TODO: are we already upgrading this cert via some other peer?

	return h.send(&pb.NetworkHandshake{
		Body: &pb.NetworkHandshake_CertificateUpgradeRequest_{
			CertificateUpgradeRequest: &pb.NetworkHandshake_CertificateUpgradeRequest{
				Certificate:        client.Network.Certificate(),
				CertificateRequest: csr,
			},
		},
	})
}

func (h *bootstrap) handleCertificateUpgradeRequest(req *pb.NetworkHandshake_CertificateUpgradeRequest) error {
	networkKey := dao.GetRootCert(req.Certificate).Key

	renewReq := &pb.CARenewRequest{
		Certificate:        req.Certificate,
		CertificateRequest: req.CertificateRequest,
	}
	renewRes := &pb.CARenewResponse{}
	err := func() error {
		client, _, err := h.clientAndLinkForNetworkKey(networkKey)
		if err != nil {
			return err
		}

		caClient, err := NewCAClient(h.logger, client)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return caClient.Renew(ctx, renewReq, renewRes)
	}()

	res := &pb.NetworkHandshake_CertificateUpgradeResponse{
		NetworkKey: networkKey,
	}
	if err != nil {
		res.Body = &pb.NetworkHandshake_CertificateUpgradeResponse_Error{
			Error: err.Error(),
		}
	} else {
		res.Body = &pb.NetworkHandshake_CertificateUpgradeResponse_Certificate{
			Certificate: renewRes.Certificate,
		}
	}

	return h.send(&pb.NetworkHandshake{
		Body: &pb.NetworkHandshake_CertificateUpgradeResponse_{
			CertificateUpgradeResponse: res,
		},
	})
}

func (h *bootstrap) handleCertificateUpgradeResponse(req *pb.NetworkHandshake_CertificateUpgradeResponse) error {
	if err := req.GetError(); err != "" {
		// TODO: propagate to user
		return errors.New(err)
	}

	cert := req.GetCertificate()
	if err := dao.VerifyCertificate(cert); err != nil {
		return err
	}

	// TODO: confirm that we want to replace the cert

	h.certificates.Emit(cert)
	return nil
}

func (h *bootstrap) sendCertificateUpdate(cert *pb.Certificate) error {
	if _, link, _ := h.clientAndLinkForNetworkKey(dao.GetRootCert(cert).Key); link == nil {
		return nil
	}

	return h.send(&pb.NetworkHandshake{
		Body: &pb.NetworkHandshake_CertificateUpdate_{
			CertificateUpdate: &pb.NetworkHandshake_CertificateUpdate{
				Certificate: cert,
			},
		},
	})
}

func (h *bootstrap) handleCertificateUpdate(req *pb.NetworkHandshake_CertificateUpdate) error {
	if err := dao.VerifyCertificate(req.Certificate); err != nil {
		return err
	}

	networkKey := dao.GetRootCert(req.Certificate).Key

	client, link, err := h.clientAndLinkForNetworkKey(networkKey)
	if err != nil {
		return err
	}
	if link == nil {
		return errors.New("peer network binding not negotiated")
	}
	if link.open {
		return errors.New("peer already added to network")
	}

	// handle invite certs
	// TODO: also handle expired certificates
	if !bytes.Equal(networkKey, req.Certificate.GetParent().Key) {
		return errors.New("invalid peer cert")
	}

	h.logger.Debug(
		"adding peer to network",
		zap.Stringer("peer", h.peer.HostID()),
		logutil.ByteHex("network", networkKey),
		zap.Uint16("localPort", link.localPort),
		zap.Uint16("peerPort", link.peerPort),
	)

	client.Network.AddPeer(h.peer, link.localPort, link.peerPort)
	link.open = true

	return nil
}

func (h *bootstrap) send(msg protoreflect.ProtoMessage) error {
	err := protoutil.WriteStream(h.conn, msg)
	if err != nil {
		return err
	}
	return h.conn.Flush()
}

func (h *bootstrap) sync() {
	if !atomic.CompareAndSwapInt32(&h.syncing, 0, 1) {
		return
	}

	keys := h.host.NetworkKeys()
	err := h.send(&pb.NetworkHandshake{
		Body: &pb.NetworkHandshake_Init_{
			Init: &pb.NetworkHandshake_Init{
				KeyCount: int32(len(keys)),
			},
		},
	})
	if err != nil {
		atomic.StoreInt32(&h.syncing, 0)
		h.errors <- fmt.Errorf("sync failed: %w", err)
		return
	}

	if err := h.exchangeBindings(keys); err != nil {
		h.errors <- fmt.Errorf("sync failed: %w", err)
	}
	atomic.StoreInt32(&h.syncing, 0)
}

func (h *bootstrap) exchangeBindings(keys [][]byte) error {
	select {
	case <-time.After(time.Second * 10):
		return errors.New("timeout")
	case peerInit := <-h.peerInit:
		if len(keys) == 0 || peerInit.KeyCount == 0 {
			return nil
		}

		preferSend := h.peer.HostID().Less(h.host.VNIC().ID())
		if len(keys) > int(peerInit.KeyCount) || (len(keys) == int(peerInit.KeyCount) && preferSend) {
			return h.exchangeBindingsAsSender(keys)
		} else {
			return h.exchangeBindingsAsReceiver(keys)
		}
	}
}

func (h *bootstrap) exchangeBindingsAsReceiver(keys [][]byte) error {
	keys, err := h.broker.ReceiveKeys(h.brokerConn, keys)
	if err != nil {
		return err
	}
	networkBindings, err := h.sendNetworkBindings(keys)
	if err != nil {
		return err
	}
	peerNetworkBindings := <-h.bindings
	if _, err = h.verifyNetworkBindings(peerNetworkBindings); err != nil {
		return err
	}
	return h.handleNetworkBindings(networkBindings, peerNetworkBindings.NetworkBindings)
}

func (h *bootstrap) exchangeBindingsAsSender(keys [][]byte) error {
	if err := h.broker.SendKeys(h.brokerConn, keys); err != nil {
		return err
	}
	peerNetworkBindings := <-h.bindings
	keys, err := h.verifyNetworkBindings(peerNetworkBindings)
	if err != nil {
		return err
	}
	networkBindings, err := h.sendNetworkBindings(keys)
	if err != nil {
		return err
	}
	return h.handleNetworkBindings(networkBindings, peerNetworkBindings.NetworkBindings)
}

func (h *bootstrap) sendNetworkBindings(keys [][]byte) ([]*pb.NetworkHandshake_NetworkBinding, error) {
	var bindings []*pb.NetworkHandshake_NetworkBinding

	for _, key := range keys {
		client, link, err := h.clientAndLinkForNetworkKey(key)
		if err != nil {
			return nil, err
		}
		if link != nil {
			continue
		}

		port, err := h.peer.ReservePort()
		if err != nil {
			return nil, err
		}

		bindings = append(
			bindings,
			&pb.NetworkHandshake_NetworkBinding{
				Port:        uint32(port),
				Certificate: client.Network.Certificate(),
			},
		)
	}
	err := protoutil.WriteStream(h.conn, &pb.NetworkHandshake{
		Body: &pb.NetworkHandshake_NetworkBindings_{
			NetworkBindings: &pb.NetworkHandshake_NetworkBindings{
				NetworkBindings: bindings,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if err := h.conn.Flush(); err != nil {
		return nil, err
	}
	return bindings, nil
}

func (h *bootstrap) verifyNetworkBindings(bindings *pb.NetworkHandshake_NetworkBindings) ([][]byte, error) {
	if bindings == nil {
		return nil, ErrNetworkBindingsEmpty
	}

	keys := make([][]byte, len(bindings.NetworkBindings))
	for i, b := range bindings.NetworkBindings {
		if err := dao.VerifyCertificate(b.Certificate); err != nil {
			return nil, err
		}
		keys[i] = dao.GetRootCert(b.Certificate).Key
	}
	return keys, nil
}

func (h *bootstrap) handleNetworkBindings(networkBindings, peerNetworkBindings []*pb.NetworkHandshake_NetworkBinding) error {
	for i, peerBinding := range peerNetworkBindings {
		binding := networkBindings[i]
		networkKey := dao.GetRootCert(peerBinding.Certificate).Key

		if !bytes.Equal(h.peer.Certificate.Key, peerBinding.Certificate.Key) {
			return ErrNetworkOwnerMismatch
		}
		if !bytes.Equal(dao.GetRootCert(binding.Certificate).Key, networkKey) {
			return ErrNetworkAuthorityMismatch
		}
		if peerBinding.Port > uint32(math.MaxUint16) {
			return ErrNetworkIDBounds
		}

		client, _, err := h.clientAndLinkForNetworkKey(networkKey)
		if err != nil {
			return err
		}

		link := &bootstrapLink{
			localPort: uint16(binding.Port),
			peerPort:  uint16(peerBinding.Port),
		}
		h.links[client.Network] = link

		// handle invite certs
		// TODO: also handle expired certificates
		if !bytes.Equal(networkKey, peerBinding.Certificate.GetParent().Key) {
			// jsonDump(peerBinding)
			err := h.send(&pb.NetworkHandshake{
				Body: &pb.NetworkHandshake_CertificateUpgradeOffer_{
					CertificateUpgradeOffer: &pb.NetworkHandshake_CertificateUpgradeOffer{
						NetworkKey: networkKey,
					},
				},
			})
			if err != nil {
				h.logger.Debug("cert upgrade offer failed", zap.Error(err))
			}
			continue
		}

		h.logger.Debug(
			"adding peer to network",
			zap.Stringer("peer", h.peer.HostID()),
			logutil.ByteHex("network", networkKey),
			zap.Uint16("localPort", link.localPort),
			zap.Uint16("peerPort", link.peerPort),
		)

		client.Network.AddPeer(h.peer, link.localPort, link.peerPort)
		link.open = true

		// h.host.peerNetworkObservers.Emit(PeerNetwork{h.peer, c.Network})
	}
	return nil
}

type bootstrapLink struct {
	localPort uint16
	peerPort  uint16
	open      bool
}

func jsonDump(i interface{}) {
	_, file, line, _ := runtime.Caller(1)
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf(
		"%s %s:%d: %s\n",
		time.Now().Format("2006/01/02 15:04:05.000000"),
		path.Base(file),
		line, string(b),
	)
}
