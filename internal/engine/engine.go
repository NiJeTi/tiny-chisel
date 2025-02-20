package engine

import (
	"context"
	"log/slog"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/nijeti/graphics/internal/simulation"
)

type Engine struct {
	logger            *slog.Logger
	window            *glfw.Window
	prog              uint32
	vao, vbo, texture uint32
	pixels            []byte
	simulation        *simulation.Simulation
}

func Init(ctx context.Context, logger *slog.Logger) (*Engine, error) {
	e := &Engine{
		logger: logger,
	}

	if err := e.initGLFW(ctx); err != nil {
		return nil, err
	}

	if err := e.initOpenGL(ctx); err != nil {
		return nil, err
	}

	if err := e.initRender(ctx); err != nil {
		return nil, err
	}

	if err := e.initSimulation(ctx); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Engine) Run(ctx context.Context) {
	particles := e.simulation.Particles()

	for !e.window.ShouldClose() && ctx.Err() == nil {
		e.simulation.Tick()

		for x := range particles {
			for y, p := range particles[x] {
				offset := (x*textureHeight + y) * sizeColor
				e.pixels[offset+0] = p.Color.R
				e.pixels[offset+1] = p.Color.G
				e.pixels[offset+2] = p.Color.B
				e.pixels[offset+3] = p.Color.A
			}
		}

		e.render(ctx)
	}
}

func (e *Engine) Shutdown() {
	e.shutdownRender()
	e.shutdownOpenGL()
	e.shutdownGLFW()
}
