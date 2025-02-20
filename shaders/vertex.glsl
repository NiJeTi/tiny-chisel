#version 410 core

layout (location = 0) in vec2 position;
layout (location = 1) in vec2 gridTextureCoordinates;

out vec2 gridTexturePos;

void main() {
    gridTexturePos = gridTextureCoordinates;
    gl_Position = vec4(position, 0, 1);
}
