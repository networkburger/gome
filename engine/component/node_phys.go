package component

import (
	en "jamesraine/grl/engine"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsManager interface {
	Register(n *en.Node)
	Unregister(n *en.Node)
}

type ObstacleProvider interface {
	// Not in love with this interface - puts a lot of responsibility on the implementor
	// to know the internal workings of the solver.
	// Better perhaps to provide a function for the implementor to call with each hit rectangle
	Obstacles(n *en.Node, pos rl.Vector2, radius float32, hits []rl.Rectangle, nhits *int)
}

type PhysicsObstacleComponent struct {
	PhysicsManager
	ObstacleProvider
}

func (s *PhysicsObstacleComponent) Tick(gs *en.GameState, n *en.Node) {}
func (s *PhysicsObstacleComponent) Event(e en.NodeEvent, n *en.Node) {
	if e == en.NodeEventLoad {
		if s.ObstacleProvider == nil {
			slog.Warn("PhysicsObstacleComponent: no ObstacleProvider; ignoring")
		} else {
			s.PhysicsManager.Register(n)
		}
	} else if e == en.NodeEventUnload {
		s.PhysicsManager.Unregister(n)
	}
}

type PhysicsBodyComponent struct {
	PhysicsManager
	Radius   float32
	OnGround float64
}

func (p *PhysicsBodyComponent) IsOnGround(t float64) bool {
	return t-p.OnGround < 0.05
}
func (p *PhysicsBodyComponent) IsOnGroundIsh(t, grace float64) bool {
	return t-p.OnGround < grace
}

func (s *PhysicsBodyComponent) Tick(gs *en.GameState, n *en.Node) {}
func (s *PhysicsBodyComponent) Event(e en.NodeEvent, n *en.Node) {
	if e == en.NodeEventLoad {
		s.PhysicsManager.Register(n)
	} else if e == en.NodeEventUnload {
		s.PhysicsManager.Unregister(n)
	}
}

type PhysicsSignalComponent struct {
	PhysicsManager
	Radius float32
	Kind   int
}

func (s *PhysicsSignalComponent) Tick(gs *en.GameState, n *en.Node) {}
func (s *PhysicsSignalComponent) Event(e en.NodeEvent, n *en.Node) {
	if e == en.NodeEventLoad {
		s.PhysicsManager.Register(n)
	} else if e == en.NodeEventUnload {
		s.PhysicsManager.Unregister(n)
	}
}
