package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/nijeti/graphics/internal/controllers/physics"
	"github.com/nijeti/graphics/internal/engine"
)

const (
	codeOk  = 0
	codeErr = 1
)

func main() {
	runtime.LockOSThread()

	code := run()
	os.Exit(code)
}

func run() (code int) {
	loggerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, loggerOpts))

	defer func() {
		if err := recover(); err != nil {
			logger.Error("panic", "error", err)
			code = codeErr
		}
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	err := engine.Run(
		ctx,
		engine.WithLogger(logger.With("module", "engine")),
		engine.ConfigureWindow("tiny-chisel", 1280, 720, true),
		engine.ConfigureSpace(640, 360),
		engine.WithControllers(physics.New()),
	)
	if err != nil {
		logger.Error("failed to run engine", "error", err)
		return codeErr
	}

	return codeOk
}
