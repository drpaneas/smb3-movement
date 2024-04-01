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
	MaxFallSpeed        spx

	// States
	idle         State
	jumping      State
	walking      State
	running      State
	pivoting     State
	currentState State

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

	player := &Player{
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
		MaxFallSpeed:        -64, // -64

		// Keyframes and animations
		AnimationFrame: 0,
		AnimationTimer: delayByVelocity[0],
	}

	return player
}

// ResetY resets the Y position and velocity of the player
// This is used when the player lands on the ground after jumping or falling
func (p *Player) ResetY() {
	p.Sprite.Y = InitialSpriteY
	p.PositionY = PxToSubpixels(px(InitialSpriteY))
	p.VelocityY = 0
}

func (p *Player) setState(s State) {
	p.currentState = s
}
