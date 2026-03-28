package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (p *Player) Render(debug bool) {
	velDir := rl.Vector2Angle(rl.NewVector2(0,0),rl.NewVector2(p.Vel.X,p.Vel.Z))

	normalVel := rl.Vector3Normalize(p.Vel)

	runOffset := math.Sin(float64(p.RunTime*2))*float64(rl.Vector3Length(rl.Vector3Multiply(p.Vel,normalVel)))

	//runOffsetHead := math.Sin(float64(p.RunTime*4))*float64(rl.Vector3Length(rl.Vector3Multiply(p.Vel,normalVel)))

	//debug
	if debug {
		rl.DrawCircle3D(rl.NewVector3(p.LegLimitR.X,0,p.LegLimitR.Y),p.LegDist,rl.NewVector3(1,0,0),90,rl.Green)
		rl.DrawCircle3D(rl.NewVector3(p.LegLimitL.X,0,p.LegLimitL.Y),p.LegDist,rl.NewVector3(1,0,0),90,rl.Green)
		rl.DrawCircle3D(p.LegL,0.1,rl.NewVector3(1,0,0),90,rl.Red)
		rl.DrawCircle3D(p.LegR,0.1,rl.NewVector3(1,0,0),90,rl.Red)
		rl.DrawCubeWires(rl.Vector3Add(p.Pos,rl.NewVector3(0,0.5,0)), 1, 1, 1, rl.Maroon)
		rl.DrawLine3D(rl.NewVector3(p.Pos.X,p.Pos.Y+1,p.Pos.Z),rl.Vector3Add(rl.NewVector3(p.Pos.X,p.Pos.Y+1,p.Pos.Z),rl.Vector3Scale(p.Vel,0.5)),rl.Blue)
		rl.DrawSphereWires(rl.Vector3Add(rl.NewVector3(p.Pos.X,p.Pos.Y+1,p.Pos.Z),rl.Vector3Scale(p.Vel,0.5)),0.3,4,4,rl.Blue)
		rl.DrawCubeWires(p.HeadPos,0.4,0.4,0.4,rl.Red)
	}

	//Body

	//feet
	rl.DrawSphere(p.RealLegL,0.2,rl.NewColor(230, 153, 67,255))
	rl.DrawSphere(p.RealLegR,0.2,rl.NewColor(230, 153, 67,255))

	//chest
	rl.DrawSphere(rl.NewVector3(p.Pos.X,p.Pos.Y+0.75+(float32(runOffset)*0.01),p.Pos.Z),0.5,rl.White)

	//after
	rl.DrawSphere(rl.NewVector3(p.Pos.X-((0.05*p.Vel.X)+normalVel.X*0.7),p.Pos.Y+0.85+(float32(runOffset)*0.01),p.Pos.Z-((0.05*p.Vel.Z)+normalVel.Z*0.7)),0.3,rl.White)

	rl.DrawCylinderEx(rl.NewVector3(p.Pos.X-((0.05*p.Vel.X)+normalVel.X*0.7),p.Pos.Y+0.85+(float32(runOffset)*0.01),p.Pos.Z-((0.05*p.Vel.Z)+normalVel.Z*0.7)), rl.NewVector3(p.Pos.X,p.Pos.Y+0.75+(float32(runOffset)*0.01),p.Pos.Z), 0.3, 0.5, 8, rl.White)

	rl.DrawSphere(rl.NewVector3(p.Pos.X-((0.05*p.Vel.X)+normalVel.X*0.3),p.Pos.Y+0.85+(float32(runOffset)*0.02),p.Pos.Z-((0.05*p.Vel.Z)+normalVel.Z*0.2)),0.35,rl.White)

	KneeL := rl.Vector2Add(rl.NewVector2(p.Pos.X,p.Pos.Z),rl.Vector2Rotate(rl.NewVector2(0,0.25),velDir))
	KneeR := rl.Vector2Add(rl.NewVector2(p.Pos.X,p.Pos.Z),rl.Vector2Rotate(rl.NewVector2(0,-0.25),velDir))

	//legs
	rl.DrawCylinderEx(p.RealLegL, rl.NewVector3(KneeL.X,p.Pos.Y+0.75,KneeL.Y), 0.1, 0.1, 8, rl.NewColor(235, 149, 52,255))
	rl.DrawCylinderEx(p.RealLegR, rl.NewVector3(KneeR.X,p.Pos.Y+0.75,KneeR.Y), 0.1, 0.1, 8, rl.NewColor(235, 149, 52,255))

	//neck
	rl.DrawCylinderEx(p.HeadPos, rl.NewVector3(p.Pos.X,p.Pos.Y+0.75,p.Pos.Z), 0.2, 0.2, 8, rl.White)
	//head
	rl.DrawSphere(p.HeadPos,0.2,rl.White)

	//eyes
	rl.DrawSphere(rl.Vector3Add(p.HeadPos, rl.NewVector3(rl.Vector2Rotate(rl.NewVector2(0,0.2),velDir).X, 0, rl.Vector2Rotate(rl.NewVector2(0,0.2),velDir).Y)),0.1,rl.Black)
	rl.DrawSphere(rl.Vector3Add(p.HeadPos, rl.NewVector3(rl.Vector2Rotate(rl.NewVector2(0,-0.2),velDir).X, 0, rl.Vector2Rotate(rl.NewVector2(0,-0.2),velDir).Y)),0.1,rl.Black)

	//beak
	rl.DrawSphere(rl.Vector3Add(p.HeadPos, rl.NewVector3(rl.Vector2Rotate(rl.NewVector2(0.3,0),velDir).X, 0, rl.Vector2Rotate(rl.NewVector2(0.3,0),velDir).Y)),0.15,rl.Orange)

}