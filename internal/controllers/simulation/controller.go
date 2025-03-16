package simulation

import (
	"math/rand/v2"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/nijeti/graphics/internal/engine"
	"github.com/nijeti/graphics/internal/types"
	"github.com/nijeti/graphics/internal/utils"
)

const brushRadius = 3

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
				Color:      utils.ColorBlack(),
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

func (c *Controller) Tick(ctx engine.Context) {
	if ctx.MouseButtonState(glfw.MouseButtonLeft) {
		x, y := ctx.MousePos()

		// Fill particles with sand within the brushRadius around (x, y)
		for dx := -brushRadius; dx <= brushRadius; dx++ {
			for dy := -brushRadius; dy <= brushRadius; dy++ {
				nx, ny := x+dx, y+dy
				// Check if the position is within the bounds of the grid
				if nx >= 0 && nx < c.width && ny >= 0 && ny < c.height {
					// Check if the position falls within the circular brush radius
					if dx*dx+dy*dy <= brushRadius*brushRadius {
						if rand.IntN(2) == 1 {
							c.particles[nx][ny] = ParticleSand()
						}
					}
				}
			}
		}
	}

	for x := c.width - 1; x >= 0; x-- {
		for y := c.height - 1; y >= 0; y-- {
			belowX, belowY := x, y+1
			belowLeftX := belowX - 1
			belowRightY := belowX + 1

			switch {
			case c.isEmpty(belowX, belowY):
				c.swapParticles(x, y, belowX, belowY)
			case c.isEmpty(belowLeftX, belowY):
				c.swapParticles(x, y, belowLeftX, belowY)
			case c.isEmpty(belowRightY, belowY):
				c.swapParticles(x, y, belowRightY, belowY)
			}
		}
	}

	c.toSpace(ctx)
}

func (c *Controller) isEmpty(x, y int) bool {
	p, ok := c.particle(x, y)
	return ok && p.Type == ParticleTypeEmpty
}

func (c *Controller) particle(x, y int) (Particle, bool) {
	if x >= c.width || x < 0 {
		return Particle{}, false
	}
	if y >= c.height || y < 0 {
		return Particle{}, false

	}

	return c.particles[x][y], true
}

func (c *Controller) swapParticles(srcX, srcY int, dstX, dstY int) {
	temp := c.particles[srcX][srcY]
	c.particles[srcX][srcY] = c.particles[dstX][dstY]
	c.particles[dstX][dstY] = temp
}

func (c *Controller) toSpace(ctx engine.Context) {
	for x := range c.particles {
		for y := range c.particles[x] {
			ctx.SetPixel(x, y, c.particles[x][y].Color)
		}
	}
}
