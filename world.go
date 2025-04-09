package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	worldSize = 2
)

type World struct {
	Chunks []Chunk
}

func CreateWorld() World {
	var chunks = make([]Chunk, worldSize*worldSize)

	for i := 0; i < worldSize; i++ {
		for j := 0; j < worldSize; j++ {
			chunks[i*worldSize+j] = CreateChunk(rl.NewVector2(float32(i), float32(j)))
		}
	}

	return World{chunks}
}
