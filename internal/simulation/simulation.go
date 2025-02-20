package simulation

import (
	"errors"

	"github.com/nijeti/graphics/internal/types"
)

type Simulation struct {
	width, height int
	matrix        [][]Particle
}

func Init(width, height int) (*Simulation, error) {
	if width <= 0 {
		return nil, errors.New("width must be a positive number")
	}
	if height <= 0 {
		return nil, errors.New("height must be a positive number")
	}

	matrix := make([][]Particle, 0, width)
	for range width {
		particles := make([]Particle, 0, height)
		for range height {
			p := Particle{
				Type:       ParticleTypeEmpty,
				Lifetime:   0,
				Velocity:   types.VectorZero(),
				Color:      types.ColorBlack(),
				IsUpToDate: true,
			}
			particles = append(particles, p)
		}
		matrix = append(matrix, particles)
	}

	return &Simulation{
		width:  width,
		height: height,
		matrix: matrix,
	}, nil
}

func (m *Simulation) Particles() [][]Particle {
	return m.matrix
}

func (m *Simulation) Tick() {
	for x := range m.width {
		for y := range m.height {
			m.matrix[x][y].Color = types.ColorRandom()
		}
	}
}

func (m *Simulation) swapParticles(fromX, fromY int, toX, toY int) {
	buf := m.matrix[fromX][fromY]
	m.matrix[fromX][fromY] = m.matrix[toX][toY]
	m.matrix[toX][toY] = buf
}
