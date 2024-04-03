package main

import "math"

type RunState struct {
	player *Player
	input  *Joypad
}

func NewRunState(player *Player, input *Joypad) *RunState {
	return &RunState{
		player: player,
		input:  input,
	}
}

func (s *RunState) Update() {
	s.updateTargetVelocity()
	s.updateVelocity()
	s.applyVelocity()
	s.checkBoundaries()
	s.checkJump()
	s.checkPivot()
}

func (s *RunState) updateTargetVelocity() {
	if s.input.HoldDown[Left] && s.input.HoldDown[B] {
		s.player.TargetVelocityX = -maxRunSpeed
	} else if s.input.HoldDown[Right] && s.input.HoldDown[B] {
		s.player.TargetVelocityX = maxRunSpeed
	} else if (s.input.HoldDown[Right] || s.input.HoldDown[Left]) && !s.input.HoldDown[B] {
		s.player.setState(s.player.walking)
	} else {
		s.player.TargetVelocityX = 0
	}

	// It never goes to exactly 0, so this is an approximation.
	if math.Abs(s.player.VelocityX) <= 0.01 { // Adjust threshold as needed
		s.player.setState(s.player.idle)
	}
}

func (s *RunState) updateVelocity() {
	if s.player.VelocityX < s.player.TargetVelocityX {
		s.player.VelocityX += runAcceleration
	} else if s.player.VelocityX > s.player.TargetVelocityX {
		s.player.VelocityX -= runAcceleration
	} else if s.player.VelocityX == 0 {
		s.player.setState(s.player.idle)
	}
}

func (s *RunState) applyVelocity() {
	s.player.PositionX += s.player.VelocityX
	s.player.Sprite.X = s.player.PositionX
}

func (s *RunState) checkBoundaries() {
	if s.player.Sprite.X > RightBound {
		s.player.Sprite.X = RightBound
		s.player.PositionX = RightBound
		s.player.setState(s.player.idle)
	} else if s.player.Sprite.X < LeftBound {
		s.player.Sprite.X = LeftBound
		s.player.PositionX = LeftBound
		s.player.setState(s.player.idle)
	}
}

func (s *RunState) checkJump() {
	if s.input.JustPressed[A] {
		s.player.setState(s.player.jumping)
	}
}

func (s *RunState) checkPivot() {
	if (s.input.HoldDown[Left] || s.input.HoldDown[Right] || s.input.JustPressed[Left] || s.input.JustPressed[Right]) &&
		((s.player.TargetVelocityX > 0 && s.player.VelocityX < 0) || (s.player.TargetVelocityX < 0 && s.player.VelocityX > 0)) {
		s.player.setState(s.player.pivoting)
	}
}
