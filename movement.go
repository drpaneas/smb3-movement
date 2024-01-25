package main

// updateMovement updates the motion of a player in both X and Y axis
func (p *Player) updateMovement(j Joypad) {
	p.updateVerticalMotion(j)   // all the logic for Y
	p.updateHorizontalMotion(j) // all the logix for X
}

// --------------------------------------------------
// Y axis Movement Logic
// --------------------------------------------------

// updateVerticalMotion updates the vertical motion of a player based on the current motion state and input from the joypad
func (p *Player) updateVerticalMotion(j Joypad) {
	// If the player is not airborne, check if he is going to be in this frame
	if p.MotionState != Airborne {
		// is A button just pressed in this frame?
		// if so, that means the player is going to jump in this frame!
		if j.JustPressed[A] {
			p.VelocityY = p.JumpInitialVelocity
			p.MotionState = Airborne
			p.applyVelocityY()
			p.Sprite.Y = int(SubpixelsToPx(p.PositionY))
			slowJumpFramesCounter = 24
		} else {
			// otherwise, the player is not airborne, and he is not going to jump in this frame
			p.VelocityY = 0
		}
	} else {
		slowJumpFramesCounter--
		// if the player is airborne, then he is going to fall in this frame
		p.updateJumpVelocity(j)
		p.applyVelocityY()
		p.boundPositionY()
	}
}

const (
	slowDeceleration spx = 1
	fastDeceleration spx = 5
)

var slowJumpFramesCounter = 25
var decelerate spx = fastDeceleration
var accelerationX spx = 0

// updateJumpVelocity updates the vertical velocity of a player when he is airborne
// it depends whether the A button is held down and how long it has been held down
func (p *Player) updateJumpVelocity(j Joypad) {
	if slowJumpFramesCounter > 0 && j.HoldDown[A] {
		decelerate = slowDeceleration
	} else {
		decelerate = fastDeceleration
	}

	// the velocity decreases by the deceleration rate, until it reaches the maximum fall speed
	p.VelocityY -= decelerate

	// Has it reached the maximum falling speed?
	if p.VelocityY <= p.JumpMaxFallSpeed {
		p.VelocityY = p.JumpMaxFallSpeed
	}
}

// applyVelocityY applies the vertical velocity to the vertical position of the player
func (p *Player) applyVelocityY() {
	p.PositionY -= p.VelocityY
}

func (p *Player) boundPositionY() {
	// Convert from subpixels into screen coordinates
	p.Sprite.Y = int(SubpixelsToPx(p.PositionY))

	// Check to see if the player has landed
	if p.Sprite.Y >= FloorHeight {
		// Land by re-initializing the Y variables and resetting the motion state
		p.ResetY()
		p.MotionState = Idle
	}
}

// --------------------------------------------------
// X axis Movement Logic
// --------------------------------------------------

// updateHorizontalMotion updates the horizontal motion of a player based on the current motion state and input from the joypad
func (p *Player) updateHorizontalMotion(j Joypad) {
	p.setTargetXVelocity(j)
	p.accelerateX(j)
	p.applyVelocityX()
	p.boundPositionX()
}

var (
	rightVelocity = []spx{24, 40}   // normal walk, running
	leftVelocity  = []spx{-24, -40} // normal walk, running (negative)
)

// setTargetXVelocity sets the target velocity of the player in the X-axis based on the given Joypad input.
func (p *Player) setTargetXVelocity(j Joypad) {
	index := 0
	if j.HoldDown[B] {
		index = 1
	}

	if j.HoldDown[Right] {
		p.TargetVelocityX = rightVelocity[index]
	} else if j.HoldDown[Left] {
		p.TargetVelocityX = leftVelocity[index]
	} else {
		p.TargetVelocityX = 0
	}
}

// accelerateX updates the X Velocity of the player based on the motion state and target X velocity
func (p *Player) accelerateX(j Joypad) {
	// When airborne, there is no friction, so ignore target velocities of 0
	if p.MotionState == Airborne && p.TargetVelocityX == 0 {
		return
	}

	if p.MotionState == Walk {
		accelerationX = 0.5
	} else if p.MotionState == Pivot {
		accelerationX = 2
	}

	if p.VelocityX < p.TargetVelocityX {
		p.VelocityX = p.VelocityX + accelerationX
	} else if p.VelocityX > p.TargetVelocityX {
		p.VelocityX = p.VelocityX - accelerationX
	}

}

// applyVelocityX applies the horizontal velocity to the player's position.
func (p *Player) applyVelocityX() {
	p.PositionX += p.VelocityX
}

const (
	SpriteWidth = 16
	RightBound  = screenWidth - SpriteWidth/2 - 8
	LeftBound   = SpriteWidth/2 + 8 // anchor for the player is at the middle
)

// boundPositionX checks to see if the player has hit the screen boundaries
// and if so, re-initializes the X variables and sets the X velocity to 0
// to stop the player from moving.
func (p *Player) boundPositionX() {
	p.Sprite.X = int(SubpixelsToPx(p.PositionX)) // Convert subpixels to screen coordinates

	// Check to see if the player has hit the screen boundaries
	if p.Sprite.X > RightBound {
		// Hit the right boundary: re-initialize the X variables and set the X velocity to 0
		p.Sprite.X = RightBound
		p.PositionX = PxToSubpixels(px(RightBound))
		p.VelocityX = 0
	} else if p.Sprite.X < LeftBound {
		// Hit the left boundary: re-initialize the X variables and set the X velocity to 0
		p.Sprite.X = LeftBound
		p.PositionX = PxToSubpixels(px(LeftBound))
		p.VelocityX = 0
	}
}
