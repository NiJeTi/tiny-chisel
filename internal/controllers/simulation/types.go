package simulation

import (
	"github.com/nijeti/graphics/internal/types"
)

type ParticleType int

const (
	ParticleTypeEmpty = ParticleType(iota)
)

type Particle struct {
	Type       ParticleType
	Lifetime   float32
	Velocity   types.Vector
	Color      types.Color
	IsUpToDate bool
}
