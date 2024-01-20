package main

// We need to use subpixels to achieve smooth movement.
// This is a trick that is used in many games back in the day.

type px int
type spx float64

// 1 pixel = 16 subpixels
const SubpixelsPerPx = 16.0

// Convert pixels to subpixels and vice versa
func PxToSubpixels(p px) spx {
	return spx(p * SubpixelsPerPx)
}

func SubpixelsToPx(s spx) px {
	return px(s / SubpixelsPerPx)
}
