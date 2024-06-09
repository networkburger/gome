package component

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BallisticComponent struct {
	Gravity         rl.Vector2
	Velocity        rl.Vector2
	VelocityDamping rl.Vector2
	AngularVelocity en.AngleD
	AngularDamping  float32
	Impulse         rl.Vector2
	Torque          float32
}

func (s *BallisticComponent) Event(e en.NodeEvent, n *en.Node) {}

func (b *BallisticComponent) Tick(gs *en.GameState, n *en.Node) {
	accel := rl.Vector2Add(b.Gravity, b.Impulse)
	b.Velocity = rl.Vector2Add(b.Velocity, rl.Vector2Scale(accel, gs.DT))
	b.Velocity = rl.Vector2Subtract(b.Velocity,
		rl.Vector2Multiply(b.Velocity, rl.Vector2Scale(b.VelocityDamping, gs.DT)))

	avel := float32(b.AngularVelocity) + b.Torque*gs.DT
	avel -= avel * b.AngularDamping * gs.DT
	b.AngularVelocity = en.AngleD(avel)

	b.Impulse = rl.NewVector2(0, 0)
	b.Torque = 0

	n.Position = rl.Vector2Add(n.Position, rl.Vector2Scale(b.Velocity, gs.DT))
	n.Rotation += en.AngleD(avel * gs.DT)
}
