package main

type HasWalkState struct {
	player *Player
	input  *Joypad
}

func (i *HasWalkState) Update() {
	if i.input.HoldDown[Left] {
		i.player.TargetVelocityX = -24
	} else if i.input.HoldDown[Right] {
		i.player.TargetVelocityX = 24
	} else {
		i.player.TargetVelocityX = 0
	}

	accelerationX = 0.5

	if i.input.HoldDown[B] {
		i.player.setState(i.player.running)
		i.player.MotionState = Walk
	}

	if i.player.VelocityX < i.player.TargetVelocityX {
		i.player.VelocityX += accelerationX // walk right
	} else if i.player.VelocityX > i.player.TargetVelocityX {
		i.player.VelocityX -= accelerationX // walk left
	} else {
		if i.player.VelocityX == 0 {
			i.player.setState(i.player.idle)
			i.player.MotionState = Idle
		}
	}

	// Apply the velocity
	i.player.PositionX += i.player.VelocityX

	// Update the Sprite's position
	i.player.Sprite.X = int(SubpixelsToPx(i.player.PositionX)) // Convert subpixels to screen coordinates

	// Check to see if the player has hit the screen boundaries
	if i.player.Sprite.X > RightBound {
		// Hit the right boundary: re-initialize the X variables and set the X velocity to 0
		i.player.Sprite.X = RightBound
		i.player.PositionX = PxToSubpixels(px(RightBound))
		i.player.setState(i.player.idle)
		i.player.MotionState = Idle
		i.player.VelocityX = 0 // Add this line to set X velocity to 0

	} else if i.player.Sprite.X < LeftBound {
		// Hit the left boundary: re-initialize the X variables and set the X velocity to 0
		i.player.Sprite.X = LeftBound
		i.player.PositionX = PxToSubpixels(px(LeftBound))
		i.player.setState(i.player.idle)
		i.player.MotionState = Idle
		i.player.VelocityX = 0 // Add this line to set X velocity to 0

	}

	if i.input.JustPressed[A] {
		i.player.setState(i.player.jumping)
		i.player.MotionState = Airborne
	}

	if (i.input.HoldDown[Left] || i.input.HoldDown[Right] || i.input.JustPressed[Left] || i.input.JustPressed[Right]) &&
		((i.player.TargetVelocityX > 0 && i.player.VelocityX < 0) || (i.player.TargetVelocityX < 0 && i.player.VelocityX > 0)) {

		i.player.MotionState = Pivot
		i.player.setState(i.player.pivoting)
	}

}
