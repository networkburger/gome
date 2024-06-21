package physics

import (
	"jamesraine/grl/engine"
	"log/slog"
)

type PhysicsObstacleComponent struct {
	CollisionSurfaceProvider
}

func (o *PhysicsObstacleComponent) Event(e engine.NodeEvent, s *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventLoad {
		if o.CollisionSurfaceProvider == nil {
			slog.Warn("PhysicsObstacleComponent: no ObstacleProvider; ignoring")
		} else {
			s.Physics.Register(n)
		}
	} else if e == engine.NodeEventUnload {
		s.Physics.Unregister(n)
	}
}
