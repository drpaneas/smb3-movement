package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SpriteWidth = 16
	RightBound  = screenWidth - SpriteWidth
	LeftBound   = 0
)

func (g *Game) Update() error {
	g.joypad.Update()              // 1. Get input from the keyboard or the gamepad
	g.player.currentState.Update() // 2. Update the player state movement (physics)
	g.player.updateSprite()        // 3. Update the player's sprite & animation keyframes

	if g.player.currentState == g.player.idle {
		fmt.Println("idle")
	} else if g.player.currentState == g.player.jumping {
		fmt.Println("jumping")
	} else if g.player.currentState == g.player.walking {
		fmt.Println("walking")
	} else if g.player.currentState == g.player.pivoting {
		fmt.Println("pivoting")
	} else if g.player.currentState == g.player.running {
		fmt.Println("running")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(bg, nil) // Draw the background

	// Draw playerAnimation at g.player.Sprite.X, g.player.Sprite.Y coordinates
	// If the player is facing left, flip the sprite
	if g.player.Sprite.HFlip {
		op := &ebiten.DrawImageOptions{}
		w, h := g.player.Sprite.SPR.Bounds().Dx(), g.player.Sprite.SPR.Bounds().Dy()
		op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
		op.GeoM.Scale(-1, 1) // Flip horizontally
		op.GeoM.Translate(float64(w)/2, float64(h)/2)
		op.GeoM.Translate(float64(g.player.Sprite.X), float64(g.player.Sprite.Y))
		op.Filter = ebiten.FilterNearest
		screen.DrawImage(g.player.Sprite.SPR, op)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(g.player.Sprite.X), float64(g.player.Sprite.Y))
		op.Filter = ebiten.FilterNearest
		screen.DrawImage(g.player.Sprite.SPR, op)
	}
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
