package main

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
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
func (j *Joypad) Update(win *opengl.Window) {

	// Update HoldDown
	j.HoldDown[B] = win.Pressed(pixel.KeyX) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadSquare)
	j.HoldDown[A] = win.Pressed(pixel.KeyZ) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadA)
	j.HoldDown[Select] = win.Pressed(pixel.KeyBackspace)
	j.HoldDown[Start] = win.Pressed(pixel.KeyEnter)
	j.HoldDown[Up] = win.Pressed(pixel.KeyUp) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadDpadUp)
	j.HoldDown[Down] = win.Pressed(pixel.KeyDown) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadDpadDown)
	j.HoldDown[Left] = win.Pressed(pixel.KeyLeft) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadDpadLeft)
	j.HoldDown[Right] = win.Pressed(pixel.KeyRight) || win.JoystickPressed(pixel.Joystick1, pixel.GamepadDpadRight)

	// Update JustPressed
	j.JustPressed[B] = win.JustPressed(pixel.KeyX) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadSquare)
	j.JustPressed[A] = win.JustPressed(pixel.KeyZ) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadA)
	j.JustPressed[Select] = win.JustPressed(pixel.KeyBackspace)
	j.JustPressed[Start] = win.JustPressed(pixel.KeyEnter)
	j.JustPressed[Up] = win.JustPressed(pixel.KeyUp) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadDpadUp)
	j.JustPressed[Down] = win.JustPressed(pixel.KeyDown) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadDpadDown)
	j.JustPressed[Left] = win.JustPressed(pixel.KeyLeft) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadDpadLeft)
	j.JustPressed[Right] = win.JustPressed(pixel.KeyRight) || win.JoystickJustPressed(pixel.Joystick1, pixel.GamepadDpadRight)
}
