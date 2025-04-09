package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	maxColumns   = 20
	screenWidth  = 1080
	screenHeight = 720
	False        = 0
	True         = 1
)

var camModes = map[rl.CameraMode]string{
	rl.CameraFree:        "FREE",
	rl.CameraOrbital:     "ORBITAL",
	rl.CameraFirstPerson: "FIRST_PERSON",
	rl.CameraThirdPerson: "THIRD_PERSON",
}

var camProjections = map[rl.CameraProjection]string{
	rl.CameraPerspective:  "PERSPECTIVE",
	rl.CameraOrthographic: "ORTHOGRAPHIC",
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Blocks")
	defer rl.CloseWindow()

	gameState := CreateGameState()

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		gameState.Update()
		gameState.Render()
	}
}

// rndU generates a random uint8 value between min and max
func rndU(min, max int32) uint8 {
	return uint8(rl.GetRandomValue(min, max))
}

// rndF generates a random float32 value between min and max
func rndF(min, max int32) float32 {
	return float32(rl.GetRandomValue(min, max))
}

// s2t generates a Status item string from a name and value
func s2T(name, value string) string {
	return fmt.Sprintf(" - %s %s", name, value)
}

// v2T generates a string from a rl.Vector3
func v2T(v rl.Vector3) string {
	return fmt.Sprintf("%6.3f, %6.3f, %6.3f", v.X, v.Y, v.Z)
}
