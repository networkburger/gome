package game_physicstest

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(e *engine.Engine, screenWidth, screenHeight int) {
	rootNode := e.NewNode("RootNode")
	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {})

	rootNode.AddChild(newLine(e, v.V2(-200, 30), v.V2(200, 30), rl.Red, &solver))
	rootNode.AddChild(newLine(e, v.V2(-200, -130), v.V2(-200, 30), rl.Green, &solver))
	rootNode.AddChild(newLine(e, v.V2(200, 30), v.V2(200, -130), rl.Blue, &solver))
	rootNode.AddChild(newLine(e, v.V2(200, -130), v.V2(-200, -130), rl.Brown, &solver))

	circleBody := e.NewNode("Circle")
	circleBallistics := physics.BallisticComponent{
		VelocityDamping: v.V2(0.3, 0.3),
		AngularDamping:  0.8,
		Gravity:         v.V2(0, 30),
	}
	circlePhys := physics.PhysicsBodyComponent{
		PhysicsSolver: &solver,
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
	circleBody.AddComponent(&circleBallistics)
	circleBody.AddComponent(&circlePhys)
	rootNode.AddChild(circleBody)

	///////////////////////////////////////////
	// GET GOING PLZ

	rl.SetTargetFPS(90)

	e.PushScene(rootNode)

	beforeRun := func(gs *engine.GameState) {
		if rl.IsKeyReleased(rl.KeyW) {
			circleBallistics.Impulse = v.V2(0, -7000)
		}
		if rl.IsKeyReleased(rl.KeyS) {
			circleBallistics.Impulse = v.V2(0, 7000)
		}
		if rl.IsKeyReleased(rl.KeyA) {
			circleBallistics.Impulse = v.V2(-7000, 0)
		}
		if rl.IsKeyReleased(rl.KeyD) {
			circleBallistics.Impulse = v.V2(7000, 0)
		}

	}
	afterRun := func(gs *engine.GameState) {
		gs.Camera.Position.X = -float32(screenWidth) / 2
		gs.Camera.Position.Y = -float32(screenHeight) / 2
	}

	convenience.LegacyLoop(e, screenWidth, screenHeight, beforeRun, afterRun)
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
