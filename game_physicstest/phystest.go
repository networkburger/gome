package game_physicstest

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type physicsTestScene struct {
	*engine.Engine
	physics.PhysicsSolver
	circleBallistics physics.BallisticComponent
}

func (s *physicsTestScene) Event(event engine.NodeEvent, gs *engine.GameState, n *engine.Node) {
	switch event {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(90)
	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
		gs.Camera.Position.X = -500
		gs.Camera.Position.Y = -300
	case engine.NodeEventTick:
		if rl.IsKeyReleased(rl.KeyW) {
			s.circleBallistics.Impulse = v.V2(0, -7000)
		}
		if rl.IsKeyReleased(rl.KeyS) {
			s.circleBallistics.Impulse = v.V2(0, 7000)
		}
		if rl.IsKeyReleased(rl.KeyA) {
			s.circleBallistics.Impulse = v.V2(-7000, 0)
		}
		if rl.IsKeyReleased(rl.KeyD) {
			s.circleBallistics.Impulse = v.V2(7000, 0)
		}
	case engine.NodeEventLateTick:
		s.PhysicsSolver.Solve(gs)
	}
}

func PhysicsTest(e *engine.Engine) *engine.Node {
	s := physicsTestScene{}
	rootNode := e.NewNode("RootNode")
	rootNode.AddComponent(&s)
	s.PhysicsSolver = physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {})

	rootNode.AddChild(newLine(e, v.V2(-200, 30), v.V2(200, 30), rl.Red, &s.PhysicsSolver))
	rootNode.AddChild(newLine(e, v.V2(-200, -130), v.V2(-200, 30), rl.Green, &s.PhysicsSolver))
	rootNode.AddChild(newLine(e, v.V2(200, 30), v.V2(200, -130), rl.Blue, &s.PhysicsSolver))
	rootNode.AddChild(newLine(e, v.V2(200, -130), v.V2(-200, -130), rl.Brown, &s.PhysicsSolver))

	circleBody := e.NewNode("Circle")
	s.circleBallistics = physics.BallisticComponent{
		VelocityDamping: v.V2(0.3, 0.3),
		AngularDamping:  0.8,
		Gravity:         v.V2(0, 30),
	}
	circlePhys := physics.PhysicsBodyComponent{
		PhysicsSolver: &s.PhysicsSolver,
		Radius:        8,
		SurfaceProperties: physics.SurfaceProperties{
			Friction:    0,
			Restitution: 0.33,
		},
	}
	circleVis := component.CircleComponent{
		Radius: 8,
		Color:  rl.Green,
	}
	circleBody.Position = v.V2(0, 0)
	circleBody.AddComponent(&circleVis)
	circleBody.AddComponent(&s.circleBallistics)
	circleBody.AddComponent(&circlePhys)
	rootNode.AddChild(circleBody)

	return rootNode
}

type PhysicsLineSegment struct {
	A, B, N v.Vec2
	Color   rl.Color
}

func (p *PhysicsLineSegment) Event(event engine.NodeEvent, gs *engine.GameState, n *engine.Node) {
	if event == engine.NodeEventDraw {
		a := gs.Camera.Transform(p.A)
		b := gs.Camera.Transform(p.B)
		rl.DrawLine(int32(a.X), int32(a.Y), int32(b.X), int32(b.Y), p.Color)
	}
}
func (p *PhysicsLineSegment) Surfaces(n *engine.Node, pos v.Vec2, radius float32, hits []physics.CollisionSurface, nhits *int) {
	didHit, hitAt := physics.CircleSegmentIntersection(radius, pos, p.A, p.B)
	if didHit {
		hits[*nhits] = physics.CollisionSurface{
			Normal:       p.N,
			ContactPoint: hitAt,
			SurfaceProperties: physics.SurfaceProperties{
				Friction:    0,
				Restitution: 0.9,
			},
		}
		*nhits++
	}
}

func newLine(e *engine.Engine, a, b v.Vec2, col rl.Color, solver *physics.PhysicsSolver) *engine.Node {
	groundNode := e.NewNode("Ground")

	ab := b.Sub(a)
	nm := v.V2(ab.Y, -ab.X).Nrm()
	groundLine := PhysicsLineSegment{
		A:     a,
		B:     b,
		N:     nm,
		Color: col,
	}
	groundSurface := physics.PhysicsObstacleComponent{
		PhysicsSolver:            solver,
		CollisionSurfaceProvider: &groundLine,
	}
	groundNode.AddComponent(&groundLine)
	groundNode.AddComponent(&groundSurface)
	return groundNode
}
