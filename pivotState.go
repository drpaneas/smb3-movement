package main

type HasPivotState struct {
	player *Player
	input  *Joypad
}

func (i *HasPivotState) Update() {

	if i.player.VelocityX < i.player.TargetVelocityX {
		i.player.VelocityX += 1.5 // walk right
		if i.player.VelocityX > 0 {
			if i.input.HoldDown[Left] || i.input.HoldDown[Right] {
				i.player.setState(i.player.walking)
				i.player.MotionState = Walk
			} else {
				i.player.setState(i.player.idle)
				i.player.MotionState = Idle
			}
		}
	} else if i.player.VelocityX > i.player.TargetVelocityX {
		i.player.VelocityX -= 1.5 // walk left
		if i.player.VelocityX < 0 {
			if i.input.HoldDown[Left] || i.input.HoldDown[Right] {
				i.player.setState(i.player.walking)
				i.player.MotionState = Walk
			} else {
				i.player.setState(i.player.idle)
				i.player.MotionState = Idle
			}
		}
	}

	// Apply the velocity
	i.player.PositionX += i.player.VelocityX

	// Update the Sprite's position
	i.player.Sprite.X = int(SubpixelsToPx(i.player.PositionX)) // Convert subpixels to screen coordinates
}
