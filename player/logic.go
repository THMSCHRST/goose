package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos rl.Vector3
	Vel rl.Vector3
	EnvVel rl.Vector3
	LegLimitL rl.Vector2
	LegLimitR rl.Vector2
	LegDist float32
	LegL rl.Vector3
	LegR rl.Vector3
	RealLegR rl.Vector3
	RealLegL rl.Vector3
	HeadPos rl.Vector3
	Step bool
	Stepped bool
	RunTime float32
	HeadState1 float32
	HeadState2 float32
	HeadState3 float32
}

func NewPlayer(pos rl.Vector3) (Player) {
	return Player{
		Pos: pos,
		Vel: rl.NewVector3(0.01,0,0.01),
		HeadState1: 0.05,
		HeadState2: 0.1,
		HeadState3: 1.75,
	}
}

func (p *Player) MovementTick(cam rl.Camera3D) {
	p.Stepped = false
	dt := rl.GetFrameTime()
	var speed float32 = 10
	var acceleration float32 = 1
	var friction float32 = 0.1

	if rl.Vector3Length(p.Vel) > 0.1{
		p.RunTime += dt
	} else {
		p.RunTime = 0
	}

	velDir := rl.Vector2Angle(rl.NewVector2(0,0),rl.NewVector2(p.Vel.X,p.Vel.Z))

	p.LegLimitL = rl.Vector2Add(rl.NewVector2(p.Pos.X,p.Pos.Z),rl.Vector2Rotate(rl.NewVector2(0,0.3),velDir))

	p.LegLimitR = rl.Vector2Add(rl.NewVector2(p.Pos.X,p.Pos.Z),rl.Vector2Rotate(rl.NewVector2(0,-0.3),velDir))

    forward := rl.Vector3Subtract(cam.Target, cam.Position)
    forward.Y = 0
    forward = rl.Vector3Normalize(forward)

    right := rl.Vector3CrossProduct(forward, cam.Up)
    right.Y = 0
    right = rl.Vector3Normalize(right)

    // 2. Build desired movement direction based on keys
    newVel := rl.NewVector3(0, 0, 0)

    if rl.IsKeyDown(rl.KeyW) {
        newVel = rl.Vector3Add(newVel, forward)
    }
    if rl.IsKeyDown(rl.KeyS) {
        newVel = rl.Vector3Subtract(newVel, forward)
    }
    if rl.IsKeyDown(rl.KeyD) {
        newVel = rl.Vector3Add(newVel, right)
    }
    if rl.IsKeyDown(rl.KeyA) {
        newVel = rl.Vector3Subtract(newVel, right)
    }

    if newVel.X != 0 || newVel.Y != 0 || newVel.Z != 0 {
        newVel = rl.Vector3Normalize(newVel)
    }

	desiredVel := rl.Vector3Scale(newVel, speed)

	p.Vel.X += (desiredVel.X - p.Vel.X) * acceleration * dt
	p.Vel.Y += (desiredVel.Y - p.Vel.Y) * acceleration * dt
	p.Vel.Z += (desiredVel.Z - p.Vel.Z) * acceleration * dt

	p.Vel = rl.Vector3Scale(p.Vel, 1 - friction*dt)

	p.Pos.X += p.Vel.X * dt
	p.Pos.Y += p.Vel.Y * dt
	p.Pos.Z += p.Vel.Z * dt


	if rl.IsKeyPressed(rl.KeyE) {
		p.EnvVel.X += 5
	}

	p.EnvVel = rl.Vector3Subtract(p.EnvVel,rl.NewVector3(p.EnvVel.X*1.5*dt,p.EnvVel.Y*1.5*dt,p.EnvVel.Z*1.5*dt))

	p.Pos.X += p.EnvVel.X * dt
	p.Pos.Y += p.EnvVel.Y * dt
	p.Pos.Z += p.EnvVel.Z * dt

	if rl.Vector3Length(p.EnvVel) < 0.05 {p.EnvVel = rl.NewVector3(0,0,0)}


	p.LegDist = max(rl.Vector2Length(rl.NewVector2(p.Vel.X,p.Vel.Z))/10,0.5)

	lNeedsStep := rl.Vector2Distance(p.LegLimitL,rl.NewVector2(p.LegL.X,p.LegL.Z)) > p.LegDist || rl.Vector2Length(rl.NewVector2(p.Vel.X,p.Vel.Z)) < 0.25
	rNeedsStep := rl.Vector2Distance(p.LegLimitR,rl.NewVector2(p.LegR.X,p.LegR.Z)) > p.LegDist || rl.Vector2Length(rl.NewVector2(p.Vel.X,p.Vel.Z)) < 0.25
	rNeedsStepUrgent :=  rl.Vector2Distance(p.LegLimitR,rl.NewVector2(p.LegR.X,p.LegR.Z)) > p.LegDist*1.5 || rl.Vector2Length(rl.NewVector2(p.Vel.X,p.Vel.Z)) < 0.25
	lNeedsStepUrgent :=  rl.Vector2Distance(p.LegLimitL,rl.NewVector2(p.LegL.X,p.LegL.Z)) > p.LegDist*1.5 || rl.Vector2Length(rl.NewVector2(p.Vel.X,p.Vel.Z)) < 0.25

	if lNeedsStep && rNeedsStep && p.Step == false && p.Stepped == false || lNeedsStepUrgent {
		p.LegL = rl.Vector3Add(rl.NewVector3(p.LegLimitL.X,p.Pos.Y,p.LegLimitL.Y),rl.NewVector3(float32(math.Cos(float64(velDir))*float64(p.LegDist)),p.Pos.Y,float32(math.Sin(float64(velDir))*float64(p.LegDist))))
		p.Step = true
		p.Stepped = true
	}

	if rNeedsStep && lNeedsStep && p.Step && p.Stepped == false || rNeedsStepUrgent {
		p.LegR = rl.Vector3Add(rl.NewVector3(p.LegLimitR.X,p.Pos.Y,p.LegLimitR.Y),rl.NewVector3(float32(math.Cos(float64(velDir))*float64(p.LegDist)),p.Pos.Y,float32(math.Sin(float64(velDir))*float64(p.LegDist))))
		p.Step = false
		p.Stepped = true
	}

	if rl.Vector3Length(p.Vel) < 0.3 {
		p.LegL = rl.Vector3Add(rl.NewVector3(p.LegLimitL.X,p.Pos.Y,p.LegLimitL.Y),rl.Vector3Scale(p.Vel,1.0/10))
		p.LegR = rl.Vector3Add(rl.NewVector3(p.LegLimitR.X,p.Pos.Y,p.LegLimitR.Y),rl.Vector3Scale(p.Vel,1.0/10))
	}

	p.RealLegL = rl.Vector3Add(p.RealLegL,rl.Vector3Scale(rl.Vector3Subtract(p.LegL,p.RealLegL),30*dt))
	p.RealLegL.Y = p.Pos.Y+rl.Vector3Distance(p.LegL,p.RealLegL)/12

	p.RealLegR = rl.Vector3Add(p.RealLegR,rl.Vector3Scale(rl.Vector3Subtract(p.LegR,p.RealLegR),30*dt))
	p.RealLegR.Y = p.Pos.Y+rl.Vector3Distance(p.LegR,p.RealLegR)/12

	normalVel := rl.Vector3Normalize(p.Vel)

	runOffsetHead := math.Sin(float64(p.RunTime*4))*float64(rl.Vector3Length(rl.Vector3Multiply(p.Vel,normalVel)))

	if rl.IsKeyDown(rl.KeyLeftControl) {
		p.HeadPos = rl.Vector3Add(p.Pos,rl.Vector3Add(rl.Vector3Scale(p.Vel,p.HeadState1),rl.Vector3Scale(normalVel,p.HeadState2)))
		p.HeadPos.Y += p.HeadState3 + float32(runOffsetHead*0.01)
		p.HeadState1 += (0.01-p.HeadState1)*7*dt
		p.HeadState2 += (1.2-p.HeadState2)*7*dt
		p.HeadState3 += (1.1-p.HeadState3)*7*dt
	} else {
		p.HeadPos = rl.Vector3Add(p.Pos,rl.Vector3Add(rl.Vector3Scale(p.Vel,p.HeadState1),rl.Vector3Scale(normalVel,p.HeadState2)))
		p.HeadPos.Y += p.HeadState3 + float32(runOffsetHead*0.01)
		p.HeadState1 += (0.05-p.HeadState1)*7*dt
		p.HeadState2 += (0.1-p.HeadState2)*7*dt
		p.HeadState3 += (1.75-p.HeadState3)*7*dt
	}
}