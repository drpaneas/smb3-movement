package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FloorHeight    = screenHeight - 17 - 16 // 17 is the height of the floor and 8 (16/2) is half height of the sprite (anchor point is in the middle)
	InitialSpriteY = FloorHeight
	InitialSpriteX = 30
)

// X and Y should be in the same unit that Draw function supports in the GoPXL library
type Sprite struct {
	Y, X  float64
	SPR   *ebiten.Image
	HFlip bool
}

// NewSprite creates a new sprite instance  with the given pictures
func NewSprite(pic [5]*ebiten.Image) *Sprite {

	sp := &Sprite{
		Y:     InitialSpriteY,
		HFlip: false,
		X:     InitialSpriteX,
		SPR:   ebiten.NewImageFromImage(pic[0]),
	}

	return sp
}

// updateSprite updates the sprite of the player
func (p *Player) updateSprite() {
	// p.updateMotionState(j)
	p.updateAnimationFrame()
	p.updateHeading()
	p.updateSpriteTiles()
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
			index := int(math.Abs(float64(p.VelocityX))) * 15 // this is a total random thing but it works :P
			p.AnimationTimer = delayByVelocity[index]
			// fmt.Println(p.AnimationTimer)

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
	8, 8, 8, 8, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 6, 6, 5, 5, 5, 5, 5, 5, 5, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2,
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

// updateSpriteTiles updates the sprite tiles of the player
// based on the current motion state and animation frame of the player
func (p *Player) updateSpriteTiles() {
	var (
		// Indeces of the tiles to use for the animation keyframes
		// based on the current motion state of the player
		jumpingTiles = []int{4}
		pivotTiles   = []int{2}
		walkTiles    = []int{0, 3}
		idleTiles    = []int{0}

		// The index of the tile to use for the current frame
		tiles []int
		index int
	)

	// Set the tiles and index based on the current motion state
	switch p.currentState {
	case p.idle:
		tiles = idleTiles
		index = 0
	case p.jumping:
		tiles = jumpingTiles
		index = 0
	case p.walking:
		tiles = walkTiles
		index = p.AnimationFrame
	case p.running:
		tiles = walkTiles
		index = p.AnimationFrame
	case p.pivoting:
		tiles = pivotTiles
		index = 0
	}

	// Set the sprite to the current tile
	pic := p.Pictures[tiles[index]]
	p.Sprite.SPR = ebiten.NewImageFromImage(pic)
}
