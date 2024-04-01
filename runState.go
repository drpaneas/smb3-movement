package main

const (
	accelerationX spx = 0.5
)

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
		s.player.TargetVelocityX = -40
	} else if s.input.HoldDown[Right] && s.input.HoldDown[B] {
		s.player.TargetVelocityX = 40
	} else if (s.input.HoldDown[Right] || s.input.HoldDown[Left]) && !s.input.HoldDown[B] {
		s.player.setState(s.player.walking)
	} else {
		s.player.TargetVelocityX = 0
	}
}

func (s *RunState) updateVelocity() {
	if s.player.VelocityX < s.player.TargetVelocityX {
		s.player.VelocityX += accelerationX
	} else if s.player.VelocityX > s.player.TargetVelocityX {
		s.player.VelocityX -= accelerationX
	} else if s.player.VelocityX == 0 {
		s.player.setState(s.player.idle)
	}
}

func (s *RunState) applyVelocity() {
	s.player.PositionX += s.player.VelocityX
	s.player.Sprite.X = int(SubpixelsToPx(s.player.PositionX))
}

func (s *RunState) checkBoundaries() {
	if s.player.Sprite.X > RightBound {
		s.player.Sprite.X = RightBound
		s.player.PositionX = PxToSubpixels(px(RightBound))
		s.player.setState(s.player.idle)
	} else if s.player.Sprite.X < LeftBound {
		s.player.Sprite.X = LeftBound
		s.player.PositionX = PxToSubpixels(px(LeftBound))
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
