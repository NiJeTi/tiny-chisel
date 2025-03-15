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

	window.SetKeyCallback(e.keyCallback)
	window.SetMouseButtonCallback(e.mouseCallback)
	window.SetCursorPosCallback(e.cursorPositionCallback)

	window.MakeContextCurrent()
	e.window = window

	e.logger.DebugContext(ctx, "glfw initialized")

	return nil
}

func (e *Engine) shutdownGLFW() {
	glfw.Terminate()

	e.logger.Debug("glfw shutdown complete")
}

func (e *Engine) keyCallback(
	_ *glfw.Window,
	key glfw.Key,
	_ int,
	action glfw.Action,
	_ glfw.ModifierKey,
) {
	e.keyStates[key] = action == glfw.Press
}

func (e *Engine) mouseCallback(
	_ *glfw.Window,
	button glfw.MouseButton,
	action glfw.Action,
	_ glfw.ModifierKey,
) {
	e.mouseStates[button] = action == glfw.Press
}

func (e *Engine) cursorPositionCallback(
	w *glfw.Window, x float64, y float64,
) {
	width, height := w.GetSize()
	if x < 0 || x > float64(width) {
		return
	}
	if y < 0 || y > float64(height) {
		return
	}

	e.mouseX = int(x / float64(width) * float64(e.textureW))
	e.mouseY = int(y / float64(height) * float64(e.textureH))
}
