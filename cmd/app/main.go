package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

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
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.TimeValue(a.Value.Time().UTC())
			}
			return a
		},
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

	// init
	logger.Info("initializing")

	e, err := engine.Init(ctx, logger.With("module", "engine"))
	if err != nil {
		logger.Error("failed to initialize engine", "error", err)
		return codeErr
	}
	defer e.Shutdown()
	logger.Info("engine initialized")

	// run
	logger.InfoContext(ctx, "running")

	e.Run(ctx)

	// exit
	logger.Info("exiting")

	return codeOk
}
