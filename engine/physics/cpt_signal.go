package physics

import (
	"jamesraine/grl/engine"
)

type PhysicsSignalComponent struct {
	Radius float32
	Kind   int
}

func (g *PhysicsSignalComponent) Event(e engine.NodeEvent, s *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.Physics.Register(n)
	} else if e == engine.NodeEventUnload {
		s.Physics.Unregister(n)
	}
}
