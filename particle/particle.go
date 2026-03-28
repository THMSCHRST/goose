package particle

import rl "github.com/gen2brain/raylib-go/raylib"

type Particle struct {
	Pos rl.Vector3
	Vel rl.Vector3
	Grav float64
	LifeTime float64
	Alive bool
}

func NewParticle(p, v rl.Vector3, g, l float64) Particle {return Particle{Alive: true,Pos: p,Vel: v,Grav: g,LifeTime: l}}

func (p *Particle) Tick() {
	if p.Alive {
		dt := rl.GetFrameTime()
		p.LifeTime -= float64(dt)

		p.Pos = rl.Vector3Add(p.Pos,rl.Vector3Scale(p.Vel,dt))
		p.Vel.Y += float32(p.Grav)*dt
		if p.LifeTime <= 0 {p.Alive=false}
	}
}