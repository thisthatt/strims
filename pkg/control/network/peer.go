package network

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MemeLabs/go-ppspp/pkg/api"
	"github.com/MemeLabs/go-ppspp/pkg/control/event"
	"github.com/MemeLabs/go-ppspp/pkg/dao"
	"github.com/MemeLabs/go-ppspp/pkg/logutil"
	"github.com/MemeLabs/go-ppspp/pkg/pb"
	"github.com/MemeLabs/go-ppspp/pkg/vnic"
	"github.com/MemeLabs/go-ppspp/pkg/vpn"
	"github.com/petar/GoLLRB/llrb"
	"go.uber.org/zap"
)

// PeerClient ..
type PeerClient interface {
	Bootstrap() *api.BootstrapPeerClient
	Network() *api.NetworkPeerClient
	CA() *api.CAPeerClient
}

// NewPeer ...
func NewPeer(
	id uint64,
	peer *vnic.Peer,
	client PeerClient,
	logger *zap.Logger,
	observers *event.Observers,
	broker Broker,
	vpn *vpn.Host,
	certificates *certificateMap,
) *Peer {
	return &Peer{
		id:           id,
		client:       client,
		peer:         peer,
		logger:       logger,
		observers:    observers,
		broker:       broker,
		vpn:          vpn,
		certificates: certificates,

		keyCount:   make(chan uint32, 1),
		bindings:   make(chan []*pb.NetworkPeerBinding, 1),
		brokerConn: peer.Channel(vnic.NetworkBrokerPort),
	}
}

// Peer ...
type Peer struct {
	id           uint64
	peer         *vnic.Peer
	client       PeerClient
	logger       *zap.Logger
	observers    *event.Observers
	broker       Broker
	vpn          *vpn.Host
	certificates *certificateMap

	lock        sync.Mutex
	links       llrb.LLRB
	negotiating uint32
	keyCount    chan uint32
	bindings    chan []*pb.NetworkPeerBinding
	brokerConn  *vnic.FrameReadWriter
}

// HandlePeerNegotiate ...
func (p *Peer) HandlePeerNegotiate(keyCount uint32) {
	select {
	case p.keyCount <- keyCount:
	default:
	}

	if atomic.LoadUint32(&p.negotiating) == 0 {
		go func() {
			if err := p.negotiateNetworks(context.Background()); err != nil {
				p.logger.Debug("network negotiation failed", zap.Error(err))
			}
		}()
	}
}

// HandlePeerOpen ...
func (p *Peer) HandlePeerOpen(bindings []*pb.NetworkPeerBinding) {
	select {
	case p.bindings <- bindings:
	default:
	}
}

// HandlePeerClose ...
func (p *Peer) HandlePeerClose(networkKey []byte) {
	p.closeNetworkWithoutNotifyingPeer(networkKey)
}

// HandlePeerUpdateCertificate ...
func (p *Peer) HandlePeerUpdateCertificate(cert *pb.Certificate) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := dao.VerifyCertificate(cert); err != nil {
		return err
	}
	if !isPeerCertificateOwner(p.peer, cert) {
		return ErrCertificateOwnerMismatch
	}
	if !isCertificateTrusted(cert) {
		return ErrProvisionalCertificate
	}

	li := p.links.Get(&networkBinding{networkKey: networkKeyForCertificate(cert)})
	if li == nil {
		return ErrNetworkBindingNotFound
	}

	link := li.(*networkBinding)
	link.peerCertTrusted = true

	return p.openNetwork(link)
}

func (p *Peer) sendCertificateUpdate(network *pb.Network) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	li := p.links.Get(&networkBinding{networkKey: networkKeyForCertificate(network.Certificate)})
	if li == nil {
		return ErrNetworkBindingNotFound
	}

	link := li.(*networkBinding)
	link.localCertTrusted = true

	err := p.client.Network().UpdateCertificate(
		context.Background(),
		&pb.NetworkPeerUpdateCertificateRequest{Certificate: network.Certificate},
		&pb.NetworkPeerUpdateCertificateResponse{},
	)
	if err != nil {
		return err
	}

	return p.openNetwork(link)
}

func (p *Peer) closeNetworkWithoutNotifyingPeer(networkKey []byte) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	li := p.links.Get(&networkBinding{networkKey: networkKey})
	if li == nil {
		return ErrNetworkBindingNotFound
	}
	p.links.Delete(li)

	client, ok := p.vpn.Client(networkKey)
	if !ok {
		return ErrNetworkNotFound
	}
	client.Network.RemovePeer(p.peer.HostID())

	p.observers.Network.Emit(event.NetworkPeerClose{
		PeerID:     p.id,
		NetworkID:  li.(*networkBinding).networkID,
		NetworkKey: networkKey,
	})

	p.logger.Info(
		"removed peer from network",
		zap.Stringer("peer", p.peer.HostID()),
		logutil.ByteHex("network", networkKey),
	)

	if p.links.Len() == 0 {
		p.peer.Close()
	}

	return nil
}

func (p *Peer) closeNetwork(networkKey []byte) {
	p.closeNetworkWithoutNotifyingPeer(networkKey)
	p.client.Network().Close(context.Background(), &pb.NetworkPeerCloseRequest{Key: networkKey}, &pb.NetworkPeerCloseResponse{})
}

func (p *Peer) networkKeysForLinks() [][]byte {
	p.lock.Lock()
	defer p.lock.Unlock()

	keys := make([][]byte, p.links.Len())
	p.links.AscendLessThan(llrb.Inf(1), func(li llrb.Item) bool {
		keys = append(keys, li.(*networkBinding).networkKey)
		return true
	})
	return keys
}

func (p *Peer) close() {
	for _, key := range p.networkKeysForLinks() {
		p.closeNetworkWithoutNotifyingPeer(key)
	}
}

func (p *Peer) hasNetworkBinding(networkKey []byte) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.links.Has(&networkBinding{networkKey: networkKey})
}

func (p *Peer) negotiateNetworks(ctx context.Context) error {
	if !atomic.CompareAndSwapUint32(&p.negotiating, 0, 1) {
		return errors.New("cannot begin new negotiation until previous negotiation finishes")
	}
	defer atomic.StoreUint32(&p.negotiating, 0)

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	keys := p.certificates.Keys()
	err := p.client.Network().Negotiate(ctx, &pb.NetworkPeerNegotiateRequest{KeyCount: uint32(len(keys))}, &pb.NetworkPeerNegotiateResponse{})
	if err != nil {
		return fmt.Errorf("sending network negotiation init failed: %w", err)
	}

	return p.exchangeBindings(ctx, keys)
}

func (p *Peer) exchangeBindings(ctx context.Context, keys [][]byte) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("no network negotiation init received: %w", ctx.Err())
	case keyCount := <-p.keyCount:
		if len(keys) == 0 || keyCount == 0 {
			return errors.New("one or both peers have zero keys")
		}

		// the psz sender role scales better than the receiver so by default we
		// pick role by comparing key counts. the role choice has to be symmetric
		// so host ids break ties.
		preferSend := p.peer.HostID().Less(p.vpn.VNIC().ID())
		if len(keys) > int(keyCount) || (len(keys) == int(keyCount) && preferSend) {
			return p.exchangeBindingsAsSender(ctx, keys)
		}
		return p.exchangeBindingsAsReceiver(ctx, keys)
	}
}

func (p *Peer) exchangeBindingsAsReceiver(ctx context.Context, keys [][]byte) error {
	keys, err := p.broker.ReceiveKeys(p.brokerConn, keys)
	if err != nil {
		return fmt.Errorf("network key broker failed: %w", err)
	}
	networkBindings, err := p.sendNetworkBindings(ctx, keys)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("peer network bindings not received: %w", ctx.Err())
	case peerNetworkBindings := <-p.bindings:
		if _, err = p.verifyNetworkBindings(peerNetworkBindings); err != nil {
			return err
		}

		p.observers.Network.Emit(event.NetworkPeerBindings{PeerID: p.id, NetworkKeys: keys})

		return p.handleNetworkBindings(networkBindings, peerNetworkBindings)
	}
}

func (p *Peer) exchangeBindingsAsSender(ctx context.Context, keys [][]byte) error {
	if err := p.broker.SendKeys(p.brokerConn, keys); err != nil {
		return fmt.Errorf("network key broker failed: %w", err)
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("peer network bindings not received: %w", ctx.Err())
	case peerNetworkBindings := <-p.bindings:
		keys, err := p.verifyNetworkBindings(peerNetworkBindings)
		if err != nil {
			return err
		}

		p.observers.Network.Emit(event.NetworkPeerBindings{PeerID: p.id, NetworkKeys: keys})

		networkBindings, err := p.sendNetworkBindings(ctx, keys)
		if err != nil {
			return err
		}

		return p.handleNetworkBindings(networkBindings, peerNetworkBindings)
	}
}

func (p *Peer) sendNetworkBindings(ctx context.Context, keys [][]byte) ([]*pb.NetworkPeerBinding, error) {
	var bindings []*pb.NetworkPeerBinding

	for _, key := range keys {
		c, ok := p.certificates.Get(key)
		if !ok {
			return nil, fmt.Errorf("certificate not found: %w", ErrNetworkNotFound)
		}

		if p.links.Has(&networkBinding{networkKey: key}) {
			continue
		}

		port, err := p.peer.ReservePort()
		if err != nil {
			return nil, err
		}

		bindings = append(
			bindings,
			&pb.NetworkPeerBinding{
				Port:        uint32(port),
				Certificate: c.certificate,
			},
		)
	}

	err := p.client.Network().Open(ctx, &pb.NetworkPeerOpenRequest{Bindings: bindings}, &pb.NetworkPeerOpenResponse{})
	if err != nil {
		return nil, err
	}
	return bindings, nil
}

func (p *Peer) verifyNetworkBindings(bindings []*pb.NetworkPeerBinding) ([][]byte, error) {
	if bindings == nil {
		return nil, ErrNetworkBindingsEmpty
	}

	keys := make([][]byte, len(bindings))
	for i, b := range bindings {
		if err, ok := dao.VerifyCertificate(b.Certificate).(dao.Errors); ok && !err.IncludesOnly(dao.ErrNotAfterRange) {
			return nil, err
		}
		keys[i] = networkKeyForCertificate(b.Certificate)
	}
	return keys, nil
}

func (p *Peer) handleNetworkBindings(networkBindings, peerNetworkBindings []*pb.NetworkPeerBinding) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	for i, peerBinding := range peerNetworkBindings {
		binding := networkBindings[i]
		networkKey := networkKeyForCertificate(peerBinding.Certificate)

		if !isPeerCertificateOwner(p.peer, peerBinding.Certificate) {
			return ErrCertificateOwnerMismatch
		}
		if !bytes.Equal(networkKeyForCertificate(binding.Certificate), networkKey) {
			return ErrNetworkAuthorityMismatch
		}
		if peerBinding.Port > uint32(math.MaxUint16) {
			return ErrNetworkPortBounds
		}

		c, ok := p.certificates.Get(networkKey)
		if !ok {
			continue
		}

		link := &networkBinding{
			networkKey:       networkKey,
			networkID:        c.networkID,
			localPort:        uint16(binding.Port),
			peerPort:         uint16(peerBinding.Port),
			localCertTrusted: isCertificateTrusted(binding.Certificate),
			peerCertTrusted:  isCertificateTrusted(peerBinding.Certificate),
		}
		p.links.ReplaceOrInsert(link)

		if !isCertificateTrusted(binding.Certificate) || !isCertificateTrusted(peerBinding.Certificate) {
			continue
		}

		if err := p.openNetwork(link); err != nil {
			return err
		}
	}
	return nil
}

func (p *Peer) openNetwork(link *networkBinding) error {
	if link.open || !link.localCertTrusted || !link.peerCertTrusted {
		return nil
	}
	link.open = true

	client, ok := p.vpn.Client(link.networkKey)
	if !ok {
		return ErrNetworkNotFound
	}
	client.Network.AddPeer(p.peer, link.localPort, link.peerPort)

	p.observers.Network.Emit(event.NetworkPeerOpen{
		PeerID:     p.id,
		NetworkID:  link.networkID,
		NetworkKey: link.networkKey,
	})

	p.logger.Info(
		"added peer to network",
		zap.Stringer("peer", p.peer.HostID()),
		logutil.ByteHex("network", link.networkKey),
		zap.Uint16("localPort", link.localPort),
		zap.Uint16("peerPort", link.peerPort),
	)

	return nil
}

type networkBinding struct {
	networkKey       []byte
	networkID        uint64
	localPort        uint16
	peerPort         uint16
	localCertTrusted bool
	peerCertTrusted  bool
	open             bool
}

func (b *networkBinding) Less(o llrb.Item) bool {
	if o, ok := o.(*networkBinding); ok {
		return bytes.Compare(b.networkKey, o.networkKey) == -1
	}
	return !o.Less(b)
}
