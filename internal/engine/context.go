package engine

import (
	"context"
	colorPkg "image/color"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Context interface {
	context.Context

	KeyState(key glfw.Key) bool

	MouseButtonState(button glfw.MouseButton) bool
	MousePos() (x int, y int)

	SpaceSize() (width, height int)
	SetPixel(x, y int, color colorPkg.RGBA)
}

type input struct {
	keys           *map[glfw.Key]bool
	mouseButtons   *map[glfw.MouseButton]bool
	mouseX, mouseY *int
}

type engineCtx struct {
	context.Context
	input input

	textureW, textureH *int
	textureData        *[]byte
}

func (c *engineCtx) KeyState(key glfw.Key) bool {
	state, ok := (*c.input.keys)[key]
	return ok && state
}

func (c *engineCtx) MouseButtonState(button glfw.MouseButton) bool {
	state, ok := (*c.input.mouseButtons)[button]
	return ok && state
}

func (c *engineCtx) MousePos() (x int, y int) {
	return *c.input.mouseX, *c.input.mouseY
}

func (c *engineCtx) SpaceSize() (width, height int) {
	return *c.textureW, *c.textureH
}

func (c *engineCtx) SetPixel(x, y int, color colorPkg.RGBA) {
	w := *c.textureW
	h := *c.textureH

	if x >= w || x < 0 || y >= h || y < 0 {
		panic("failed to set pixel out of bounds")
	}

	index := (x + (h-1-y)*w) * sizeColor

	const (
		rOffset = iota
		gOffset
		bOffset
		aOffset
	)

	(*c.textureData)[index+rOffset] = color.R
	(*c.textureData)[index+gOffset] = color.G
	(*c.textureData)[index+bOffset] = color.B
	(*c.textureData)[index+aOffset] = color.A
}

func (e *Engine) initContext(ctx context.Context) *engineCtx {
	ectx := &engineCtx{
		Context: ctx,
		input: input{
			keys:         &e.keyStates,
			mouseButtons: &e.mouseStates,
			mouseX:       &e.mouseX,
			mouseY:       &e.mouseY,
		},
		textureW:    &e.textureW,
		textureH:    &e.textureH,
		textureData: &e.textureData,
	}
	return ectx
}
