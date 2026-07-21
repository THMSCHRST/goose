package main

import (
	"fmt"
	"goose/gen"
	player "goose/player" // player
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
    rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "goose")
    defer rl.CloseWindow()

    camera := rl.Camera3D{
        Position:   rl.NewVector3(0, 0, 0),
        Target:     rl.NewVector3(0, 0, 1),
        Up:         rl.NewVector3(0, 1, 0),
        Fovy:       95.0,
        Projection: rl.CameraPerspective,
    }

    rl.SetTargetFPS(240)
    rl.HideCursor()
    rl.DisableCursor()

    yaw := -math.Pi / 4
    pitch := 0.0
    mouseSensitivity := 0.002

    p := player.NewPlayer(rl.NewVector3(0, 0, 0))

	var eyeHeight float32 = 1.0
	distance := 5.0

	debug := false

	trees := []gen.Tree{}

	for range 40 {
		trees = append(trees, gen.NewTree(rl.NewVector3(float32(rand.IntN(100)-50),0,float32(rand.IntN(100)-50))))
	}

	var grassFloor rl.Material = rl.LoadMaterialDefault()
	grassFloor.GetMap(0).Color = rl.NewColor(92, 135, 22, 255)

	floor := rl.LoadModelFromMesh(rl.GenMeshPlane(100,100,1,1))
	floor.Materials = &grassFloor

	//var dust []particle.Particle

    for !rl.WindowShouldClose() {
		/*
		if rl.Vector3Length(p.Vel) > 0.75 {
			if rand.Float32() < 4.5*rl.GetFrameTime() {
				dust = append(dust, particle.NewParticle(p.Pos,rl.Vector3Add(rl.Vector3Scale(p.Vel,-0.3),rl.NewVector3(0,3,0)),-10,3))
			}
		}*/

        mouseDelta := rl.GetMouseDelta()
        yaw += float64(mouseDelta.X) * mouseSensitivity
        pitch += float64(mouseDelta.Y) * mouseSensitivity

        if pitch > math.Pi/2-0.01 {
            pitch = math.Pi/2 - 0.01
        }
        if pitch < -math.Pi/2+0.01 {
            pitch = -math.Pi/2 + 0.01
        }

		distance -= float64(rl.GetMouseWheelMove())

		offsetX := float32(math.Cos(yaw) * math.Cos(pitch) * distance)
		offsetY := float32(math.Sin(pitch) * distance)
		offsetZ := float32(math.Sin(yaw) * math.Cos(pitch) * distance)

		camera.Position = rl.NewVector3(p.Pos.X+offsetX, p.Pos.Y+eyeHeight+offsetY, p.Pos.Z+offsetZ)
		camera.Target = rl.NewVector3(p.Pos.X, p.Pos.Y+eyeHeight, p.Pos.Z) // look at player

        p.MovementTick(camera)

        rl.BeginDrawing()
        rl.ClearBackground(rl.NewColor(175, 216, 219,255))

        rl.BeginMode3D(camera)
        p.Render(debug)
		for _,tree := range trees {
			tree.Render(debug)
		}

		if debug {rl.DrawGrid(100, 1.0)}
		rl.DrawModel(floor,rl.NewVector3(0,0,0),1,rl.White)

		/*
		newDust := []particle.Particle{}

		for _,p := range dust {
			p.Tick()
			rl.DrawSphere(p.Pos,0.4,rl.NewColor(40,40,40,50))
			if p.Alive {
				newDust = append(newDust, p)
			}
		}
		dust = newDust
		*/

		rl.EndMode3D()

        rl.DrawText("GAMEPLAY IS SUBJECT TO CHANGE", 10, 10, 20, rl.DarkGray)
		rl.DrawText(fmt.Sprint(rl.GetFPS()), 10, 30, 20, rl.DarkGray)

        rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyF3) {
			if debug {
				debug = false
			} else {
				debug = true
			}
		}
    }
}