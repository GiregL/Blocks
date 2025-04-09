package main

import (
	"blocks/generation"
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	Player    Player
	World     World
	Generator *perlin.Perlin
}

func CreateGameState() GameState {
	gameState := GameState{
		Player:    createPlayer(),
		Generator: generation.CreatePerlin(),
	}
	gameState.World = CreateWorld(gameState.Generator)
	return gameState
}

func createPlayer() Player {
	player := Player{
		Velocity: 1.0,
	}

	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0, 1, 0)
	camera.Target = rl.NewVector3(player.Position().X, player.Position().Y, player.Position().Z+1)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 90.0
	camera.Projection = rl.CameraPerspective

	player.Camera = camera

	return player
}

func (gs *GameState) Update() {
	// Input update
	rl.UpdateCamera(&gs.Player.Camera, rl.CameraFirstPerson)
	gs.Player.Update()
}

func (gs *GameState) Render() {

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	// 3D Rendering
	rl.BeginMode3D(gs.Player.Camera)

	for _, chunk := range gs.World.Chunks {
		chunk.Render()
	}

	rl.EndMode3D()

	// HUD Rendering

	rl.EndDrawing()
}
