package engine

import (
	"context"
	"fmt"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Run(ctx context.Context, opts ...Option) error {
	ectx := newCtx(ctx)
	for _, opt := range opts {
		opt(ectx)
	}
	ectx.Init()

	ectx.logger.InfoContext(ctx, "initializing")

	window, err := initGLFW(ectx)
	if err != nil {
		return err
	}
	defer shutdownGLFW(ectx)

	program, err := initOpenGL(ectx)
	if err != nil {
		return err
	}
	defer shutdownOpenGL(ectx, program)

	vao, vbo, texture, err := initRender(ectx, program)
	if err != nil {
		return err
	}
	defer shutdownRender(ectx, vao, vbo, texture)

	ectx.logger.InfoContext(ctx, "running")

	for _, c := range ectx.controllers {
		c.Init(ectx)
	}

	for frame := 0; !window.ShouldClose() && ctx.Err() == nil; frame++ {
		frameStart := time.Now()

		preTick()

		for _, c := range ectx.controllers {
			c.Tick(ectx)
		}

		render(ectx, window)

		frameTime := time.Since(frameStart)
		outputStats(ectx, window, frame, frameTime)
		time.Sleep(frameDelta - frameTime)
	}

	ectx.logger.InfoContext(ctx, "shutting down")

	return nil
}

func preTick() {
	glfw.PollEvents()
}

func outputStats(
	ctx *ectx, window *glfw.Window, frame int, frameTime time.Duration,
) {
	if frame%3 != 0 {
		return
	}

	title := fmt.Sprintf(
		"%s | FT: %.2fms",
		ctx.windowTitle, float64(frameTime.Microseconds())/1000,
	)
	window.SetTitle(title)
}
