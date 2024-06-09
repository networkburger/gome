package game_dig

import (
	"fmt"
	en "jamesraine/grl/engine"
	pt "jamesraine/grl/engine/parts"
	ph "jamesraine/grl/engine/physics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	en.G = en.NewEngine()
	assets := pt.NewAssets("ass")

	solver := ph.NewPhysicsSolver(func(b ph.PhysicsBodyInfo, s ph.PhysicsSignalInfo) {
		// something hit something
	})

	rootNode := en.NewNode("RootNode")

	mapNode := NewDigMap(&solver, &assets)

	player := StandardPlayerNode(&solver)
	player.Position = rl.NewVector2(985*20, 35*20)
	player.Rotation = 90

	en.G.AddChild(rootNode, mapNode)
	en.G.AddChild(rootNode, player)

	rl.SetTargetFPS(30)
	gs := en.GameState{
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
		Camera: &en.Camera{
			Position: rl.NewVector2(0, 0),
			Zoom:     1,
			Rotation: en.AngleD(0),
		},
	}

	en.G.SetScene(rootNode)

	for !rl.WindowShouldClose() {
		gs.DT = rl.GetFrameTime()
		gs.T = rl.GetTime()

		rl.BeginDrawing()

		//rl.ClearBackground(rl.NewColor(18, 65, 68, 255))

		gs.Camera.Position.X = player.Position.X - (float32(screenWidth) / 2)
		gs.Camera.Position.Y = player.Position.Y - (float32(screenHeight) / 2)

		en.G.Run(&gs)

		pos := player.Position
		ang := player.Rotation
		rl.DrawText(fmt.Sprintf("FPS: %d, X: %d, Y: %d, R: %d", rl.GetFPS(), int32(pos.X), int32(pos.Y), int32(ang)), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)

		rl.EndDrawing()
		solver.Solve(&gs)
	}
}
