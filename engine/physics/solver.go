package physics

import (
	en "jamesraine/grl/engine"
	cm "jamesraine/grl/engine/component"
	"jamesraine/grl/engine/util"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Cache types - store a node and its relevant components
type PhysicsBodyInfo struct {
	*en.Node
	*cm.PhysicsBodyComponent
	*cm.BallisticComponent
}
type PhysicsSignalInfo struct {
	*en.Node
	*cm.PhysicsSignalComponent
}
type PhysicsObstacleInfo struct {
	*en.Node
	*cm.PhysicsObstacleComponent
}

type PhysicsContactNotifier func(PhysicsBodyInfo, PhysicsSignalInfo)

type PhysicsSolver struct {
	PhysicsContactNotifier
	obstacles []PhysicsObstacleInfo
	bodies    []PhysicsBodyInfo
	signals   []PhysicsSignalInfo
	hits      []rl.Rectangle
}

func NewPhysicsSolver(notifier PhysicsContactNotifier) PhysicsSolver {
	return PhysicsSolver{
		PhysicsContactNotifier: notifier,
		bodies:                 make([]PhysicsBodyInfo, 0),
		signals:                make([]PhysicsSignalInfo, 0),
		obstacles:              make([]PhysicsObstacleInfo, 0),
		hits:                   make([]rl.Rectangle, 32),
	}
}

// Register a node with the solver
// Typically a component will call this at NodeEventLoad
func (s *PhysicsSolver) Register(n *en.Node) {
	obstacle := en.FindComponent[*cm.PhysicsObstacleComponent](n.Components)
	if obstacle != nil {
		b := PhysicsObstacleInfo{
			Node:                     n,
			PhysicsObstacleComponent: *obstacle,
		}
		s.obstacles = append(s.obstacles, b)
	}

	body := en.FindComponent[*cm.PhysicsBodyComponent](n.Components)
	if body != nil {
		ballistics := en.FindComponent[*cm.BallisticComponent](n.Components)
		if ballistics == nil {
			slog.Warn("PhysicsSolver.AddBody: Node has no BallisticComponent; ignoring")
		} else {
			b := PhysicsBodyInfo{
				Node:                 n,
				BallisticComponent:   *ballistics,
				PhysicsBodyComponent: *body,
			}
			s.bodies = append(s.bodies, b)
		}
	}

	signal := en.FindComponent[*cm.PhysicsSignalComponent](n.Components)
	if signal != nil {
		b := PhysicsSignalInfo{
			Node:                   n,
			PhysicsSignalComponent: *signal,
		}
		s.signals = append(s.signals, b)
	}
}

// Unregister a node with the solver
// Typically a component will call this at NodeEventUnload
func (s *PhysicsSolver) Unregister(n *en.Node) {
	s.bodies = util.SliceRemoveAll(s.bodies, func(_ int, b PhysicsBodyInfo) bool {
		return b.Node == n
	})
	s.signals = util.SliceRemoveAll(s.signals, func(_ int, b PhysicsSignalInfo) bool {
		return b.Node == n
	})
	s.obstacles = util.SliceRemoveAll(s.obstacles, func(_ int, b PhysicsObstacleInfo) bool {
		return b.Node == n
	})
}

func (s *PhysicsSolver) Solve(gs *en.GameState) {
	nhits := 0

	for _, b := range s.bodies {
		bpos := b.Node.Position
		radius := b.PhysicsBodyComponent.Radius * b.Node.AbsoluteScale()

		for _, o := range s.obstacles {
			o.ObstacleProvider.Obstacles(o.Node, bpos, radius, s.hits, &nhits)
		}

		for i := 0; i < nhits; i++ {
			hit := s.hits[i]
			dt := bpos.Y - hit.Y
			db := hit.Y + hit.Height - bpos.Y
			dl := bpos.X - hit.X
			dr := hit.X + hit.Width - bpos.X
			if dt < db && dt < dl && dt < dr { // TOP
				b.Node.Position = rl.NewVector2(b.Node.Position.X, hit.Y-radius)
				b.BallisticComponent.Velocity = rl.NewVector2(b.BallisticComponent.Velocity.X, 0)
				b.OnGround = gs.T
			} else if db < dt && db < dl && db < dr { // BOTTOM
				b.Node.Position = rl.NewVector2(b.Node.Position.X, hit.Y+hit.Height+radius)
				b.BallisticComponent.Velocity = rl.NewVector2(b.BallisticComponent.Velocity.X, 0)
			} else if dl < dt && dl < db && dl < dr { // LEFT
				b.Node.Position = rl.NewVector2(hit.X-radius, b.Node.Position.Y)
				b.BallisticComponent.Velocity = rl.NewVector2(0, b.BallisticComponent.Velocity.Y)
			} else if dr < dt && dr < db && dr < dl { // RIGHT
				b.Node.Position = rl.NewVector2(hit.X+hit.Width+radius, b.Node.Position.Y)
				b.BallisticComponent.Velocity = rl.NewVector2(0, b.BallisticComponent.Velocity.Y)
			}
		}

		// check for colliding signals
		if s.PhysicsContactNotifier != nil {
			for _, sig := range s.signals {
				sigRadius := sig.Node.AbsoluteScale() * sig.PhysicsSignalComponent.Radius
				sigPos := sig.Node.AbsolutePosition()
				if rl.Vector2DistanceSqr(bpos, sigPos) < (radius+sigRadius)*(radius+sigRadius) {
					s.PhysicsContactNotifier(b, sig)
				}
			}
		}
	}
}
