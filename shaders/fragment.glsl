#version 410 core

layout (location = 0) in vec2 gridTexturePos;

out vec4 color;

uniform sampler2D gridTexture;

void main() {
    color = texture(gridTexture, gridTexturePos);
}
