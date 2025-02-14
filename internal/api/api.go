// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	networkv1 "github.com/MemeLabs/strims/pkg/apis/network/v1"
	networkv1bootstrap "github.com/MemeLabs/strims/pkg/apis/network/v1/bootstrap"
	networkv1ca "github.com/MemeLabs/strims/pkg/apis/network/v1/ca"
	transferv1 "github.com/MemeLabs/strims/pkg/apis/transfer/v1"
)

// PeerClient ...
type PeerClient interface {
	Bootstrap() *networkv1bootstrap.PeerServiceClient
	CA() *networkv1ca.CAPeerClient
	Transfer() *transferv1.TransferPeerClient
	Network() *networkv1.NetworkPeerClient
}
