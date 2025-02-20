package engine

import (
	"context"
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func (e *Engine) initGLFW(ctx context.Context) error {
	if err := glfw.Init(); err != nil {
		return fmt.Errorf("failed to init glfw: %w", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, glVersionMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, glVersionMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(
		windowWidth, windowHeight, windowTitle, nil, nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create window: %w", err)
	}

	window.MakeContextCurrent()
	e.window = window

	e.logger.DebugContext(ctx, "glfw initialized")

	return nil
}

func (e *Engine) shutdownGLFW() {
	glfw.Terminate()

	e.logger.Debug("glfw shutdown complete")
}
