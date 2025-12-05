package main

import (
	"time"

	sdk "agones.dev/agones/pkg/sdk"
	agones "agones.dev/agones/sdks/go"
)

// RealSDK wraps the official SDK to satisfy the AgonesSDK interface.
type RealSDK struct {
	client *agones.SDK
}

func NewRealSDK() (*RealSDK, error) {
	s, err := agones.NewSDK()
	if err != nil {
		return nil, err
	}
	return &RealSDK{client: s}, nil
}

// Passthrough methods
func (r *RealSDK) Ready() error                    { return r.client.Ready() }
func (r *RealSDK) Allocate() error                 { return r.client.Allocate() }
func (r *RealSDK) Shutdown() error                 { return r.client.Shutdown() }
func (r *RealSDK) Reserve(d time.Duration) error   { return r.client.Reserve(d) }
func (r *RealSDK) SetLabel(k, v string) error      { return r.client.SetLabel(k, v) }
func (r *RealSDK) SetAnnotation(k, v string) error { return r.client.SetAnnotation(k, v) }
func (r *RealSDK) Health() error {
	return r.client.Health()
}

// Data methods - Returning the raw struct as interface{}
func (r *RealSDK) GameServer() (interface{}, error) { return r.client.GameServer() }
func (r *RealSDK) WatchGameServer(f func(interface{})) error {
	return r.client.WatchGameServer(func(gs *sdk.GameServer) {
		f(gs)
	})
}

// Alpha Wrapper
// We need this because r.client.Alpha() returns *sdk.Alpha (struct),
// but our interface expects AlphaFeature (interface).
func (r *RealSDK) Alpha() AlphaFeature {
	return r.client.Alpha()
}
