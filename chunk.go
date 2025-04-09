package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ChunkSize   = 4
	ChunkHeight = 1
)

type Chunk struct {
	Position rl.Vector2
	Blocks   []Block
}

func CreateChunk(position rl.Vector2) Chunk {
	var blocks = make([]Block, ChunkSize*ChunkHeight*ChunkSize)

	for i := 0; i < ChunkSize; i++ {
		for j := 0; j < ChunkSize; j++ {
			blocks[(i*ChunkSize)+j] = Block{
				Color:    ColorFromPosition(rl.NewVector3(float32(i), 0, float32(j))),
				Position: rl.NewVector3(float32(i), 0, float32(j)),
				Size:     1.0,
			}
		}
	}

	return Chunk{position, blocks}
}

func (c *Chunk) Render() {
	translateX := c.Position.X * ChunkSize
	translateZ := c.Position.Y * ChunkSize

	for _, block := range c.Blocks {
		drawPosition := rl.NewVector3(float32(block.Position.X)+translateX, float32(block.Position.Y), float32(block.Position.Z)+translateZ)
		rl.DrawCube(drawPosition, block.Size, block.Size, block.Size, block.Color)
	}
}
