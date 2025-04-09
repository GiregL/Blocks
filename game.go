package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameState struct {
	Player Player
	World  World
}

func CreateGameState() GameState {
	return GameState{
		Player: createPlayer(),
		World:  CreateWorld(),
	}
}

func createPlayer() Player {
	player := Player{
		Position: rl.NewVector3(0, 1, 0),
		Velocity: 1.0,
	}

	camera := rl.Camera3D{}
	camera.Position = player.Position
	camera.Target = rl.NewVector3(player.Position.X, player.Position.Y, player.Position.Z+1)
	camera.Up = rl.NewVector3(0, 1, 0)
	camera.Fovy = 90.0
	camera.Projection = rl.CameraPerspective

	player.Camera = camera

	return player
}

func (gs *GameState) Update() {
	// Input update
	rl.UpdateCamera(&gs.Player.Camera, rl.CameraFirstPerson)
	gs.Player.Position = gs.Player.Camera.Position
	defer gs.Player.SyncCameraPos()

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

	rl.DrawRectangle(5, 5, 330, 100, rl.Fade(rl.SkyBlue, 0.5))
	rl.DrawRectangleLines(5, 5, 330, 100, rl.Blue)

	rl.DrawText("Camera controls:", 15, 15, 10, rl.Black)
	rl.DrawText("- Move keys: W, A, S, D, Space, Left-Ctrl", 15, 30, 10, rl.Black)
	rl.DrawText("- Look around: arrow keys or mouse", 15, 45, 10, rl.Black)
	rl.DrawText("- Camera mode keys: 1,2,3,4", 15, 60, 10, rl.Black)
	rl.DrawText("- Zoom keys: num-plus, num-minus or mouse scroll", 15, 75, 10, rl.Black)

	rl.EndDrawing()
}
