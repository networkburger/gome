package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerStats struct {
	Speed        float32
	ReverseSpeed float32
	TurnSpeed    float32
	MaxHealth    int
}

func StartingPlayerStats() PlayerStats {
	return PlayerStats{
		Speed:     1200,
		TurnSpeed: 250,
		MaxHealth: 100,
	}
}

type Player struct {
	Stats      PlayerStats
	Health     int
	Ballistics *physics.BallisticComponent
}

func (p *Player) Event(event engine.NodeEvent, gs *engine.Scene, node *engine.Node) {
	if event == engine.NodeEventTick {
		engine.ProcessInputs(InputOverworld, func(action engine.ActionID, power float32) {
			switch action {
			case Accelerate:
				p.Ballistics.Impulse = node.Forward().Scl(power * p.Stats.Speed)
			case TurnLeft:
				p.Ballistics.Torque = p.Stats.TurnSpeed * -power
			case TurnRight:
				p.Ballistics.Torque = p.Stats.TurnSpeed * power
			case Turn:
				p.Ballistics.Torque = p.Stats.TurnSpeed * power
			}
		})
	}
}

func StandardPlayerNode(e *engine.Engine, phys *physics.PhysicsSolver) *engine.Node {
	r := float32(20)
	player := Player{
		Stats:  StartingPlayerStats(),
		Health: 100,
	}
	playerNode := e.NewNode("Player")
	playerNode.AddComponent(&player)

	a45 := math.Pi / 4

	polyNode := NewLineStripComponent(rl.White, []v.Vec2{
		v.V2(float32(math.Sin(a45*0))*r, float32(math.Cos(a45*0))*r),
		v.V2(float32(math.Sin(a45*1))*r, float32(math.Cos(a45*1))*r),
		v.V2(float32(math.Sin(a45*2))*r, float32(math.Cos(a45*2))*r),
		v.V2(float32(math.Sin(a45*3))*r, float32(math.Cos(a45*3))*r),
		v.V2(float32(math.Sin(a45*4))*r, float32(math.Cos(a45*4))*r),
		v.V2(float32(math.Sin(a45*5))*r, float32(math.Cos(a45*5))*r),
		v.V2(float32(math.Sin(a45*6))*r, float32(math.Cos(a45*6))*r),
		v.V2(float32(math.Sin(a45*7))*r, float32(math.Cos(a45*7))*r),
		v.V2(float32(math.Sin(a45*0))*r, float32(math.Cos(a45*0))*r),
		v.V2(0, 0),
	})
	playerNode.AddComponent(&polyNode)

	ballistics := physics.BallisticComponent{
		VelocityDamping: v.V2(4, 4),
		AngularDamping:  5,
		Gravity:         v.V2(0, 1000),
	}
	playerNode.AddComponent(&ballistics)
	player.Ballistics = &ballistics

	physBody := physics.PhysicsBodyComponent{
		PhysicsSolver: phys,
		Radius:        20,
		SurfaceProperties: physics.SurfaceProperties{
			Friction:    0,
			Restitution: 0.2,
		},
	}
	playerNode.AddComponent(&physBody)

	return playerNode
}
