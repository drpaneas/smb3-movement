package main

import (
	"math"
)

const (
	walkAcceleration     = 0.09
	walkDecel            = 0.16
	walkDecelerationStop = 0.06
	runAcceleration      = 0.09
	// runDeceleration      = 0.31

	maxWalkSpeed   = 1.25 // $14 -> 20 subpixels is 1.25 pixels per frame (75 pixels per second at 60 FPS)
	maxRunSpeed    = 2.25 // $24 -> 36 subpixels is 2.25 pixels per frame (135 pixels per second at 60 FPS)
	maxPMeterSpeed = 3    // $30 -> 48 subpixels is 3.00 pixels per frame (180 pixels per second at 60 FPS)
)

type WalkState struct {
	player *Player
	input  *Joypad
}

func NewWalkState(player *Player, input *Joypad) *WalkState {
	return &WalkState{
		player: player,
		input:  input,
	}
}

func (ws *WalkState) Update() {
	ws.updateTargetVelocity()
	ws.updateVelocity()
	ws.applyVelocity()
	ws.checkBoundaries()
	ws.handleJump()
	ws.handlePivot()
}

func (ws *WalkState) updateTargetVelocity() {

	// Check for any horizontal movement input
	isMovingLeft := ws.input.HoldDown[Left]
	isMovingRight := ws.input.HoldDown[Right]

	// Determine target velocity based on input
	if isMovingLeft {
		ws.player.TargetVelocityX = -maxWalkSpeed // Set negative for left movement
	} else if isMovingRight {
		ws.player.TargetVelocityX = maxWalkSpeed
	} else {
		ws.player.TargetVelocityX = 0 // Set to 0 if no movement input
	}

}

func (ws *WalkState) updateVelocity() {
	// fmt.Println(math.Abs(ws.player.VelocityX))
	if ws.input.HoldDown[B] {
		ws.player.setState(ws.player.running)
		return
	}

	if ws.player.TargetVelocityX == maxWalkSpeed {
		if ws.player.VelocityX < ws.player.TargetVelocityX {
			ws.player.VelocityX += walkAcceleration
		}
	}

	if ws.player.TargetVelocityX == -maxWalkSpeed {
		if ws.player.VelocityX > ws.player.TargetVelocityX {
			ws.player.VelocityX -= walkAcceleration
		}
	}

	if ws.player.TargetVelocityX == 0 {
		if ws.player.VelocityX > 0 {
			ws.player.VelocityX -= walkDecelerationStop
		}
		if ws.player.VelocityX < 0 {
			ws.player.VelocityX += walkDecelerationStop
		}

		// It never goes to exactly 0, so this is an approximation.
		if math.Abs(ws.player.VelocityX) <= 0.06 { // Adjust threshold as needed
			ws.player.setState(ws.player.idle)
		}
	}
}

func (ws *WalkState) applyVelocity() {
	ws.player.PositionX += ws.player.VelocityX
	ws.player.Sprite.X = ws.player.PositionX
}

func (ws *WalkState) checkBoundaries() {
	switch {
	case ws.player.Sprite.X > RightBound:
		ws.player.Sprite.X = RightBound
		ws.player.PositionX = RightBound
		ws.player.setState(ws.player.idle)
		ws.player.VelocityX = 0
	case ws.player.Sprite.X < LeftBound:
		ws.player.Sprite.X = LeftBound
		ws.player.PositionX = LeftBound
		ws.player.setState(ws.player.idle)
		ws.player.VelocityX = 0
	}
}

func (ws *WalkState) handleJump() {
	if ws.input.JustPressed[A] {
		ws.player.setState(ws.player.jumping)
	}
}

func (ws *WalkState) handlePivot() {
	if (ws.input.HoldDown[Left] || ws.input.HoldDown[Right] || ws.input.JustPressed[Left] || ws.input.JustPressed[Right]) &&
		((ws.player.TargetVelocityX > 0 && ws.player.VelocityX < 0) || (ws.player.TargetVelocityX < 0 && ws.player.VelocityX > 0)) {
		ws.player.setState(ws.player.pivoting)
	}
}
