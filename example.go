package main

import (
	"log"
	"os"
)

// GameServerLogic contains your actual game code.
// It asks for the INTERFACE, not the struct.
type GameServerLogic struct {
	agones AgonesSDK
}

func (g *GameServerLogic) Start() {
	// This code works identical for both Mock and Real
	err := g.agones.Ready()
	if err != nil {
		log.Printf("Failed to mark ready: %v", err)
	}

	// Use Alpha features
	g.agones.Alpha().SetPlayerCapacity(100)
}

func main() {
	var sdk AgonesSDK
	var err error

	// Check environment variable to decide which SDK to load
	if os.Getenv("AGONES_ENV") == "local" {
		log.Println("Initializing MOCK SDK...")
		sdk, err = NewMockSDK() // From previous step
	} else {
		log.Println("Initializing REAL SDK...")
		// var err error
		sdk, err = NewRealSDK()
	}

	if err != nil {
		panic(err)
	}

	// Inject dependency
	game := &GameServerLogic{
		agones: sdk,
	}

	game.Start()
}
