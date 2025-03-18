package engine

import (
	"context"
	"log/slog"
	"sync"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Engine struct {
	logger *slog.Logger

	window *glfw.Window

	prog uint32

	textureW, textureH int
	vao, vbo, texture  uint32
	textureData        []byte

	fixedDelta time.Duration
	delta      time.Duration
	lastFrame  time.Time

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
		fixedDelta:  fixedDelta,
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	done := func() { wg.Done() }

	go e.fixedTick(ectx, done)
	e.tick(ectx)

	wg.Wait()
}

func (e *Engine) Shutdown() {
	e.shutdownRender()
	e.shutdownOpenGL()
	e.shutdownGLFW()
}

func (e *Engine) tick(ctx Context) {
	for e.isRunning(ctx) {
		now := time.Now()
		e.delta = now.Sub(e.lastFrame)
		e.lastFrame = now

		e.preFrame()

		for _, c := range e.controllers {
			c.Tick(ctx)
		}

		e.render(ctx)
	}
}

func (e *Engine) fixedTick(ctx Context, done func()) {
	defer done()

	for e.isRunning(ctx) {
		start := time.Now()
		for _, c := range e.controllers {
			c.FixedTick(ctx)
		}
		duration := time.Since(start)

		time.Sleep(e.fixedDelta - duration)
	}
}

func (*Engine) preFrame() {
	glfw.PollEvents()
}

func (e *Engine) isRunning(ctx Context) bool {
	return !e.window.ShouldClose() && ctx.Err() == nil
}
