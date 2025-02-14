// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package network

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math"
	"sync"
	"sync/atomic"

	"github.com/MemeLabs/strims/internal/api"
	"github.com/MemeLabs/strims/internal/dao"
	"github.com/MemeLabs/strims/internal/event"
	networkv1 "github.com/MemeLabs/strims/pkg/apis/network/v1"
	"github.com/MemeLabs/strims/pkg/apis/type/certificate"
	"github.com/MemeLabs/strims/pkg/logutil"
	"github.com/MemeLabs/strims/pkg/vnic"
	"github.com/MemeLabs/strims/pkg/vnic/qos"
	"github.com/MemeLabs/strims/pkg/vpn"
	"github.com/petar/GoLLRB/llrb"
	"go.uber.org/zap"
)

type Peer interface {
	HandlePeerNegotiate(keyCount uint32)
	HandlePeerOpen(bindings []*networkv1.NetworkPeerBinding)
	HandlePeerClose(networkKey []byte)
	HandlePeerUpdateCertificate(cert *certificate.Certificate) error
}

var _ Peer = (*peer)(nil)

// NewPeer ...
func newPeer(
	id uint64,
	vnicPeer *vnic.Peer,
	client api.PeerClient,
	logger *zap.Logger,
	observers *event.Observers,
	broker Broker,
	vpn *vpn.Host,
	qosc *qos.Class,
	certificates *certificateMap,
) *peer {
	return &peer{
		id:       id,
		client:   client,
		vnicPeer: vnicPeer,
		logger: logger.With(
			zap.Uint64("id", id),
			logutil.ByteHex("host", vnicPeer.HostID().Bytes(nil)),
		),
		observers:    observers,
		broker:       broker,
		vpn:          vpn,
		certificates: certificates,

		keyCount:   make(chan uint32, 1),
		bindings:   make(chan []*networkv1.NetworkPeerBinding, 1),
		brokerConn: vnicPeer.Channel(vnic.NetworkBrokerPort, qosc),
	}
}

// Peer ...
type peer struct {
	id           uint64
	vnicPeer     *vnic.Peer
	client       api.PeerClient
	logger       *zap.Logger
	observers    *event.Observers
	broker       Broker
	vpn          *vpn.Host
	certificates *certificateMap

	lock        sync.Mutex
	links       llrb.LLRB
	negotiating atomic.Bool
	keyCount    chan uint32
	bindings    chan []*networkv1.NetworkPeerBinding
	brokerConn  *vnic.FrameReadWriter
}

// HandlePeerNegotiate ...
func (p *peer) HandlePeerNegotiate(keyCount uint32) {
	select {
	case p.keyCount <- keyCount:
	default:
	}

	if !p.negotiating.Load() {
		go func() {
			if err := p.negotiateNetworks(context.Background()); err != nil {
				p.logger.Debug("network negotiation failed", zap.Error(err))
			}
		}()
	}
}

// HandlePeerOpen ...
func (p *peer) HandlePeerOpen(bindings []*networkv1.NetworkPeerBinding) {
	select {
	case p.bindings <- bindings:
	default:
	}
}

// HandlePeerClose ...
func (p *peer) HandlePeerClose(networkKey []byte) {
	p.closeNetworkWithoutNotifyingPeer(networkKey)
}

// HandlePeerUpdateCertificate ...
func (p *peer) HandlePeerUpdateCertificate(cert *certificate.Certificate) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	if err := dao.VerifyCertificate(cert); err != nil {
		return err
	}
	if !isPeerCertificateOwner(p.vnicPeer, cert) {
		return ErrCertificateOwnerMismatch
	}
	if !isCertificateTrusted(cert) {
		return ErrProvisionalCertificate
	}

	li := p.links.Get(&networkBinding{networkKey: dao.CertificateNetworkKey(cert)})
	if li == nil {
		return ErrNetworkBindingNotFound
	}

	link := li.(*networkBinding)
	link.peerCertTrusted = true

	return p.openNetwork(link)
}

func (p *peer) sendCertificateUpdate(network *networkv1.Network) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	li := p.links.Get(&networkBinding{networkKey: dao.NetworkKey(network)})
	if li == nil {
		return ErrNetworkBindingNotFound
	}

	link := li.(*networkBinding)
	link.localCertTrusted = true

	err := p.client.Network().UpdateCertificate(
		context.Background(),
		&networkv1.NetworkPeerUpdateCertificateRequest{Certificate: network.Certificate},
		&networkv1.NetworkPeerUpdateCertificateResponse{},
	)
	if err != nil {
		return err
	}

	return p.openNetwork(link)
}

func (p *peer) closeNetworkWithoutNotifyingPeer(networkKey []byte) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	li := p.links.Get(&networkBinding{networkKey: networkKey})
	if li == nil {
		return ErrNetworkBindingNotFound
	}
	p.links.Delete(li)

	if li.(*networkBinding).open {
		node, ok := p.vpn.Node(networkKey)
		if !ok {
			return ErrNetworkNotFound
		}
		node.Network.RemovePeer(p.vnicPeer.HostID())

		defer p.observers.EmitLocal(event.NetworkPeerClose{
			PeerID:     p.id,
			NetworkID:  li.(*networkBinding).networkID,
			NetworkKey: networkKey,
		})

		p.logger.Info(
			"removed peer from network",
			zap.Stringer("peer", p.vnicPeer.HostID()),
			logutil.ByteHex("network", networkKey),
		)
	}

	if p.links.Len() == 0 {
		p.vnicPeer.Close()
	}

	return nil
}

func (p *peer) closeNetwork(networkKey []byte) {
	p.closeNetworkWithoutNotifyingPeer(networkKey)
	p.client.Network().Close(context.Background(), &networkv1.NetworkPeerCloseRequest{Key: networkKey}, &networkv1.NetworkPeerCloseResponse{})
}

func (p *peer) networkKeysForLinks() [][]byte {
	p.lock.Lock()
	defer p.lock.Unlock()

	keys := make([][]byte, p.links.Len())
	p.links.AscendLessThan(llrb.Inf(1), func(li llrb.Item) bool {
		keys = append(keys, li.(*networkBinding).networkKey)
		return true
	})
	return keys
}

func (p *peer) close() {
	for _, key := range p.networkKeysForLinks() {
		p.closeNetworkWithoutNotifyingPeer(key)
	}
}

func (p *peer) hasNetworkBinding(networkKey []byte) bool {
	p.lock.Lock()
	defer p.lock.Unlock()
	return p.links.Has(&networkBinding{networkKey: networkKey})
}

func (p *peer) negotiateNetworks(ctx context.Context) error {
	if !p.negotiating.CompareAndSwap(false, true) {
		return errors.New("cannot begin new negotiation until previous negotiation finishes")
	}
	defer p.negotiating.Store(false)

	select {
	case <-p.certificates.Loaded():
	case <-ctx.Done():
		return ctx.Err()
	}

	keys := p.certificates.Keys()
	err := p.client.Network().Negotiate(ctx, &networkv1.NetworkPeerNegotiateRequest{KeyCount: uint32(len(keys))}, &networkv1.NetworkPeerNegotiateResponse{})
	if err != nil {
		return fmt.Errorf("sending network negotiation init failed: %w", err)
	}

	return p.exchangeBindings(ctx, keys)
}

func (p *peer) exchangeBindings(ctx context.Context, keys [][]byte) error {
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
		preferSend := p.vnicPeer.HostID().Less(p.vpn.VNIC().ID())
		if len(keys) > int(keyCount) || (len(keys) == int(keyCount) && preferSend) {
			return p.exchangeBindingsAsSender(ctx, keys)
		}
		return p.exchangeBindingsAsReceiver(ctx, keys)
	}
}

func (p *peer) exchangeBindingsAsReceiver(ctx context.Context, keys [][]byte) error {
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

		p.observers.EmitLocal(event.NetworkPeerBindings{PeerID: p.id, NetworkKeys: keys})

		return p.handleNetworkBindings(networkBindings, peerNetworkBindings)
	}
}

func (p *peer) exchangeBindingsAsSender(ctx context.Context, keys [][]byte) error {
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

		p.observers.EmitLocal(event.NetworkPeerBindings{PeerID: p.id, NetworkKeys: keys})

		networkBindings, err := p.sendNetworkBindings(ctx, keys)
		if err != nil {
			return err
		}

		return p.handleNetworkBindings(networkBindings, peerNetworkBindings)
	}
}

func (p *peer) sendNetworkBindings(ctx context.Context, keys [][]byte) ([]*networkv1.NetworkPeerBinding, error) {
	var bindings []*networkv1.NetworkPeerBinding

	for _, key := range keys {
		c, ok := p.certificates.Get(key)
		if !ok {
			return nil, fmt.Errorf("certificate not found: %w", ErrNetworkNotFound)
		}

		if p.links.Has(&networkBinding{networkKey: key}) {
			continue
		}

		port, err := p.vnicPeer.ReservePort()
		if err != nil {
			return nil, err
		}

		bindings = append(
			bindings,
			&networkv1.NetworkPeerBinding{
				Port:        uint32(port),
				Certificate: c.certificate,
			},
		)
	}

	err := p.client.Network().Open(ctx, &networkv1.NetworkPeerOpenRequest{Bindings: bindings}, &networkv1.NetworkPeerOpenResponse{})
	if err != nil {
		return nil, err
	}
	return bindings, nil
}

func (p *peer) verifyNetworkBindings(bindings []*networkv1.NetworkPeerBinding) ([][]byte, error) {
	if bindings == nil {
		return nil, ErrNetworkBindingsEmpty
	}

	keys := make([][]byte, len(bindings))
	for i, b := range bindings {
		if err, ok := dao.VerifyCertificate(b.Certificate).(dao.Errors); ok && !err.IncludesOnly(dao.ErrNotAfterRange) {
			return nil, err
		}
		keys[i] = dao.CertificateNetworkKey(b.Certificate)
	}
	return keys, nil
}

func (p *peer) handleNetworkBindings(networkBindings, peerNetworkBindings []*networkv1.NetworkPeerBinding) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	for i, peerBinding := range peerNetworkBindings {
		binding := networkBindings[i]
		networkKey := dao.CertificateNetworkKey(peerBinding.Certificate)

		if !isPeerCertificateOwner(p.vnicPeer, peerBinding.Certificate) {
			return ErrCertificateOwnerMismatch
		}
		if !bytes.Equal(dao.CertificateNetworkKey(binding.Certificate), networkKey) {
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

func (p *peer) openNetwork(link *networkBinding) error {
	if link.open || !link.localCertTrusted || !link.peerCertTrusted {
		return nil
	}
	link.open = true

	node, ok := p.vpn.Node(link.networkKey)
	if !ok {
		return ErrNetworkNotFound
	}
	node.Network.AddPeer(p.vnicPeer, link.localPort, link.peerPort)

	p.observers.EmitLocal(event.NetworkPeerOpen{
		PeerID:     p.id,
		NetworkID:  link.networkID,
		NetworkKey: link.networkKey,
	})

	p.logger.Info(
		"added peer to network",
		zap.Stringer("peer", p.vnicPeer.HostID()),
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
