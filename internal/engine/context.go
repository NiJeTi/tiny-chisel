package engine

import (
	"context"
	colorPkg "image/color"
	"log/slog"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type Context interface {
	context.Context

	Delta() time.Duration

	KeyState(key glfw.Key) bool

	MouseButtonState(button glfw.MouseButton) bool
	MousePos() (x int, y int)

	SpaceSize() (width, height int)
	SetPixel(x, y int, color colorPkg.RGBA)
}

type ectx struct {
	context.Context

	logger *slog.Logger

	windowTitle      string
	windowW, windowH int
	windowResizable  bool

	spaceW, spaceH int
	textureData    []byte

	keys           map[glfw.Key]bool
	mouseButtons   map[glfw.MouseButton]bool
	mouseX, mouseY int

	controllers []Controller
}

func newCtx(ctx context.Context) *ectx {
	return &ectx{
		Context:         ctx,
		logger:          slog.New(slog.NewTextHandler(nil, nil)),
		windowTitle:     defaultWindowTitle,
		windowW:         defaultWindowW,
		windowH:         defaultWindowH,
		windowResizable: defaultWindowResizable,
		spaceW:          defaultWindowW / 2,
		spaceH:          defaultWindowH / 2,
	}
}

func (ctx *ectx) Init() {
	ctx.textureData = make([]byte, ctx.spaceW*ctx.spaceH*sizeColor)
	ctx.keys = make(map[glfw.Key]bool)
	ctx.mouseButtons = make(map[glfw.MouseButton]bool)
}

func (ctx *ectx) Delta() time.Duration {
	return frameDelta
}

func (ctx *ectx) KeyState(key glfw.Key) bool {
	state, ok := ctx.keys[key]
	return ok && state
}

func (ctx *ectx) MouseButtonState(button glfw.MouseButton) bool {
	state, ok := ctx.mouseButtons[button]
	return ok && state
}

func (ctx *ectx) MousePos() (x int, y int) {
	return ctx.mouseX, ctx.mouseY
}

func (ctx *ectx) SpaceSize() (width, height int) {
	return ctx.spaceW, ctx.spaceH
}

func (ctx *ectx) SetPixel(x, y int, color colorPkg.RGBA) {
	if x >= ctx.spaceW || x < 0 || y >= ctx.spaceH || y < 0 {
		panic("failed to set pixel out of bounds")
	}

	index := (x + (ctx.spaceH-1-y)*ctx.spaceW) * sizeColor

	const (
		rOffset = iota
		gOffset
		bOffset
		aOffset
	)

	ctx.textureData[index+rOffset] = color.R
	ctx.textureData[index+gOffset] = color.G
	ctx.textureData[index+bOffset] = color.B
	ctx.textureData[index+aOffset] = color.A
}
