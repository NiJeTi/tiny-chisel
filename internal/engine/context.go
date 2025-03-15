package engine

import (
	"context"

	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/nijeti/graphics/internal/types"
)

type Context interface {
	context.Context

	KeyState(key glfw.Key) bool

	MouseButtonState(button glfw.MouseButton) bool
	MousePos() (x int, y int)

	SpaceSize() (width, height int)
	SetPixel(x, y int, color types.Color)
}

type input struct {
	keys           map[glfw.Key]bool
	mouseButtons   map[glfw.MouseButton]bool
	mouseX, mouseY int
}

type eCtx struct {
	context.Context
	input input

	textureW, textureH int
	texture            [][]types.Color
}

func (c *eCtx) KeyState(key glfw.Key) bool {
	state, ok := c.input.keys[key]
	return ok && state
}

func (c *eCtx) MouseButtonState(button glfw.MouseButton) bool {
	state, ok := c.input.mouseButtons[button]
	return ok && state
}

func (c *eCtx) MousePos() (x int, y int) {
	return c.input.mouseX, c.input.mouseY
}

func (c *eCtx) SpaceSize() (width, height int) {
	return c.textureW, c.textureH
}

func (c *eCtx) SetPixel(x, y int, color types.Color) {
	if x > len(c.texture) || y > len(c.texture[0]) {
		panic("failed to set pixel out of bounds")
	}

	c.texture[x][y] = color
}

func (e *Engine) initContext(ctx context.Context) *eCtx {
	texture := make([][]types.Color, 0, e.textureW)
	for range e.textureW {
		pixels := make([]types.Color, 0, e.textureH)
		for range e.textureH {
			pixels = append(pixels, types.ColorBlack())
		}
		texture = append(texture, pixels)
	}
	ectx := &eCtx{
		Context: ctx,
		input: input{
			keys:         make(map[glfw.Key]bool),
			mouseButtons: make(map[glfw.MouseButton]bool),
			mouseX:       e.mouseX,
			mouseY:       e.mouseY,
		},
		textureW: e.textureW,
		textureH: e.textureH,
		texture:  texture,
	}
	return ectx
}
