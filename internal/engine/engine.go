package engine

import (
	"context"
	"log/slog"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Engine struct {
	logger *slog.Logger

	window *glfw.Window

	prog uint32

	textureW, textureH int
	vao, vbo, texture  uint32
	pixelData          []byte

	keyStates      map[glfw.Key]bool
	mouseStates    map[glfw.MouseButton]bool
	mouseX, mouseY int

	controllers []Controller
}

func Init(
	ctx context.Context, logger *slog.Logger, controllers ...Controller,
) (*Engine, error) {
	e := &Engine{
		logger:      logger,
		keyStates:   make(map[glfw.Key]bool),
		mouseStates: make(map[glfw.MouseButton]bool),
		controllers: controllers,
	}

	if err := e.initGLFW(ctx); err != nil {
		return nil, err
	}

	if err := e.initOpenGL(ctx); err != nil {
		return nil, err
	}

	if err := e.initRender(ctx, textureWidth, textureHeight); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Engine) Run(ctx context.Context) {
	ectx := e.initContext(ctx)

	for _, c := range e.controllers {
		c.Init(ectx)
	}

	for !e.window.ShouldClose() && ctx.Err() == nil {
		e.preTick(ectx)

		for _, c := range e.controllers {
			c.Tick(ectx)
		}

		e.render(ectx)
	}
}

func (e *Engine) Shutdown() {
	e.shutdownRender()
	e.shutdownOpenGL()
	e.shutdownGLFW()
}

func (e *Engine) preTick(ctx *eCtx) {
	glfw.PollEvents()

	for key, state := range e.keyStates {
		ctx.input.keys[key] = state
	}
	for button, state := range e.mouseStates {
		ctx.input.mouseButtons[button] = state
	}
	ctx.input.mouseX, ctx.input.mouseY = e.mouseX, e.mouseY
}
