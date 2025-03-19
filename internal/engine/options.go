package engine

import (
	"log/slog"
)

type Option func(ctx *ectx)

func WithLogger(logger *slog.Logger) Option {
	if logger == nil {
		panic("logger is nil")
	}

	return func(ctx *ectx) {
		ctx.logger = logger
	}
}

func ConfigureWindow(title string, width, height int, resizable bool) Option {
	if width <= 0 || height <= 0 {
		panic("window size is invalid")
	}

	return func(ctx *ectx) {
		ctx.windowTitle = title
		ctx.windowW = width
		ctx.windowH = height
		ctx.windowResizable = resizable
	}
}

func ConfigureSpace(width, height int) Option {
	if width <= 0 || height <= 0 {
		panic("space size is invalid")
	}

	return func(ctx *ectx) {
		ctx.spaceW = width
		ctx.spaceH = height
	}
}

func WithControllers(cs ...Controller) Option {
	return func(ctx *ectx) {
		ctx.controllers = cs
	}
}
