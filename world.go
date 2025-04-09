package main

import (
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	worldSize = 1
)

type World struct {
	Chunks []Chunk
}

func CreateWorld(perlin *perlin.Perlin) World {
	var chunks = make([]Chunk, worldSize*worldSize)

	for i := 0; i < worldSize; i++ {
		for j := 0; j < worldSize; j++ {
			chunks[i*worldSize+j] = CreateChunk(perlin, rl.NewVector2(float32(i), float32(j)))
		}
	}

	return World{chunks}
}
