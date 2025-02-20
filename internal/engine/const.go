package engine

const (
	windowTitle  = "graphics"
	windowWidth  = 1280
	windowHeight = 720
)

const (
	glVersionMajor = 4
	glVersionMinor = 1
)

const (
	shadersDir = "shaders/"
	shadersExt = ".glsl"

	vertexShaderFile   = "vertex.glsl"
	fragmentShaderFile = "fragment.glsl"
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

const (
	pixelSize     = 4
	textureWidth  = windowWidth / pixelSize
	textureHeight = windowHeight / pixelSize
)
