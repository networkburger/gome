package game_dig

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	engine.G = engine.NewEngine()
	assets := parts.NewAssets("ass")

	solver := physics.NewPhysicsSolver(func(b physics.PhysicsBodyInfo, s physics.PhysicsSignalInfo) {
		// something hit something
	})

	rootNode := engine.NewNode("RootNode")

	mapNode := NewDigMap(&solver, &assets)

	player := StandardPlayerNode(&solver)
	player.Position = v.V2(985*20, 35*20)
	player.Rotation = 90

	engine.G.AddChild(rootNode, mapNode)
	engine.G.AddChild(rootNode, player)

	rl.SetTargetFPS(30)
	gs := engine.GameState{
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
		Camera: &engine.Camera{
			Position: v.R(0, 0, float32(screenWidth), float32(screenHeight)),
		},
	}

	engine.G.SetScene(rootNode)

	for !rl.WindowShouldClose() {
		gs.DT = rl.GetFrameTime()
		gs.T = rl.GetTime()

		rl.BeginDrawing()

		//rl.ClearBackground(rl.NewColor(18, 65, 68, 255))

		gs.Camera.Position.X = player.Position.X - (float32(screenWidth) / 2)
		gs.Camera.Position.Y = player.Position.Y - (float32(screenHeight) / 2)

		engine.G.Run(&gs)

		pos := player.Position
		ang := player.Rotation
		rl.DrawText(fmt.Sprintf("FPS: %d, X: %d, Y: %d, R: %d", rl.GetFPS(), int32(pos.X), int32(pos.Y), int32(ang)), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)

		rl.EndDrawing()
		solver.Solve(&gs)
	}
}
