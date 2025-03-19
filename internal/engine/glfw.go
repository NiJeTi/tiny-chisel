package engine

import (
	"fmt"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	glVersionMajor = 4
	glVersionMinor = 1
)

const (
	shadersDir = "shaders/"
	shadersExt = ".glsl"

	vertexShaderFile   = "vertex" + shadersExt
	fragmentShaderFile = "fragment" + shadersExt
)

func initGLFW(ctx *ectx) (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, fmt.Errorf("failed to init glfw: %w", err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, glVersionMajor)
	glfw.WindowHint(glfw.ContextVersionMinor, glVersionMinor)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	resizable := glfw.False
	if ctx.windowResizable {
		resizable = glfw.True
	}
	glfw.WindowHint(glfw.Resizable, resizable)

	window, err := glfw.CreateWindow(
		ctx.windowW, ctx.windowH, ctx.windowTitle, nil, nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %w", err)
	}

	window.SetKeyCallback(keyCallback(ctx))
	window.SetMouseButtonCallback(mouseCallback(ctx))
	window.SetCursorPosCallback(cursorPositionCallback(ctx))

	window.MakeContextCurrent()

	ctx.logger.DebugContext(ctx, "glfw initialized")

	return window, nil
}

func shutdownGLFW(ctx *ectx) {
	glfw.Terminate()

	ctx.logger.Debug("glfw shutdown complete")
}

func keyCallback(ctx *ectx) glfw.KeyCallback {
	return func(
		_ *glfw.Window,
		key glfw.Key,
		_ int,
		action glfw.Action,
		_ glfw.ModifierKey,
	) {
		ctx.keys[key] = action == glfw.Press
	}
}

func mouseCallback(ctx *ectx) glfw.MouseButtonCallback {
	return func(
		_ *glfw.Window,
		button glfw.MouseButton,
		action glfw.Action,
		_ glfw.ModifierKey,
	) {
		ctx.mouseButtons[button] = action == glfw.Press
	}
}

func cursorPositionCallback(ctx *ectx) glfw.CursorPosCallback {
	return func(w *glfw.Window, x float64, y float64) {
		width, height := w.GetSize()

		if x < 0 || x > float64(width) {
			return
		}
		if y < 0 || y > float64(height) {
			return
		}

		ctx.mouseX = int(x / float64(width) * float64(ctx.spaceW))
		ctx.mouseY = int(y / float64(height) * float64(ctx.spaceH))
	}
}
