package main

// This is a simple example of how to use the EbitenEngine to program
// the physics of a platformer game that is inspired by the movement of Super Mario Bros 3.
// Although the movement it's not 100% identical, it has the very same feeling as SMB3;
//
// Based on NES Hacker's code: https://github.com/NesHacker/PlatformerMovement/
// Youtube video: https://www.youtube.com/watch?v=ZuKIUjw_tNU

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	img, err := loadPicture(background)
	if err != nil {
		panic(err)
	}

	bg = img

}

func main() {
	// Configure the game window
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeDisabled)
	ebiten.SetWindowTitle("SMB3 inspired movement")
	ebiten.SetFullscreen(false)

	game := newGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
