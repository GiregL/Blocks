package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Position rl.Vector3
	Velocity float32
	Camera   rl.Camera
}

func (player *Player) SyncCameraPos() {
	player.Camera.Position = player.Position
}

func (player *Player) Update() {

	if rl.IsKeyDown(rl.KeySpace) {
		player.Position.Y += player.Velocity * 0.2
		player.Camera.Target.Y += player.Velocity * 0.2
	}

	if rl.IsKeyDown(rl.KeyLeftShift) {
		player.Position.Y -= player.Velocity * 0.5
		player.Camera.Target.Y -= player.Velocity * 0.5
	}

	player.SyncCameraPos()
}
