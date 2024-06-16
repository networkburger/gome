package game_physicstest

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/contact"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	engine.G = engine.NewEngine()

	rootNode := engine.NewNode("RootNode")

	solver := physics.NewPhysicsSolver(func(b physics.PhysicsBodyInfo, s physics.PhysicsSignalInfo) {

	})

	addLine(v.V2(-200, 30), v.V2(200, 30), rl.Red, &solver, rootNode)
	addLine(v.V2(-200, -130), v.V2(-200, 30), rl.Green, &solver, rootNode)
	addLine(v.V2(200, 30), v.V2(200, -130), rl.Blue, &solver, rootNode)
	addLine(v.V2(200, -130), v.V2(-200, -130), rl.Brown, &solver, rootNode)

	circleBody := engine.NewNode("Circle")
	circleBallistics := component.BallisticComponent{
		VelocityDamping: v.V2(0.3, 0.3),
		AngularDamping:  0.8,
		Gravity:         v.V2(0, 30),
	}
	circlePhys := component.PhysicsBodyComponent{
		PhysicsManager: &solver,
		Radius:         8,
		SurfaceProperties: contact.SurfaceProperties{
			Friction:    0,
			Restitution: 0.33,
		},
	}
	circleVis := component.CircleComponent{
		Radius: 8,
		Color:  rl.Green,
	}
	circleBody.Position = v.V2(0, 0)
	engine.G.AddComponent(circleBody, &circleVis)
	engine.G.AddComponent(circleBody, &circleBallistics)
	engine.G.AddComponent(circleBody, &circlePhys)
	engine.G.AddChild(rootNode, circleBody)

	///////////////////////////////////////////
	// GET GOING PLZ

	rl.SetTargetFPS(90)

	gs := engine.GameState{
		WindowPixelHeight: screenHeight,
		WindowPixelWidth:  screenWidth,
		Camera: &engine.Camera{
			Position: v.R(0, 0, float32(screenWidth), float32(screenHeight)),
		},
	}

	engine.G.SetScene(rootNode)

	for !rl.WindowShouldClose() {
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

		gs.DT = rl.GetFrameTime()
		gs.T = rl.GetTime()
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
		engine.G.Run(&gs)

		gs.Camera.Position.X = -float32(screenWidth) / 2
		gs.Camera.Position.Y = -float32(screenHeight) / 2

		rl.DrawText(fmt.Sprintf("FPS: %d", rl.GetFPS()), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)
		rl.EndDrawing()
		solver.Solve(&gs)

	}
}

type PhysicsLineSegment struct {
	A, B, N v.Vec2
	Color   rl.Color
}

func (s *PhysicsLineSegment) Event(e engine.NodeEvent, n *engine.Node) {}

func (p *PhysicsLineSegment) Tick(gs *engine.GameState, n *engine.Node) {
	a := gs.Camera.Transform(p.A)
	b := gs.Camera.Transform(p.B)
	rl.DrawLine(int32(a.X), int32(a.Y), int32(b.X), int32(b.Y), p.Color)
}
func (p *PhysicsLineSegment) Surfaces(n *engine.Node, pos v.Vec2, radius float32, hits []contact.CollisionSurface, nhits *int) {
	didHit, hitAt := contact.CircleSegmentIntersection(radius, pos, p.A, p.B)
	if didHit {
		hits[*nhits] = contact.CollisionSurface{
			Normal:       p.N,
			ContactPoint: hitAt,
			SurfaceProperties: contact.SurfaceProperties{
				Friction:    0,
				Restitution: 0.9,
			},
		}
		*nhits++
	}
}

func addLine(a, b v.Vec2, col rl.Color, solver *physics.PhysicsSolver, rootNode *engine.Node) {
	groundNode := engine.NewNode("Ground")

	ab := b.Sub(a)
	nm := v.V2(ab.Y, -ab.X).Nrm()
	groundLine := PhysicsLineSegment{
		A:     a,
		B:     b,
		N:     nm,
		Color: col,
	}
	groundSurface := component.PhysicsObstacleComponent{
		PhysicsManager:           solver,
		CollisionSurfaceProvider: &groundLine,
	}
	engine.G.AddComponent(groundNode, &groundLine)
	engine.G.AddComponent(groundNode, &groundSurface)
	engine.G.AddChild(rootNode, groundNode)
}
