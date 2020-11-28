package bootstrap

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/control/event"
	"github.com/MemeLabs/go-ppspp/pkg/dao"
	"github.com/MemeLabs/go-ppspp/pkg/pb"
	"github.com/MemeLabs/go-ppspp/pkg/vnic"
	"github.com/MemeLabs/go-ppspp/pkg/vpn"
	"go.uber.org/zap"
)

// NewControl ...
func NewControl(logger *zap.Logger, vpn *vpn.Host, store *dao.ProfileStore, observers *event.Observers) *Control {
	events := make(chan interface{}, 128)
	observers.Peer.Notify(events)
	observers.Network.Notify(events)

	return &Control{
		logger:    logger,
		vpn:       vpn,
		store:     store,
		observers: observers,
		events:    events,
		peers:     map[uint64]*Peer{},
	}
}

// Control ...
type Control struct {
	logger *zap.Logger
	vpn    *vpn.Host
	store  *dao.ProfileStore

	lock              sync.Mutex
	observers         *event.Observers
	events            chan interface{}
	certRenewTimeout  <-chan time.Time
	lastCertRenewTime time.Time
	nextID            uint64
	peers             map[uint64]*Peer
}

// Run ...
func (t *Control) Run(ctx context.Context) {
	go t.startClients()

	for {
		select {
		case e := <-t.events:
			switch e := e.(type) {
			case event.PeerAdd:
				t.handlePeerAdd(ctx, e.ID)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (t *Control) handlePeerAdd(ctx context.Context, id uint64) {
	peer, ok := t.peers[id]
	if !ok {
		return
	}

	var res pb.BootstrapPeerGetPublishEnabledResponse
	if err := peer.client.Bootstrap().GetPublishEnabled(ctx, &pb.BootstrapPeerGetPublishEnabledRequest{}, &res); err != nil {
		t.logger.Debug("bootstrap publish enabled check failed", zap.Error(err))
	}

	// TODO: locking...
	peer.PublishEnabled = res.Enabled
}

// AddPeer ...
func (t *Control) AddPeer(id uint64, peer *vnic.Peer, client PeerClient) *Peer {
	p := &Peer{
		vnic:   peer,
		client: client,
	}

	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[id] = p

	return p
}

// RemovePeer ...
func (t *Control) RemovePeer(id uint64) {
	t.lock.Lock()
	defer t.lock.Unlock()

	delete(t.peers, id)
}

func (t *Control) startClients() {
	clients, err := dao.GetBootstrapClients(t.store)
	if err != nil {
		t.logger.Fatal(
			"loading bootstrap clients failed",
			zap.Error(err),
		)
	}

	for _, client := range clients {
		go func(client *pb.BootstrapClient) {
			if err := t.startClient(client); err != nil {
				t.logger.Debug("starting bootstrap client failed", zap.Error(err))
			}
		}(client)
	}
}

func (t *Control) startClient(client *pb.BootstrapClient) error {
	switch client := client.ClientOptions.(type) {
	case *pb.BootstrapClient_WebsocketOptions:
		return t.vpn.VNIC().Dial(vnic.WebSocketAddr{
			URL:                   client.WebsocketOptions.Url,
			InsecureSkipVerifyTLS: client.WebsocketOptions.InsecureSkipVerifyTls,
		})
	}
	return nil
}

// PublishingEnabled ...
func (t *Control) PublishingEnabled() bool {
	return true
}

// Publish ...
func (t *Control) Publish(ctx context.Context, peerID uint64, network *pb.Network, validDuration time.Duration) error {
	peer, ok := t.peers[peerID]
	if !ok {
		return errors.New("peer id not found")
	}

	if !peer.PublishEnabled {
		return errors.New("peer does not support network bootstrapping")
	}

	networkCert, err := dao.NewNetworkCertificate(network)
	if err != nil {
		return err
	}
	csr := &pb.CertificateRequest{
		Key:      peer.vnic.Certificate.Key,
		KeyType:  peer.vnic.Certificate.KeyType,
		KeyUsage: uint32(pb.KeyUsage_KEY_USAGE_BROKER),
	}
	cert, err := dao.SignCertificateRequest(csr, validDuration, network.Key)
	if err != nil {
		return err
	}
	cert.ParentOneof = &pb.Certificate_Parent{Parent: networkCert}

	return peer.client.Bootstrap().Publish(ctx, &pb.BootstrapPeerPublishRequest{Certificate: cert}, &pb.BootstrapPeerPublishResponse{})
}
