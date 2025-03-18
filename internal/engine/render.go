package engine

import (
	"context"
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func (e *Engine) initRender(ctx context.Context, width, height int) error {
	quadVertices := []float32{
		// 1st triangle
		-1, 1, // top-left
		0, 1,

		-1, -1, // bottom-left
		0, 0,

		1, -1, // bottom-right
		1, 0,

		// 2nd triangle
		1, -1, // bottom-right
		1, 0,

		1, 1, // top-right
		1, 1,

		-1, 1, // top-left
		0, 1,
	}

	gl.GenVertexArrays(1, &e.vao)
	gl.BindVertexArray(e.vao)

	gl.GenBuffers(1, &e.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, e.vbo)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(quadVertices)*sizeFloat32,
		gl.Ptr(quadVertices),
		gl.STATIC_DRAW,
	)

	gl.VertexAttribPointerWithOffset(
		0,
		vertexPosSize,
		gl.FLOAT,
		false,
		vertexInfoSize*sizeFloat32,
		0,
	)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(
		1,
		texturePosSize,
		gl.FLOAT,
		false,
		vertexInfoSize*sizeFloat32,
		vertexPosSize*sizeFloat32,
	)
	gl.EnableVertexAttribArray(1)

	gl.GenTextures(1, &e.texture)
	gl.BindTexture(gl.TEXTURE_2D, e.texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.UseProgram(e.prog)

	e.textureW, e.textureH = width, height
	e.textureData = make([]byte, width*height*sizeColor)

	e.logger.DebugContext(ctx, "render initialized")

	return nil
}

func (e *Engine) shutdownRender() {
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.DeleteTextures(1, &e.texture)

	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.DeleteBuffers(1, &e.vbo)

	gl.BindVertexArray(0)
	gl.DeleteVertexArrays(1, &e.vao)

	e.logger.Debug("render shutdown complete")
}

func (e *Engine) render(ctx *engineCtx) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(e.textureW),
		int32(e.textureH),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(e.textureData),
	)
	gl.DrawArrays(gl.TRIANGLES, 0, quadVerticesCount)

	if err := gl.GetError(); err != gl.NO_ERROR {
		errCode := fmt.Errorf("%d (0x%x)", err, err)
		e.logger.ErrorContext(ctx, "opengl error", "error", errCode)
	}

	e.window.SwapBuffers()
}
