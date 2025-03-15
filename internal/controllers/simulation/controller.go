package simulation

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/nijeti/graphics/internal/engine"
	"github.com/nijeti/graphics/internal/types"
)

type Controller struct {
	width, height int
	particles     [][]Particle
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Init(ctx engine.Context) {
	width, height := ctx.SpaceSize()

	particles := make([][]Particle, 0, width)
	for range width {
		row := make([]Particle, 0, height)
		for range height {
			p := Particle{
				Type:       ParticleTypeEmpty,
				Lifetime:   0,
				Velocity:   types.VectorZero(),
				Color:      types.ColorBlack(),
				IsUpToDate: true,
			}
			row = append(row, p)
		}
		particles = append(particles, row)
	}

	c.width = width
	c.height = height
	c.particles = particles
}

func (*Controller) Tick(ctx engine.Context) {
	if ctx.MouseButtonState(glfw.MouseButtonLeft) {
		x, y := ctx.MousePos()
		ctx.SetPixel(x, y, types.ColorRandom())
	}
}
