package generation

import (
	"fmt"
	"github.com/aquilax/go-perlin"
	"math/rand"
)

func CreatePerlin() *perlin.Perlin {
	return perlin.NewPerlinRandSource(4, 4.0, 2, rand.NewSource(rand.Int63()))
}

func GetHeight(perlin *perlin.Perlin, x, y float32) float32 {
	asFloat := float32(perlin.Noise2D(float64(x)/20.0, float64(y)/20.0)*10.0) + 3.0
	fmt.Println("Generating Height for coordinates : X = ", x, "\tY = ", y, "\t\t\tGENERATED => ", asFloat)
	// asFloat := rand.Float32() * 32
	return float32(int32(asFloat))
}
