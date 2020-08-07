package node

import (
	"context"

	"github.com/golang/geo/s2"
)

// Driver ...
type Driver interface {
	Provider() string
	Regions(ctx context.Context, req *RegionsRequest) ([]*Region, error)
	SKUs(ctx context.Context, req *SKUsRequest) ([]*SKU, error)
	Create(ctx context.Context, req *CreateRequest) (*Node, error)
	List(ctx context.Context, req *ListRequest) ([]*Node, error)
	Delete(ctx context.Context, req *DeleteRequest) error
}

// RegionsRequest ...
type RegionsRequest struct {
}

// SKUsRequest ...
type SKUsRequest struct {
	Region string
}

// CreateRequest ...
type CreateRequest struct {
	Name    string
	Region  string
	SKU     string
	SSHKeys []string
}

// ListRequest ...
type ListRequest struct {
}

// DeleteRequest ...
type DeleteRequest struct {
	Region     string
	ProviderID string
}

// Region represents the Node's datacenter location
type Region struct {
	Name   string
	City   string
	LatLng s2.LatLng
}

// SKU ...
type SKU struct {
	Name         string
	CPUs         int
	Memory       int
	NetworkCap   int
	NetworkSpeed int
	PriceMonthly *Price
	PriceHourly  *Price
}

// Price ...
type Price struct {
	Value    float64
	Currency string
}

// Node represents a host
type Node struct {
	ProviderID string    `json:"provider_id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Memory     int       `json:"memory,omitempty"`
	CPUs       int       `json:"vcpus,omitempty"`
	Disk       int       `json:"disk,omitempty"`
	Networks   *Networks `json:"networks,omitempty"`
	Status     string    `json:"status,omitempty"`
	Region     *Region   `json:"region,omitempty"`
	SKU        *SKU      `json:"sku,omitempty"`
}

// Networks represents the Node's networks.
type Networks struct {
	V4 []string `json:"v4,omitempty"`
	V6 []string `json:"v6,omitempty"`
}
