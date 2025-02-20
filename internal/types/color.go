package types

import (
	"math/rand/v2"
)

type Color struct {
	R, G, B, A byte
}

func ColorBlack() Color {
	return Color{0, 0, 0, 255}
}

func ColorRandom() Color {
	return Color{
		R: byte(rand.UintN(256)),
		G: byte(rand.UintN(256)),
		B: byte(rand.UintN(256)),
		A: 255,
	}
}
