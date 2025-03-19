package engine

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func initRender(ctx *ectx, program uint32) (
	vao uint32, vbo uint32, texture uint32, err error,
) {
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

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
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

	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.UseProgram(program)

	ctx.logger.DebugContext(ctx, "render initialized")

	return
}

func shutdownRender(ctx *ectx, vao uint32, vbo uint32, texture uint32) {
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.DeleteTextures(1, &texture)

	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.DeleteBuffers(1, &vbo)

	gl.BindVertexArray(0)
	gl.DeleteVertexArrays(1, &vao)

	ctx.logger.Debug("render shutdown complete")
}

func render(ctx *ectx, window *glfw.Window) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(ctx.spaceW),
		int32(ctx.spaceH),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(ctx.textureData),
	)
	gl.DrawArrays(gl.TRIANGLES, 0, quadVerticesCount)

	if err := gl.GetError(); err != gl.NO_ERROR {
		errCode := fmt.Errorf("%d (0x%x)", err, err)
		ctx.logger.ErrorContext(ctx, "opengl error", "error", errCode)
	}

	window.SwapBuffers()
}
