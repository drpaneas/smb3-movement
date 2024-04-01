package main

// JumpState is the airborn state of Mario from the moment the player is in the air until it lands back down to the ground
type JumpState struct {
	player               *Player
	input                *Joypad
	slowJumpFrames       int
	slowDeceleration     spx
	fastDeceleration     spx
	initialVelocity      spx
	veryFastDeceleration spx
	isFirstJumpFrame     bool
}

func NewJumpState(player *Player, input *Joypad) *JumpState {
	return &JumpState{
		player:               player,
		input:                input,
		slowJumpFrames:       24,
		slowDeceleration:     1,
		fastDeceleration:     5,
		initialVelocity:      55,
		veryFastDeceleration: -64,
		isFirstJumpFrame:     true,
	}
}

func (s *JumpState) Update() {
	s.updateVerticalMotion()
	s.updateHorizontalMotion()
	s.boundCheck()
}

func (s *JumpState) boundCheck() {
	if s.player.Sprite.Y >= FloorHeight {
		s.player.ResetY()
		if s.player.VelocityX != 0 {
			s.isFirstJumpFrame = true
			s.player.setState(s.player.walking)
		} else {
			s.isFirstJumpFrame = true
			s.player.setState(s.player.idle)
		}
	}
}

func (s *JumpState) updateMidAir() {
	if s.input.HoldDown[Left] && s.player.VelocityX > 0 {
		s.player.VelocityX -= 1
	} else if s.input.HoldDown[Right] && s.player.VelocityX < 0 {
		s.player.VelocityX += 1
	} else if s.input.HoldDown[Left] && s.player.VelocityX < 0 {
		if s.player.VelocityX > s.player.TargetVelocityX {
			s.player.VelocityX -= 0.5 // walk left
		}
	} else if s.input.HoldDown[Right] && s.player.VelocityX > 0 {
		if s.player.VelocityX < s.player.TargetVelocityX {
			s.player.VelocityX += 0.5 // walk right
		}
	}
}

func (s *JumpState) updateVerticalMotion() {
	if s.isFirstJumpFrame {
		s.isFirstJumpFrame = false
		s.player.VelocityY = s.initialVelocity
		s.player.PositionY -= s.player.VelocityY
		s.player.Sprite.Y = int(SubpixelsToPx(s.player.PositionY))
		s.slowJumpFrames = 24
	}

	// If A is _not_ pressed, it will decelerate faster
	decelerate := s.fastDeceleration
	s.slowJumpFrames--
	if s.slowJumpFrames > 0 && s.input.HoldDown[A] {
		decelerate = s.slowDeceleration
	}

	// Velocity decreases by the deceleration rate, until it reaches the maximum fall speed
	s.player.VelocityY -= decelerate
	if s.player.VelocityY <= s.veryFastDeceleration {
		s.player.VelocityY = s.veryFastDeceleration
	}

	// Update the Position Y and apply it into the Sprite's screen coordinates
	s.player.PositionY -= s.player.VelocityY
	s.player.Sprite.Y = int(SubpixelsToPx(s.player.PositionY))
}

func (s *JumpState) updateHorizontalMotion() {
	// Check for direction change mid-air
	if s.player.PositionY > FloorHeight {
		s.updateMidAir()
	}

	// Update the PositionX and apply it into the Sprite's screen coordinates
	s.player.PositionX += s.player.VelocityX
	s.player.Sprite.X = int(SubpixelsToPx(s.player.PositionX)) // Convert subpixels to screen coordinates

}
