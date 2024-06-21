package game_physicstest

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type physicsTestScene struct {
	circleBallistics physics.BallisticComponent
}

func (s *physicsTestScene) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(90)
	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
		gs.Camera.Position.X = float32(gs.Engine.WindowPixelWidth) / -2
		gs.Camera.Position.Y = float32(gs.Engine.WindowPixelHeight) / -2
	case engine.NodeEventTick:
		engine.ProcessInputs(InputOverworld, func(action engine.ActionID, power float32) {
			switch action {
			case MoveH:
				s.circleBallistics.Impulse.X = power
			case MoveV:
				s.circleBallistics.Impulse.Y = power
			case Pause:
				game_shared.ShowPauseMenu(gs)
			}
		})
	}
}

func PhysicsTest(e *engine.Engine) *engine.Scene {
	s := physicsTestScene{}
	rootNode := e.NewNode("RootNode - PT")
	rootNode.AddComponent(&s)
	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {})

	rootNode.AddChild(newLine(e, v.V2(-200, 30), v.V2(200, 30), rl.Red))
	rootNode.AddChild(newLine(e, v.V2(-200, -130), v.V2(-200, 30), rl.Green))
	rootNode.AddChild(newLine(e, v.V2(200, 30), v.V2(200, -130), rl.Blue))
	rootNode.AddChild(newLine(e, v.V2(200, -130), v.V2(-200, -130), rl.Brown))

	circleBody := e.NewNode("Circle")
	s.circleBallistics = physics.BallisticComponent{
		VelocityDamping: v.V2(0.3, 0.3),
		AngularDamping:  0.8,
		Gravity:         v.V2(0, 30),
	}
	circlePhys := physics.PhysicsBodyComponent{
		Radius: 8,
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

	return &engine.Scene{
		Node:    rootNode,
		Physics: &solver,
	}
}

type PhysicsLineSegment struct {
	A, B, N v.Vec2
	Color   rl.Color
}

func (p *PhysicsLineSegment) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
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

func newLine(e *engine.Engine, a, b v.Vec2, col rl.Color) *engine.Node {
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
		CollisionSurfaceProvider: &groundLine,
	}
	groundNode.AddComponent(&groundLine)
	groundNode.AddComponent(&groundSurface)
	return groundNode
}
