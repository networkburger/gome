package physics

import "jamesraine/grl/engine"

type PhysicsBodyComponent struct {
	Radius   float32
	OnGround float64
	SurfaceProperties
}

func (p *PhysicsBodyComponent) IsOnGround(t float64) bool {
	return t-p.OnGround < 0.05
}
func (p *PhysicsBodyComponent) IsOnGroundIsh(t, grace float64) bool {
	return t-p.OnGround < grace
}

func (b *PhysicsBodyComponent) Event(e engine.NodeEvent, s *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.Physics.Register(n)
	} else if e == engine.NodeEventUnload {
		s.Physics.Unregister(n)
	}
}
