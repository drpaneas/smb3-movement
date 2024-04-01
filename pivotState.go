package main

type HasPivotState struct {
	player *Player
	input  *Joypad
}

func (s *HasPivotState) Update() {
	s.updateVelocity()
	s.updatePosition()
	s.updateSpritePosition()
	s.checkJump()
}

func (s *HasPivotState) updateVelocity() {
	const walkSpeed = 1.5

	switch {
	case s.player.VelocityX < s.player.TargetVelocityX:
		s.player.VelocityX += walkSpeed
		if s.player.VelocityX > 0 {
			s.transitionState()
		}
	case s.player.VelocityX > s.player.TargetVelocityX:
		s.player.VelocityX -= walkSpeed
		if s.player.VelocityX < 0 {
			s.transitionState()
		}
	}
}

func (s *HasPivotState) transitionState() {
	if s.input.HoldDown[Left] || s.input.HoldDown[Right] {
		s.player.setState(s.player.walking)
		s.player.MotionState = Walk
	} else {
		s.player.setState(s.player.idle)
		s.player.MotionState = Idle
	}
}

func (s *HasPivotState) updatePosition() {
	s.player.PositionX += s.player.VelocityX
}

func (s *HasPivotState) updateSpritePosition() {
	s.player.Sprite.X = int(SubpixelsToPx(s.player.PositionX))
}

func (s *HasPivotState) checkJump() {
	if s.input.JustPressed[A] {
		s.player.setState(s.player.jumping)
		s.player.MotionState = Airborne
	}
}
