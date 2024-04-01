package main

type HasPivotState struct {
	player *Player
	input  *Joypad
}

func NewHasPivotState(player *Player, input *Joypad) *HasPivotState {
	return &HasPivotState{
		player: player,
		input:  input,
	}
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
	} else {
		s.player.setState(s.player.idle)
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
	}
}
