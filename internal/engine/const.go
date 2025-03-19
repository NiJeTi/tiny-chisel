package engine

import (
	"time"
)

const (
	defaultWindowTitle     = "tiny-chisel"
	defaultWindowW         = 1280
	defaultWindowH         = 720
	defaultWindowResizable = true
)

const (
	frameRate  = 60
	frameDelta = time.Second / frameRate
)

const (
	sizeFloat32 = 4
	sizeColor   = 4
)

const (
	quadVerticesCount = 3 + 3
	vertexPosSize     = 2
	texturePosSize    = 2
	vertexInfoSize    = vertexPosSize + texturePosSize
)
