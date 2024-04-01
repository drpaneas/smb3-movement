package main

type HasIdleState struct {
	player *Player
	input  *Joypad
}

func NewHasIdleState(p *Player, i *Joypad) *HasIdleState {
	return &HasIdleState{player: p, input: i}
}

func (s *HasIdleState) Update() {
	s.player.VelocityX = 0
	s.player.VelocityY = 0

	if s.input.JustPressed[A] {
		s.transitionToJumping()
		return
	}

	if s.input.HoldDown[Left] {
		s.transitionToWalking()

		if s.input.HoldDown[B] {
			s.transitionToRunning(-40)
		}
	} else if s.input.HoldDown[Right] {
		s.transitionToWalking()

		if s.input.HoldDown[B] {
			s.transitionToRunning(40)
		}
	}
}

func (s *HasIdleState) transitionToJumping() {
	s.player.setState(s.player.jumping)
}

func (s *HasIdleState) transitionToWalking() {
	s.player.setState(s.player.walking)
}

func (s *HasIdleState) transitionToRunning(targetVelocityX spx) {
	s.player.setState(s.player.running)
	s.player.TargetVelocityX = targetVelocityX
}
