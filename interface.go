package main

import "time"

// AgonesSDK defines the contract that both the Real and Mock SDKs must satisfy.
type AgonesSDK interface {
	// Lifecycle
	Ready() error
	Allocate() error
	Shutdown() error
	Reserve(duration time.Duration) error

	// Health
	Health() error

	// Metadata
	SetLabel(key, value string) error
	SetAnnotation(key, value string) error

	// Data
	// We use interface{} here to avoid strict dependency on the official sdk.GameServer struct
	// in your tests, but you can cast it in your application code.
	GameServer() (interface{}, error)
	WatchGameServer(f func(gs interface{})) error

	// Features
	Alpha() AlphaFeature
}

// AlphaFeature defines the contract for Player Tracking (Alpha).
type AlphaFeature interface {
	PlayerConnect(id string) (bool, error)
	PlayerDisconnect(id string) (bool, error)
	SetPlayerCapacity(capacity int64) error
	GetPlayerCapacity() (int64, error)
	GetPlayerCount() (int64, error)
	IsPlayerConnected(id string) (bool, error)
	GetConnectedPlayers() ([]string, error)
}
