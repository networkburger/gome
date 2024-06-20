package physics

import (
	"jamesraine/grl/engine"
	"log/slog"
)

type PhysicsObstacleComponent struct {
	*PhysicsSolver
	CollisionSurfaceProvider
}

func (s *PhysicsObstacleComponent) Event(e engine.NodeEvent, _ *engine.Scene, n *engine.Node) {
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
