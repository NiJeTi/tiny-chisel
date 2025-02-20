package engine

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func (e *Engine) initOpenGL(ctx context.Context) error {
	if err := gl.Init(); err != nil {
		return fmt.Errorf("failed to initialize opengl: %w", err)
	}

	vertexShader, err := e.compileShader(
		filepath.Join(shadersDir, vertexShaderFile), gl.VERTEX_SHADER,
	)
	if err != nil {
		return err
	}
	e.logger.DebugContext(ctx, "vertex shader compiled")

	fragmentShader, err := e.compileShader(
		filepath.Join(shadersDir, fragmentShaderFile), gl.FRAGMENT_SHADER,
	)
	if err != nil {
		return err
	}
	e.logger.DebugContext(ctx, "fragment shader compiled")

	progRef := gl.CreateProgram()
	gl.AttachShader(progRef, vertexShader)
	gl.AttachShader(progRef, fragmentShader)
	gl.LinkProgram(progRef)
	e.prog = progRef

	e.logger.DebugContext(ctx, "opengl initialized")

	return nil
}

func (e *Engine) shutdownOpenGL() {
	gl.DeleteProgram(e.prog)

	e.logger.Debug("opengl shutdown complete")
}

func (*Engine) compileShader(path string, shaderType uint32) (uint32, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("failed to read shader file: %w", err)
	}

	source, free := gl.Strs(string(bytes))
	defer free()

	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, 1, source, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := string(make([]byte, logLength))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile shader: %s", log)
	}

	return shader, nil
}
