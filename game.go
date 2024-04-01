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
	idleState := NewHasIdleState(game.player, &game.joypad)
	jumpState := NewJumpState(game.player, &game.joypad)
	walkState := NewWalkState(game.player, &game.joypad)
	pivotState := NewHasPivotState(game.player, &game.joypad)
	runState := NewRunState(game.player, &game.joypad)

	// Set the states
	game.player.idle = idleState
	game.player.jumping = jumpState
	game.player.walking = walkState
	game.player.pivoting = pivotState
	game.player.running = runState

	game.player.setState(game.player.idle)

	return game
}
