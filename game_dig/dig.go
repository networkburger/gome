package game_dig

import (
	"fmt"
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(screenWidth, screenHeight int) {
	rootNode := en.NewNode("RootNode")

	bgSprite := en.NewBillboard("ass/bg.png")
	mapSprite := en.NewBillboard("ass/map.png")

	baseSize := rl.NewVector2(float32(mapSprite.Texture.Width), float32(mapSprite.Texture.Height))
	worldScale := float32(50)
	worldSize := rl.Vector2Scale(baseSize, worldScale)
	worldRect := rl.NewRectangle(0, 0, worldSize.X, worldSize.Y)

	bgSprite.DstRect = worldRect
	mapSprite.DstRect = worldRect

	mapNode := en.NewNode("Map")
	en.G.AddComponent(&mapNode, &bgSprite)
	en.G.AddComponent(&mapNode, &mapSprite)

	player := StandardPlayerNode()
	player.Position = rl.NewVector2(985*worldScale, 35*worldScale)
	player.Rotation = 90

	en.G.AddChild(&rootNode, &mapNode)
	en.G.AddChild(&rootNode, player)

	rl.SetTargetFPS(30)
	gs := en.GameState{
		WindowPixelHeight: int(screenHeight),
		WindowPixelWidth:  int(screenWidth),
	}

	en.G.SetScene(&rootNode)

	for !rl.WindowShouldClose() {
		gs.DT = rl.GetFrameTime()
		gs.T = rl.GetTime()

		rl.BeginDrawing()

		//rl.ClearBackground(rl.NewColor(18, 65, 68, 255))

		rootNode.Position.X = (float32(screenWidth) / 2) - player.Position.X
		rootNode.Position.Y = (float32(screenHeight) / 2) - player.Position.Y

		en.G.Run(&gs)

		pos := player.Position
		ang := player.Rotation
		rl.DrawText(fmt.Sprintf("FPS: %d, X: %d, Y: %d, R: %d", rl.GetFPS(), int32(pos.X), int32(pos.Y), int32(ang)), int32(screenWidth)-160, int32(screenHeight)-20, 10, rl.Gray)

		rl.EndDrawing()
	}
}
