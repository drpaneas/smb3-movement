package main

type HasIdleState struct {
	player *Player
	input  *Joypad
}

func (i *HasIdleState) Update() {
	i.player.VelocityX = 0
	i.player.VelocityY = 0

	// Change into Jumping state
	if i.input.HoldDown[A] {
		i.player.setState(i.player.jumping)
		i.player.MotionState = Airborne

		i.player.TargetVelocityX = 0
		i.player.VelocityY = i.player.JumpInitialVelocity
		i.player.applyVelocityY()
		i.player.Sprite.Y = int(SubpixelsToPx(i.player.PositionY))
		slowJumpFramesCounter = 24
	}

	// Change into walking state
	if i.input.HoldDown[Left] {
		i.player.setState(i.player.walking)
		i.player.MotionState = Walk

		// Change into running state
		if i.input.HoldDown[B] {
			i.player.setState(i.player.running)
			i.player.MotionState = Walk
			i.player.TargetVelocityX = -40
		}
	} else if i.input.HoldDown[Right] {
		i.player.setState(i.player.walking)
		i.player.MotionState = Walk

		// Change into runnign state
		if i.input.HoldDown[B] {
			i.player.setState(i.player.running)
			i.player.MotionState = Walk
			i.player.TargetVelocityX = 40
		}
	}
}
