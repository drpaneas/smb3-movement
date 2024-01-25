package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed sprites/mario0.png
var mario0 []byte

//go:embed sprites/mario1.png
var mario1 []byte

//go:embed sprites/mario2.png
var mario2 []byte

//go:embed sprites/mario3.png
var mario3 []byte

//go:embed sprites/mario4.png
var mario4 []byte

//go:embed sprites/background.png
var background []byte

// loadPicture loads a picture from a file on the disk
func loadPicture(file []byte) (*ebiten.Image, error) {

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}
