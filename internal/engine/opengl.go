package engine

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func initOpenGL(ctx *ectx) (uint32, error) {
	if err := gl.Init(); err != nil {
		return 0, fmt.Errorf("failed to initialize opengl: %w", err)
	}

	vertexShader, err := compileShader(
		filepath.Join(shadersDir, vertexShaderFile), gl.VERTEX_SHADER,
	)
	if err != nil {
		return 0, err
	}
	ctx.logger.DebugContext(ctx, "vertex shader compiled")

	fragmentShader, err := compileShader(
		filepath.Join(shadersDir, fragmentShaderFile), gl.FRAGMENT_SHADER,
	)
	if err != nil {
		return 0, err
	}
	ctx.logger.DebugContext(ctx, "fragment shader compiled")

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	ctx.logger.DebugContext(ctx, "opengl initialized")

	return program, nil
}

func shutdownOpenGL(ctx *ectx, program uint32) {
	gl.DeleteProgram(program)

	ctx.logger.Debug("opengl shutdown complete")
}

func compileShader(path string, shaderType uint32) (uint32, error) {
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
