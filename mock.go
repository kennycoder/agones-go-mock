package main

import (
	"fmt"
	"time"
)

// MockSDK mimics the Agones SDK behavior by printing logs.
type MockSDK struct {
	alpha *MockAlpha
}

// NewMockSDK initializes the mock and the nested Alpha mock.
func NewMockSDK() (*MockSDK, error) {
	return &MockSDK{
		alpha: &MockAlpha{},
	}, nil
}

// logHelper handles the formatting of the timestamp and message.
func (m *MockSDK) logHelper(method string, extras ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	if len(extras) > 0 {
		fmt.Printf("[%s] Agones Mock: %s called with args: %v\n", timestamp, method, extras)
	} else {
		fmt.Printf("[%s] Agones Mock: %s called\n", timestamp, method)
	}
}

// --- Lifecycle Methods ---

func (m *MockSDK) Ready() error {
	m.logHelper("Ready")
	return nil
}

func (m *MockSDK) Allocate() error {
	m.logHelper("Allocate")
	return nil
}

func (m *MockSDK) Shutdown() error {
	m.logHelper("Shutdown")
	return nil
}

func (m *MockSDK) Reserve(duration time.Duration) error {
	m.logHelper("Reserve", duration)
	return nil
}

// --- Health Check ---

func (m *MockSDK) Health() error {
	m.logHelper("Health Ping Sent")
	return nil
}

// --- Configuration & Metadata ---

func (m *MockSDK) SetLabel(key, value string) error {
	m.logHelper("SetLabel", key, value)
	return nil
}

func (m *MockSDK) SetAnnotation(key, value string) error {
	m.logHelper("SetAnnotation", key, value)
	return nil
}

// GameServer usually returns *sdk.GameServer.
// We return nil here to keep this mock dependency-free.
func (m *MockSDK) GameServer() (interface{}, error) {
	m.logHelper("GameServer (Get Details)")
	return nil, nil
}

// WatchGameServer usually takes a callback function.
func (m *MockSDK) WatchGameServer(f func(gs interface{})) error {
	m.logHelper("WatchGameServer")
	return nil
}

// --- Feature Gates (Alpha/Beta) ---

func (m *MockSDK) Alpha() AlphaFeature {
	return m.alpha
}

// MockAlpha handles the Alpha feature set (Player Tracking, etc.)
type MockAlpha struct{}

func (a *MockAlpha) logHelper(method string, extras ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	fmt.Printf("[%s] Agones Mock [Alpha]: %s called with args: %v\n", timestamp, method, extras)
}

func (a *MockAlpha) PlayerConnect(id string) (bool, error) {
	a.logHelper("PlayerConnect", id)
	return true, nil
}

func (a *MockAlpha) PlayerDisconnect(id string) (bool, error) {
	a.logHelper("PlayerDisconnect", id)
	return true, nil
}

func (a *MockAlpha) SetPlayerCapacity(capacity int64) error {
	a.logHelper("SetPlayerCapacity", capacity)
	return nil
}

func (a *MockAlpha) GetPlayerCapacity() (int64, error) {
	a.logHelper("GetPlayerCapacity")
	return 100, nil
}

func (a *MockAlpha) GetPlayerCount() (int64, error) {
	a.logHelper("GetPlayerCount")
	return 0, nil
}

func (a *MockAlpha) IsPlayerConnected(id string) (bool, error) {
	a.logHelper("IsPlayerConnected", id)
	return false, nil
}

func (a *MockAlpha) GetConnectedPlayers() ([]string, error) {
	a.logHelper("GetConnectedPlayers")
	return []string{}, nil
}
