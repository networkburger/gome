package physics

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/util"
	"jamesraine/grl/engine/v"
	"log/slog"
)

type physicsBodyInfo struct {
	*engine.Node
	*PhysicsBodyComponent
	*BallisticComponent
}
type physicsSignalInfo struct {
	*engine.Node
	*PhysicsSignalComponent
}
type physicsObstacleInfo struct {
	*engine.Node
	*PhysicsObstacleComponent
}

type PhysicsContactNotifier func(*engine.Node, *engine.Node)

type PhysicsSolver struct {
	PhysicsContactNotifier
	obstacles []physicsObstacleInfo
	bodies    []physicsBodyInfo
	signals   []physicsSignalInfo
	hits      []CollisionSurface
}

func NewPhysicsSolver(notifier PhysicsContactNotifier) PhysicsSolver {
	return PhysicsSolver{
		PhysicsContactNotifier: notifier,
		bodies:                 make([]physicsBodyInfo, 0),
		signals:                make([]physicsSignalInfo, 0),
		obstacles:              make([]physicsObstacleInfo, 0),
		hits:                   make([]CollisionSurface, 32),
	}
}

// Register a node with the solver
// Typically a component will call this at NodeEventLoad
func (s *PhysicsSolver) Register(n *engine.Node) {
	obstacle, ok := engine.FindComponent[*PhysicsObstacleComponent](n.Components)
	if ok {
		b := physicsObstacleInfo{
			Node:                     n,
			PhysicsObstacleComponent: obstacle,
		}
		s.obstacles = append(s.obstacles, b)
	}

	body, ok := engine.FindComponent[*PhysicsBodyComponent](n.Components)
	if ok {
		ballistics, ok := engine.FindComponent[*BallisticComponent](n.Components)
		if ok {
			slog.Warn("PhysicsSolver.AddBody: Node has no BallisticComponent; ignoring")
		} else {
			b := physicsBodyInfo{
				Node:                 n,
				BallisticComponent:   ballistics,
				PhysicsBodyComponent: body,
			}
			s.bodies = append(s.bodies, b)
		}
	}

	signal, ok := engine.FindComponent[*PhysicsSignalComponent](n.Components)
	if ok {
		b := physicsSignalInfo{
			Node:                   n,
			PhysicsSignalComponent: signal,
		}
		s.signals = append(s.signals, b)
	}
}

// Unregister a node with the solver
// Typically a component will call this at NodeEventUnload
func (s *PhysicsSolver) Unregister(n *engine.Node) {
	s.bodies = util.SliceRemoveAll(s.bodies, func(_ int, b physicsBodyInfo) bool {
		return b.Node == n
	})
	s.signals = util.SliceRemoveAll(s.signals, func(_ int, b physicsSignalInfo) bool {
		return b.Node == n
	})
	s.obstacles = util.SliceRemoveAll(s.obstacles, func(_ int, b physicsObstacleInfo) bool {
		return b.Node == n
	})
}

func velocityCorrection(inputVelocity, contactNormal v.Vec2, restitution float32) v.Vec2 {
	velm := inputVelocity.Len()
	veln := inputVelocity.Nrm()
	d := contactNormal.Dot(veln)
	scaled := contactNormal.Scl(velm * d * (1 + restitution))
	delta := scaled.Add(inputVelocity)
	correction := inputVelocity.Sub(delta)
	return correction
}

var _up = v.V2(0, -1)

func (s *PhysicsSolver) Solve(gs *engine.GameState) {
	nhits := 0

	for _, b := range s.bodies {
		bpos := b.Node.Position
		radius := b.PhysicsBodyComponent.Radius * b.Node.AbsoluteScale()
		bounces := 0

		// check for colliding signals
		if s.PhysicsContactNotifier != nil {
			for _, sig := range s.signals {
				sigRadius := sig.Node.AbsoluteScale() * sig.PhysicsSignalComponent.Radius
				sigPos := sig.Node.AbsolutePosition()
				if bpos.DistDist(sigPos) < (radius+sigRadius)*(radius+sigRadius) {
					s.PhysicsContactNotifier(b.Node, sig.Node)
				}
			}
		}

		colliding := true
		for colliding {
			for _, o := range s.obstacles {
				o.CollisionSurfaceProvider.Surfaces(o.Node, bpos, radius, s.hits, &nhits)
			}

			closest := -1
			closestDist := float32(999999999)

			for i := 0; i < nhits; i++ {
				hit := s.hits[i]
				dist := bpos.DistDist(hit.ContactPoint)
				if dist < closestDist {
					closest = i
					closestDist = dist
				}
			}
			if closest > -1 {
				hit := s.hits[closest]
				responseNormal := bpos.Sub(hit.ContactPoint).Nrm()

				bpos = hit.ContactPoint.Add(responseNormal.Scl(radius))
				b.Node.Position = bpos
				restitution := hit.Restitution * b.PhysicsBodyComponent.SurfaceProperties.Restitution
				correction := velocityCorrection(b.BallisticComponent.Velocity, responseNormal, restitution)
				b.BallisticComponent.Velocity = b.BallisticComponent.Velocity.Add(correction)
				nhits = 0
				bounces++

				if hit.Normal.Dot(_up) > 0.95 && b.BallisticComponent.Velocity.Y < 0.1 {
					b.PhysicsBodyComponent.OnGround = gs.T
				}

				if bounces > 10 {
					colliding = false
				}
			} else {
				colliding = false
			}
		}
	}
}
