package main

import (
	"math"

	"github.com/gopxl/pixel/v2"
)

const (
	FloorHeight    = 17 + 8 // 17 is the height of the floor and 8 (16/2) is half height of the sprite (anchor point is in the middle)
	InitialSpriteY = FloorHeight
	InitialSpriteX = 30
)

// X and Y should be in the same unit that Draw function supports in the GoPXL library
type Sprite struct {
	Y, X  int
	SPR   *pixel.Sprite
	HFlip bool
	Mat   pixel.Matrix
}

// NewSprite creates a new sprite instance  with the given pictures
func NewSprite(pic [5]pixel.Picture) *Sprite {

	sp := &Sprite{
		Y:     InitialSpriteY,
		HFlip: false,
		X:     InitialSpriteX,
		SPR:   pixel.NewSprite(pic[0], pic[0].Bounds()),
		Mat:   pixel.IM,
	}

	return sp
}

// updateSprite updates the sprite of the player
func (p *Player) updateSprite(j Joypad) {
	p.updateMotionState(j)
	p.updateAnimationFrame()
	p.updateHeading()
	p.updateIdleState()
	p.updateSpriteTiles()
}

// updateMotionState updates the motion state of the player
// based on the current motion state and the target velocity of the player
func (p *Player) updateMotionState(j Joypad) {
	if p.MotionState < Airborne {
		if p.Sprite.X == 0 {
			p.MotionState = Idle
		} else if p.TargetVelocityX == p.VelocityX {
			if p.VelocityX == 0 {
				p.MotionState = Idle
			} else {
				p.MotionState = Walk
			}
		} else if (j.HoldDown[Left] || j.HoldDown[Right] || j.JustPressed[Left] || j.JustPressed[Right]) &&
			((p.TargetVelocityX > 0 && p.VelocityX < 0) || (p.TargetVelocityX < 0 && p.VelocityX > 0)) {
			p.MotionState = Pivot
		} else {
			p.MotionState = Walk
		}
	}
}

// updateAnimationFrame updates the animation frame of the player
// and controls the speed of the animation based on the velocity of the player
// if the player is not moving, the animation will be paused
// if the player is moving, the animation will be played at a speed that is proportional to the velocity of the player
// using linear interpolation (lerp) to calculate the speed of the animation (see the delayByVelocity table below)
func (p *Player) updateAnimationFrame() {

	if p.VelocityX == 0 {
		// Set initial timer
		p.AnimationTimer = delayByVelocity[0]
	} else {
		// Decrement timer
		p.AnimationTimer--

		if p.AnimationTimer == 0 {
			// Reset frame timer
			index := int(math.Abs(float64(p.VelocityX)))
			p.AnimationTimer = delayByVelocity[index]

			// Toggle the frame (walk animation has 2 frames only)
			p.AnimationFrame = 1 - p.AnimationFrame
		}
	}
}

// delayByVelocity is a lookup table that maps the absolute value of the velocity to the number of frames to wait before
// transitioning to the next frame of the walk animation. The timer is decremented every frame and when it reaches 0, the
// frame is toggled. The index is the absolute value of the velocity. The array is in order from 0 to 40, because the
// it gets faster as the velocity increases (less frames to wait before change to the next frame).
// To generate this table, I used the following code:
//
// --> Go Playground: https://go.dev/play/p/wvdjEdnOYi_D
var delayByVelocity = []int{
	8, 8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2,
}

// updateHeading updates the heading of the player
func (p *Player) updateHeading() {
	// Check if the target velocity is not zero
	if p.TargetVelocityX != 0 {
		// By default, set the desired heading to FaceRight
		desiredHeading := FaceRight

		// If the target velocity is less than zero, set the desired heading to FaceLeft
		if p.TargetVelocityX < 0 {
			desiredHeading = FaceLeft
		}

		// If the current heading is not the desired heading
		if p.Heading != desiredHeading {
			// Update the heading to the desired heading
			p.Heading = desiredHeading

			// If the desired heading is FaceLeft, flip the sprite horizontally
			p.Sprite.HFlip = desiredHeading == FaceLeft
		}
	}
}

// the number of frames to wait before transitioning to the next idlestate
// 245 - is the number of frames to wait in the initial idle state.
// if the player character remains idle for this many frames, it will transition to the next idle state.
//
// 10 - is the number of frames to wait in each subsequent idle state
var timers = []int{245, 10, 10, 10}

// updateIdleState updates the idle state of the player
func (p *Player) updateIdleState() {

	if p.MotionState != Idle {
		p.IdleTimer = timers[0]
		p.IdleState = Still
	} else {
		p.IdleTimer--
		if p.IdleTimer == 0 {
			p.IdleState++
			if int(p.IdleState) >= len(timers) {
				p.IdleState = 0
			}
			p.IdleTimer = timers[p.IdleState]
		}
	}
}

// updateSpriteTiles updates the sprite tiles of the player
// based on the current motion state and animation frame of the player
func (p *Player) updateSpriteTiles() {
	var (
		// Indeces of the tiles to use for the animation keyframes
		// based on the current motion state of the player
		jumpingTiles = []int{4}
		pivotTiles   = []int{2}
		walkTiles    = []int{0, 3}
		idleTiles    = []int{0, 1, 0, 1}

		// The index of the tile to use for the current frame
		tiles []int
		index int
	)

	// Set the tiles and index based on the current motion state
	switch p.MotionState {
	case Idle:
		tiles = idleTiles
		index = int(p.IdleState)
	case Airborne:
		tiles = jumpingTiles
		index = 0
	case Walk:
		tiles = walkTiles
		index = p.AnimationFrame
	case Pivot:
		tiles = pivotTiles
		index = 0
	}

	// Set the sprite to the current tile
	pic := p.Pictures[tiles[index]]
	p.Sprite.SPR.Set(pic, pic.Bounds())

	// Set the sprite position
	screenPosition := pixel.V(float64(p.Sprite.X), float64(p.Sprite.Y))
	p.Sprite.Mat = pixel.IM.Moved(screenPosition)

	// NOTE: Flip the sprite if the player is facing left
	if p.Sprite.HFlip {
		// convert degrees to radians using math standard library
		flipLeft := pixel.IM.Rotated(screenPosition, 180*math.Pi/180)
		p.Sprite.Mat = p.Sprite.Mat.Chained(flipLeft)

		flipVertical := pixel.IM.ScaledXY(screenPosition, pixel.V(1, -1))
		p.Sprite.Mat = p.Sprite.Mat.Chained(flipVertical)
	}

	// Now the sprite is ready to be drawn (rendered to the screen)
}
