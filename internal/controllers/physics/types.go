package physics

import (
	"image/color"
	"math/rand/v2"
)

type ParticleType int

const (
	ParticleTypeEmpty = ParticleType(iota)
	ParticleTypeSand
	ParticleTypeWater
)

type Particle struct {
	Type  ParticleType
	Color color.RGBA
}

func ParticleSand() Particle {
	return Particle{
		Type:  ParticleTypeSand,
		Color: ParticleColorSand(),
	}
}

func ParticleWater() Particle {
	return Particle{
		Type:  ParticleTypeWater,
		Color: ParticleColorWater(),
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

func ParticleColorWater() color.RGBA {
	return color.RGBA{
		R: uint8(0 + rand.UintN(50)),   // Base blue color (0) with random variation up to 50
		G: uint8(0 + rand.UintN(100)),  // Base greenish-blue tint (0) with random variation
		B: uint8(255 - rand.UintN(50)), // Base blue tint (255) with random darker variation
		A: 255,                         // Fully opaque
	}
}
