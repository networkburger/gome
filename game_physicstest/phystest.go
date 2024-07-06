package game_physicstest

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/io"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"
)

type physicsTestScene struct {
	circleBallistics physics.BallisticComponent
}

func (s *physicsTestScene) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventDraw:
		render.ClearBackground(18, 65, 68)
		gs.Camera.Position.X = float32(gs.Engine.WindowPixelWidth) / -2
		gs.Camera.Position.Y = float32(gs.Engine.WindowPixelHeight) / -2
	case engine.NodeEventTick:
		io.ProcessInputs(InputOverworld, func(action io.ActionID, power float32) {
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
	solver := physics.NewPhysicsSolver()

	rootNode.AddChild(newLine(e, v.V2(-200, 30), v.V2(200, 30), v.Red))
	rootNode.AddChild(newLine(e, v.V2(-200, -130), v.V2(-200, 30), v.Green))
	rootNode.AddChild(newLine(e, v.V2(200, 30), v.V2(200, -130), v.Blue))
	rootNode.AddChild(newLine(e, v.V2(200, -130), v.V2(-200, -130), v.Brown))

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
	circleVis := game_shared.CircleComponent{
		Radius: 8,
		Color:  v.Green,
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
	Color   v.Color
}

func (p *PhysicsLineSegment) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if event == engine.NodeEventDraw {
		a := gs.Camera.Transform(p.A)
		b := gs.Camera.Transform(p.B)
		render.DrawLine(int32(a.X), int32(a.Y), int32(b.X), int32(b.Y), p.Color)
	}
}
func (p *PhysicsLineSegment) Surfaces(n *engine.Node, pos v.Vec2, radius float32, log physics.CollisionBuffferFunc) {
	didHit, hitAt := physics.CircleSegmentIntersection(radius, pos, p.A, p.B)
	if didHit {
		log(physics.CollisionSurface{
			Normal:       p.N,
			ContactPoint: hitAt,
			SurfaceProperties: physics.SurfaceProperties{
				Friction:    0,
				Restitution: 0.9,
			},
		}, n)
	}
}

func newLine(e *engine.Engine, a, b v.Vec2, col v.Color) *engine.Node {
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
