package simulation

import (
	"image/color"
	"math/rand/v2"

	"github.com/nijeti/graphics/internal/types"
)

type ParticleType int

const (
	ParticleTypeEmpty = ParticleType(iota)
	ParticleTypeSand
)

type Particle struct {
	Type       ParticleType
	Lifetime   float32
	Velocity   types.Vector
	Color      color.RGBA
	IsUpToDate bool
}

func ParticleSand() Particle {
	return Particle{
		Type:       ParticleTypeSand,
		Lifetime:   0,
		Velocity:   types.Vector{X: 0, Y: 0},
		Color:      ParticleColorSand(),
		IsUpToDate: true,
	}
}

func ParticleColorSand() color.RGBA {
	return color.RGBA{
		R: uint8(194 + rand.UintN(30)), // Base sand color (194) with random variation up to 30
		G: uint8(178 + rand.UintN(30)), // Base greenish tint (178) with random variation
		B: uint8(128 + rand.UintN(20)), // Base brownish tint (128) with random variation
		A: 255,                         // Fully opaque
	}
}
