package physics

import (
	"jamesraine/grl/engine"
	"log/slog"
)

type PhysicsObstacleComponent struct {
	*PhysicsSolver
	CollisionSurfaceProvider
}

func (s *PhysicsObstacleComponent) Tick(gs *engine.GameState, n *engine.Node) {}
func (s *PhysicsObstacleComponent) Draw(gs *engine.GameState, n *engine.Node) {}
func (s *PhysicsObstacleComponent) Event(e engine.NodeEvent, n *engine.Node) {
	if e == engine.NodeEventLoad {
		if s.CollisionSurfaceProvider == nil {
			slog.Warn("PhysicsObstacleComponent: no ObstacleProvider; ignoring")
		} else {
			s.PhysicsSolver.Register(n)
		}
	} else if e == engine.NodeEventUnload {
		s.PhysicsSolver.Unregister(n)
	}
}
