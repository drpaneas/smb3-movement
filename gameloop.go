package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.joypad.Update()                 // 1. Get input from the keyboard or the gamepad
	g.player.updateMovement(g.joypad) // 2. Update the player's movement
	g.player.updateSprite(g.joypad)   // 3. Update the player's sprite & animation keyframes

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
