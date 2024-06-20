package physics

import (
	"jamesraine/grl/engine"
)

type PhysicsSignalComponent struct {
	*PhysicsSolver
	Radius float32
	Kind   int
}

func (s *PhysicsSignalComponent) Event(e engine.NodeEvent, _ *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.PhysicsSolver.Register(n)
	} else if e == engine.NodeEventUnload {
		s.PhysicsSolver.Unregister(n)
	}
}
