package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// GameObject is considered anything that can be updated and drawn on the screen
type GameObject interface {
	Update()
	Draw(screen *ebiten.Image)
}

// Game is the main struct for our game that holds all the important information
type Game struct {
	joypad Joypad
	player *Player
}

func newGame() *Game {
	// Create the game
	game := &Game{
		joypad: NewJoypad(),
		player: NewPlayer(),
	}

	return game
}
