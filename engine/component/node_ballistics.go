package component

import (
	en "jamesraine/grl/engine"
	"jamesraine/grl/engine/v"
)

type BallisticComponent struct {
	Gravity         v.Vec2
	Velocity        v.Vec2
	VelocityDamping v.Vec2
	AngularVelocity en.AngleD
	AngularDamping  float32
	Impulse         v.Vec2
	Torque          float32
}

func (s *BallisticComponent) Event(e en.NodeEvent, n *en.Node) {}

func (b *BallisticComponent) Tick(gs *en.GameState, n *en.Node) {
	accel := b.Gravity.Add(b.Impulse)
	b.Velocity = b.Velocity.Add(accel.Scl(gs.DT))
	b.Velocity = b.Velocity.Sub(b.Velocity.Mul(b.VelocityDamping.Scl(gs.DT)))

	avel := float32(b.AngularVelocity) + b.Torque*gs.DT
	avel -= avel * b.AngularDamping * gs.DT
	b.AngularVelocity = en.AngleD(avel)

	b.Impulse = v.V2(0, 0)
	b.Torque = 0

	n.Position = n.Position.Add(b.Velocity.Scl(gs.DT))
	n.Rotation += en.AngleD(avel * gs.DT)
}
