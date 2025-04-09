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

	if position.Y >= 15 {
		return rl.LightGray
	} else if position.Y >= 10 {
		return rl.Gray
	} else if position.Y >= 5 {
		return rl.DarkGreen
	} else if position.Y >= 0 {
		return rl.DarkGray
	} else {
		return rl.DarkBlue
	}

	//r := uint8(position.X * 255.0)
	//g := uint8(position.Y * 255.0)
	//b := uint8(position.Z * 255.0)
	//a := uint8(255)
	//
	//return rl.NewColor(r, g, b, a)
}
