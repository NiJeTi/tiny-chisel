package physics

import (
	"math/rand/v2"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/nijeti/graphics/internal/engine"
	"github.com/nijeti/graphics/internal/utils"
)

const (
	brushRadius  = 3
	brushOpacity = 0.5
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
				Type:  ParticleTypeEmpty,
				Color: utils.ColorBlack(),
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
	defer c.render(ctx)

	switch {
	case ctx.MouseButtonState(glfw.MouseButtonLeft):
		c.draw(ctx, ParticleSand())
	case ctx.MouseButtonState(glfw.MouseButtonRight):
		c.draw(ctx, ParticleWater())
	}
}

func (c *Controller) FixedTick(engine.Context) {
	for x := c.width - 1; x >= 0; x-- {
		for y := c.height - 1; y >= 0; y-- {
			switch c.particles[x][y].Type {
			case ParticleTypeSand:
				c.processSand(x, y)
			case ParticleTypeWater:
				c.processWater(x, y)
			}
		}
	}
}

func (c *Controller) render(ctx engine.Context) {
	for x := range c.particles {
		for y := range c.particles[x] {
			ctx.SetPixel(x, y, c.particles[x][y].Color)
		}
	}
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
	c.particles[srcX][srcY], c.particles[dstX][dstY] =
		c.particles[dstX][dstY], c.particles[srcX][srcY]
}

func (c *Controller) draw(ctx engine.Context, p Particle) {
	x, y := ctx.MousePos()

	for dx := -brushRadius; dx <= brushRadius; dx++ {
		for dy := -brushRadius; dy <= brushRadius; dy++ {
			nx, ny := x+dx, y+dy

			if nx < 0 || nx >= c.width || ny < 0 || ny >= c.height {
				continue
			}

			if dx*dx+dy*dy > brushRadius*brushRadius {
				continue
			}

			if rand.Float32() < brushOpacity {
				c.particles[nx][ny] = p
			}
		}
	}
}

func (c *Controller) processSand(x, y int) {
	belowLeftX := x - 1
	belowRightX := x + 1
	belowY := y + 1

	switch {
	case c.isEmpty(x, belowY):
		c.swapParticles(x, y, x, belowY)
	case c.isEmpty(belowLeftX, belowY):
		c.swapParticles(x, y, belowLeftX, belowY)
	case c.isEmpty(belowRightX, belowY):
		c.swapParticles(x, y, belowRightX, belowY)
	}
}

func (c *Controller) processWater(x, y int) {
	leftX, rightX := x-1, x+1
	belowY := y + 1

	switch {
	case c.isEmpty(x, belowY):
		c.swapParticles(x, y, x, belowY)
	case c.isEmpty(leftX, belowY):
		c.swapParticles(x, y, leftX, belowY)
	case c.isEmpty(rightX, belowY):
		c.swapParticles(x, y, rightX, belowY)
	case c.isEmpty(leftX, y):
		c.swapParticles(x, y, leftX, y)
	case c.isEmpty(rightX, y):
		c.swapParticles(x, y, rightX, y)
	}
}
