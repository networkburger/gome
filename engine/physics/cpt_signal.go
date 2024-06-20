package physics

import (
	"jamesraine/grl/engine"
)

type PhysicsSignalComponent struct {
	*PhysicsSolver
	Radius float32
	Kind   int
}

func (s *PhysicsSignalComponent) Tick(gs *engine.GameState, n *engine.Node) {}
func (s *PhysicsSignalComponent) Draw(gs *engine.GameState, n *engine.Node) {}

func (s *PhysicsSignalComponent) Event(e engine.NodeEvent, n *engine.Node) {
	if e == engine.NodeEventLoad {
		s.PhysicsSolver.Register(n)
	} else if e == engine.NodeEventUnload {
		s.PhysicsSolver.Unregister(n)
	}
}
