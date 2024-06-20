package game_dig

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(e *engine.Engine, screenWidth, screenHeight int) {
	assets := parts.NewAssets("ass")
	defer assets.Close()

	solver := physics.NewPhysicsSolver(func(b *engine.Node, s *engine.Node) {
		// something hit something
	})

	rootNode := e.NewNode("RootNode")

	mapNode := NewDigMap(e, &solver, &assets)

	player := StandardPlayerNode(e, &solver)
	player.Position = v.V2(985*20, 35*20)
	player.Rotation = 90

	rootNode.AddChild(mapNode)
	rootNode.AddChild(player)
	e.PushScene(rootNode)

	afterRun := func(gs *engine.GameState) {
		gs.Camera.Position.X = player.Position.X - (float32(screenWidth) / 2)
		gs.Camera.Position.Y = player.Position.Y - (float32(screenHeight) / 2)

		pos := player.Position
		ang := player.Rotation
		rl.DrawText(fmt.Sprintf("X: %d, Y: %d, R: %d", int32(pos.X), int32(pos.Y), int32(ang)), int32(screenWidth)-120, int32(screenHeight)-20, 10, rl.Gray)

		solver.Solve(gs)
	}

	rl.SetTargetFPS(30)
	convenience.LegacyLoop(e, screenWidth, screenHeight, nil, afterRun)
}
