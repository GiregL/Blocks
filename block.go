package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Color    rl.Color
	Position rl.Vector3
	Size     float32
}

func ColorFromPosition(position rl.Vector3) rl.Color {
	r := uint8(position.X * 255.0)
	g := uint8(position.Y * 255.0)
	b := uint8(position.Z * 255.0)
	a := uint8(255)

	return rl.NewColor(r, g, b, a)
}
