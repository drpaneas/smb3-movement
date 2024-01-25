package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Button represents the different buttons on a gamepad, similar to NES.
type Button int

const (
	A Button = iota
	B
	Select
	Start
	Up
	Down
	Left
	Right
)

// Joypad represents the input from a gamepad
//   - HoldDown     : stores the state of the buttons that are being held down continuously
//   - JustPressed  : stores the state of the buttons that have been pressed in the current frame
type Joypad struct {
	HoldDown    map[Button]bool
	JustPressed map[Button]bool
}

// Create a new joypad and initialize the HoldDown and JustPressed maps
func NewJoypad() Joypad {
	return Joypad{
		HoldDown:    make(map[Button]bool),
		JustPressed: make(map[Button]bool),
	}
}

// Update the state of the joypad based on the input from the window
// This is supposed to be called every frame
func (j *Joypad) Update() {

	// gamepad ID
	var gamepadIDsBuf []ebiten.GamepadID
	gamepadIDsBuf = inpututil.AppendJustConnectedGamepadIDs(gamepadIDsBuf[:0])
	var id ebiten.GamepadID
	if len(gamepadIDsBuf) > 0 {
		id = gamepadIDsBuf[0]
	}

	// Update HoldDown
	j.HoldDown[B] = ebiten.IsKeyPressed(ebiten.KeyX) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton0)
	j.HoldDown[A] = ebiten.IsKeyPressed(ebiten.KeyZ) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton1)
	j.HoldDown[Select] = ebiten.IsKeyPressed(ebiten.KeyBackspace) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton8)
	j.HoldDown[Start] = ebiten.IsKeyPressed(ebiten.KeyEnter) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton9)
	j.HoldDown[Up] = ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton11)
	j.HoldDown[Down] = ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton12)
	j.HoldDown[Left] = ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton15)
	j.HoldDown[Right] = ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsGamepadButtonPressed(id, ebiten.GamepadButton13)

	// Update JustPressed
	j.JustPressed[B] = inpututil.IsKeyJustPressed(ebiten.KeyX) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton0)
	j.JustPressed[A] = inpututil.IsKeyJustPressed(ebiten.KeyZ) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton1)
	j.JustPressed[Select] = inpututil.IsKeyJustPressed(ebiten.KeyBackspace) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton8)
	j.JustPressed[Start] = inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton9)
	j.JustPressed[Up] = inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton11)
	j.JustPressed[Down] = inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton12)
	j.JustPressed[Left] = inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton15)
	j.JustPressed[Right] = inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsGamepadButtonJustPressed(id, ebiten.GamepadButton13)
}
