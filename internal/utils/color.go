package utils

import (
	"image/color"
	"math/rand/v2"
)

func ColorBlack() color.RGBA {
	return color.RGBA{R: 0, G: 0, B: 0, A: 255}
}

func ColorRandom() color.RGBA {
	return color.RGBA{
		R: uint8(rand.UintN(256)),
		G: uint8(rand.UintN(256)),
		B: uint8(rand.UintN(256)),
		A: 255,
	}
}
