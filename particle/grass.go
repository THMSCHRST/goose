package particle

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GrassBlade struct {
    Pos   rl.Vector3
    Scale float32
    Color rl.Color
}

func GenerateGrass(sizeX, sizeZ int, spacing float32) []GrassBlade {
    grass := make([]GrassBlade, 0)
    for x := -sizeX/2; x < sizeX/2; x++ {
        for z := -sizeZ/2; z < sizeZ/2; z++ {
            // Add some randomness to placement
            if rand.Float32() > 0.7 { // 30% coverage
                pos := rl.NewVector3(
                    float32(x)*spacing + (rand.Float32()-0.5)*spacing,
                    0.2, // ground level
                    float32(z)*spacing + (rand.Float32()-0.5)*spacing,
                )
                grass = append(grass, GrassBlade{
                    Pos:   pos,
                    Scale: 1.5 + rand.Float32()*0.5,
                    Color: rl.NewColor(255-uint8(rand.IntN(10)), 255-uint8(rand.IntN(10)), 255-uint8(rand.IntN(10)), 255), //92, 135, 22
                })
            }
        }
    }
    return grass
}