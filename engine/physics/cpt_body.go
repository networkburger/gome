package physics

import "jamesraine/grl/engine"

type PhysicsBodyComponent struct {
	*PhysicsSolver
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

func (s *PhysicsBodyComponent) Event(e engine.NodeEvent, _ *engine.GameState, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.PhysicsSolver.Register(n)
	} else if e == engine.NodeEventUnload {
		s.PhysicsSolver.Unregister(n)
	}
}
