package main

import (
	"blocks/generation"
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ChunkSize   = 16
	ChunkHeight = 128
)

type Chunk struct {
	Position rl.Vector2
	Blocks   []Block
}

func CreateChunk(perlin *perlin.Perlin, position rl.Vector2) Chunk {
	var blocks = make([]Block, ChunkSize*ChunkHeight*ChunkSize)

	for i := 0; i < ChunkSize; i++ {
		for j := 0; j < ChunkSize; j++ {
			translateX := position.X * ChunkSize
			translateZ := position.Y * ChunkSize

			absolutePosX := float32(i) + translateX
			absolutePosZ := float32(j) + translateZ

			blockHeight := generation.GetHeight(perlin, absolutePosX, absolutePosZ)
			blockPosition := rl.NewVector3(float32(i), blockHeight, float32(j))

			for k := 0; k < ChunkHeight && float32(k) < blockHeight; k++ {
				index := i + ChunkSize*(j+ChunkHeight*k)
				blocks[index] = Block{
					Color:    ColorFromPosition(blockPosition),
					Position: blockPosition,
					Size:     1.0,
				}
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
