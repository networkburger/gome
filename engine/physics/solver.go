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

type ContactResponse int32

const (
	ContactResponseNone = ContactResponse(iota)
	ContactResponseBounce
)

type ExtendedContactInfo struct {
	Surface        CollisionSurface
	BodyNode       *engine.Node
	ResponseNormal v.Vec2
}

type ContactObstacleNotifier func(*engine.Node, ExtendedContactInfo) ContactResponse
type ContactSignalNotifier func(*engine.Node, *engine.Node)

type PhysicsSolver struct {
	ContactObstacleNotifier
	ContactSignalNotifier

	obstacles []physicsObstacleInfo
	bodies    []physicsBodyInfo
	signals   []physicsSignalInfo
	hits      []ExtendedContactInfo
	nhits     int
}

func NewPhysicsSolver() PhysicsSolver {
	return PhysicsSolver{
		bodies:    make([]physicsBodyInfo, 0),
		signals:   make([]physicsSignalInfo, 0),
		obstacles: make([]physicsObstacleInfo, 0),
		hits:      make([]ExtendedContactInfo, 32),
	}
}

// Register a node with the solver
// Typically a component will call this at NodeEventLoad
func (s *PhysicsSolver) Register(n *engine.Node) {
	obstacle, psok := engine.FindComponent[*PhysicsObstacleComponent](n.Components)
	if psok {
		b := physicsObstacleInfo{
			Node:                     n,
			PhysicsObstacleComponent: obstacle,
		}
		s.obstacles = append(s.obstacles, b)
	}

	body, bok := engine.FindComponent[*PhysicsBodyComponent](n.Components)
	if bok {
		ballistics, ok := engine.FindComponent[*BallisticComponent](n.Components)
		if !ok {
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

	signal, sok := engine.FindComponent[*PhysicsSignalComponent](n.Components)
	if sok {
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

func (s *PhysicsSolver) _bufferContact(surf CollisionSurface, n *engine.Node) {
	if s.nhits >= len(s.hits) {
		slog.Error("PhysicsSolver: too many hits")
		return
	}
	s.hits[s.nhits].Surface = surf
	s.hits[s.nhits].BodyNode = n
	s.nhits++
}

func (s *PhysicsSolver) Solve(gs *engine.Scene) {
	s.nhits = 0

	for _, b := range s.bodies {
		bpos := b.Node.Position
		radius := b.PhysicsBodyComponent.Radius * b.Node.AbsoluteScale()
		bounces := 0

		// check for colliding signals
		if s.ContactSignalNotifier != nil {
			for _, sig := range s.signals {
				sigRadius := sig.Node.AbsoluteScale() * sig.PhysicsSignalComponent.Radius
				sigPos := sig.Node.AbsolutePosition()
				if bpos.DistDist(sigPos) < (radius+sigRadius)*(radius+sigRadius) {
					s.ContactSignalNotifier(b.Node, sig.Node)
				}
			}
		}

		colliding := true
		for colliding {
			for _, o := range s.obstacles {
				o.CollisionSurfaceProvider.Surfaces(o.Node, bpos, radius, s._bufferContact)
			}

			closest := -1
			closestDist := float32(999999999)

			for i := 0; i < s.nhits; i++ {
				hit := s.hits[i]
				dist := bpos.DistDist(hit.Surface.ContactPoint)
				if dist < closestDist {
					closest = i
					closestDist = dist
				}
			}
			if closest > -1 {
				hit := s.hits[closest]
				hit.ResponseNormal = bpos.Sub(hit.Surface.ContactPoint).Nrm()
				responseKind := ContactResponseBounce
				if s.ContactObstacleNotifier != nil {
					responseKind = s.ContactObstacleNotifier(b.Node, hit)
				}
				if responseKind == ContactResponseBounce {
					bpos = hit.Surface.ContactPoint.Add(hit.ResponseNormal.Scl(radius))
					b.Node.Position = bpos
					restitution := hit.Surface.Restitution * b.PhysicsBodyComponent.SurfaceProperties.Restitution
					correction := velocityCorrection(b.BallisticComponent.Velocity, hit.ResponseNormal, restitution)
					b.BallisticComponent.Velocity = b.BallisticComponent.Velocity.Add(correction)
					s.nhits = 0
					bounces++

					if hit.Surface.Normal.Dot(_up) > 0.95 && b.BallisticComponent.Velocity.Y < 0.1 {
						b.PhysicsBodyComponent.OnGround = gs.T
					}

					if bounces > 10 {
						colliding = false
					}
				} else {
					colliding = false
				}
			} else {
				colliding = false
			}
		}
	}
}
