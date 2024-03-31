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

	// Initialize states
	idleState := &HasIdleState{
		player: game.player,
		input:  &game.joypad,
	}

	jumpState := &HasJumpState{
		player: game.player,
		input:  &game.joypad,
	}

	walkState := &HasWalkState{
		player: game.player,
		input:  &game.joypad,
	}

	pivotState := &HasPivotState{
		player: game.player,
		input:  &game.joypad,
	}

	runState := &HasRunState{
		player: game.player,
		input:  &game.joypad,
	}

	// Set the states
	game.player.idle = idleState
	game.player.jumping = jumpState
	game.player.walking = walkState
	game.player.pivoting = pivotState
	game.player.running = runState

	game.player.setState(game.player.idle)

	return game
}
