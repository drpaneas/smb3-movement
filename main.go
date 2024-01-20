package main

// This is a simple example of how to use the gopxl/pixel/v2 library program
// the physics of a platformer game that is inspired by the movement of Super Mario Bros 3.
// Although the movement it's not 100% identical, it has the very same feeling as SMB3;
//
// Based on NES Hacker's code: https://github.com/NesHacker/PlatformerMovement/
// Youtube video: https://www.youtube.com/watch?v=ZuKIUjw_tNU

import (
	"log"
	"time"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
)

const (
	ScreenWidth  = 256
	ScreenHeight = 240
	FPS          = 60
)

func run() {
	// --------------------- //
	// Window configuration //
	// --------------------- //

	cfg := opengl.WindowConfig{
		Title:       "SMB3 inspired movement",
		Bounds:      pixel.R(0, 0, ScreenWidth, ScreenHeight),
		VSync:       true,
		Resizable:   false,
		Undecorated: false,
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Failed to create window: %v", err)
	}

	win.SetSmooth(false) // disable anti-aliasing for pixel art

	// ------------------------------------------------------ //
	// Initialize the game before the main loop begins to run //
	// ------------------------------------------------------ //

	// Initialize
	var (
		joypad = NewJoypad()
		player = NewPlayer()
	)

	// Load the background
	bg, err := loadPicture(background)
	if err != nil {
		panic(err)
	}

	bgSprite := pixel.NewSprite(bg, bg.Bounds())

	// Configure the game loop to run in 60 FPS
	frameTime := time.Second / FPS
	ticker := time.NewTicker(frameTime) // ticks every frame
	start := time.Now()
	frame := 0

	for !win.Closed() {
		// Create a loop that waits for the ticker to tick
		// this is the main game loop, called every frame tick
		for range ticker.C {
			frame++

			// ---------------- //
			// Main Game Loop   // It does 3 things:
			// ---------------- //

			joypad.Update(win)            // 1. Get input from the keyboard or the gamepad
			player.updateMovement(joypad) // 2. Update the player's movement
			player.updateSprite(joypad)   // 3. Update the player's sprite & animation keyframes

			// ---------------- //
			// Render Graphics  // Calls the Draw() method of the sprites
			// ---------------- //

			// Draw the background
			bgSprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

			// Draw the player
			player.Sprite.SPR.Draw(win, player.Sprite.Mat)

			// Render the frame to the screen
			win.Update()

			// If a second has passed since the ticker started, reset the frame counter
			since := time.Since(start)
			if since > time.Second {
				start = time.Now()
				frame = 0
			}
		}
	}
}

func main() {
	opengl.Run(run)
}
