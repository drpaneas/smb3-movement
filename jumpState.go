package main

import "fmt"

type HasJumpState struct {
	player *Player
	input  *Joypad
}

func (i *HasJumpState) Update() {
	slowJumpFramesCounter--

	// If A is pressed, it will decelerate faster
	if slowJumpFramesCounter > 0 && i.input.HoldDown[A] {
		decelerate = slowDeceleration
	} else {
		decelerate = fastDeceleration
	}

	// the velocity decreases by the deceleration rate, until it reaches the maximum fall speed
	i.player.VelocityY -= decelerate

	// Has it reached the maximum falling speed?
	if i.player.VelocityY <= i.player.JumpMaxFallSpeed {
		i.player.VelocityY = i.player.JumpMaxFallSpeed
	}

	// Update the Position Y
	i.player.PositionY -= i.player.VelocityY

	// Apply the position to the Sprite coordinates
	i.player.Sprite.Y = int(SubpixelsToPx(i.player.PositionY))

	// Check for direction change mid-air
	if i.player.PositionY > FloorHeight {
		fmt.Println("Mid air")
		if i.input.HoldDown[Left] && i.player.VelocityX > 0 {
			i.player.VelocityX -= 1

		} else if i.input.HoldDown[Right] && i.player.VelocityX < 0 {
			i.player.VelocityX += 1
		} else if i.input.HoldDown[Left] && i.player.VelocityX < 0 {
			if i.player.VelocityX > i.player.TargetVelocityX {
				i.player.VelocityX -= 0.5 // walk left
			}
			fmt.Println(i.player.VelocityX)
		} else if i.input.HoldDown[Right] && i.player.VelocityX > 0 {
			if i.player.VelocityX < i.player.TargetVelocityX {
				i.player.VelocityX += 0.5 // walk right
			}
			fmt.Println(i.player.VelocityX)
		}
	}

	// Apply the velocity X
	i.player.PositionX += i.player.VelocityX

	// Update the Sprite's X position
	i.player.Sprite.X = int(SubpixelsToPx(i.player.PositionX)) // Convert subpixels to screen coordinates

	// Check to see if the player has landed
	if i.player.Sprite.Y >= FloorHeight {
		// Land by re-initializing the Y variables and resetting the motion state
		i.player.ResetY()

		if i.player.VelocityX != 0 {
			i.player.setState(i.player.walking)
			i.player.MotionState = Walk
			if i.input.HoldDown[A] {
				i.player.setState(i.player.running)
				i.player.MotionState = Walk
			}
		} else {
			i.player.MotionState = Idle
			i.player.setState(i.player.idle)
			i.player.MotionState = Idle
		}
	}
}
