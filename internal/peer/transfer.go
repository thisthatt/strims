// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package peer

import (
	"context"

	"github.com/MemeLabs/strims/internal/app"
	"github.com/MemeLabs/strims/internal/transfer"
	transferv1 "github.com/MemeLabs/strims/pkg/apis/transfer/v1"
)

type transferService struct {
	Peer app.Peer
	App  app.Control
}

func (s *transferService) Announce(ctx context.Context, req *transferv1.TransferPeerAnnounceRequest) (*transferv1.TransferPeerAnnounceResponse, error) {
	id, err := transfer.ParseID(req.Id)
	if err != nil {
		return nil, err
	}
	channel, ok := s.Peer.Transfer().AssignPort(id, req.Channel)
	if ok {
		return &transferv1.TransferPeerAnnounceResponse{
			Body: &transferv1.TransferPeerAnnounceResponse_Channel{
				Channel: channel,
			},
		}, nil
	}
	return &transferv1.TransferPeerAnnounceResponse{}, nil
}

func (s *transferService) Close(ctx context.Context, req *transferv1.TransferPeerCloseRequest) (*transferv1.TransferPeerCloseResponse, error) {
	id, err := transfer.ParseID(req.Id)
	if err != nil {
		return nil, err
	}
	s.Peer.Transfer().Close(id)
	return &transferv1.TransferPeerCloseResponse{}, nil
}
