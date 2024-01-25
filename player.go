package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// MotionState is the state of the player
type MotionState int

const (
	Idle     MotionState = iota // Standing still
	Walk                        // Walking
	Pivot                       // Turning around while walking
	Airborne                    // Jumping or falling (not on the ground basically)
)

// IdleState is the state of the player when idle
// He closes his eyes (blinks) and opens them every now and then
type IdleState int

const (
	Still IdleState = iota
	Blink1
	Still2
	Blink2
)

type Heading int

const (
	FaceRight Heading = iota
	FaceLeft
)

type Player struct {
	// Sprites
	Pictures [5]*ebiten.Image // 5 animation frames
	Sprite   Sprite

	// X axis
	TargetVelocityX spx // desired velocity
	VelocityX       spx // current velocity
	PositionX       spx // current position
	Heading         Heading

	// Y axis
	VelocityY spx // current velocity
	PositionY spx // current position

	// Jump
	JumpInitialVelocity spx
	JumpMaxFallSpeed    spx

	// States
	MotionState MotionState
	IdleState   IdleState

	// Animation
	AnimationFrame int
	AnimationTimer int
	IdleTimer      int
}

// NewPlayer creates a new player instance and loads the sprites from the binary data
func NewPlayer() *Player {
	// Initialize sprites from binary data
	pic := [5]*ebiten.Image{}
	marioSprites := [5][]byte{mario0, mario1, mario2, mario3, mario4}
	for i := 0; i < 5; i++ {
		var err error
		pic[i], err = loadPicture(marioSprites[i]) // load mario0, mario1, mario2, mario3, mario4
		if err != nil {
			panic(err)
		}
	}

	sprite := NewSprite(pic)

	return &Player{
		// Initialize sprites
		Pictures: pic,
		Sprite:   *sprite,

		// Initialize X axis and direction
		TargetVelocityX: 0,
		VelocityX:       0,
		PositionX:       PxToSubpixels(px(InitialSpriteX)),
		Heading:         FaceRight,

		// Initialize Y axis
		VelocityY: 0,
		PositionY: PxToSubpixels(px(InitialSpriteY)),

		// Jump stuff
		JumpInitialVelocity: 55,  // 55
		JumpMaxFallSpeed:    -64, // -64

		// State Machines
		MotionState: Idle,
		IdleState:   Still,

		// Keyframes and animations
		AnimationFrame: 0,
		AnimationTimer: delayByVelocity[0],
		IdleTimer:      timers[0],
	}
}

// ResetY resets the Y position and velocity of the player
// This is used when the player lands on the ground after jumping or falling
func (p *Player) ResetY() {
	p.Sprite.Y = InitialSpriteY
	p.PositionY = PxToSubpixels(px(InitialSpriteY))
	p.VelocityY = 0
}
