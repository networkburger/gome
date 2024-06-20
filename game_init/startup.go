package game_init

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/convenience"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GameLoop(e *engine.Engine, screenWidth, screenHeight int) {
	assets := parts.NewAssets("ass")

	fontData, err := assets.FileBytes("robotoslab.json")
	if err != nil {
		panic(err)
	}
	fontSpec, err := parts.FontRead(fontData)
	if err != nil {
		panic(err)
	}

	font := ui.FontRenderer{
		Font:    fontSpec,
		Texture: assets.Texture(fontSpec.ImagePath),
	}

	fmt.Printf("%v", font)

	rootNode := e.NewNode("RootNode")

	rl.SetTargetFPS(15)

	e.SetScene(rootNode)

	beforeRun := func(gs *engine.GameState) {
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	}

	convenience.StandardLoop(e, screenWidth, screenHeight, beforeRun, nil)
}
