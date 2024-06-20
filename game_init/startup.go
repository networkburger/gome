package game_init

import (
	"fmt"
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/ui"
	"jamesraine/grl/game_ken"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type startupScene struct {
	parts.Assets
	*engine.Engine
}

func (s *startupScene) Event(event engine.NodeEvent, gs *engine.GameState, n *engine.Node) {
	switch event {
	case engine.NodeEventSceneActivate:
		rl.SetTargetFPS(15)
	case engine.NodeEventDraw:
		rl.ClearBackground(rl.NewColor(18, 65, 68, 255))
	case engine.NodeEventTick:
		if rl.IsKeyPressed(rl.KeyF1) {
			s.Engine.PushScene(game_ken.KenScene(s.Engine))
		}
	}
}

func StartupScene(e *engine.Engine) *engine.Node {
	k := startupScene{}
	k.Engine = e
	k.Assets = parts.NewAssets("ass")

	fontData, err := k.Assets.FileBytes("robotoslab.json")
	if err != nil {
		panic(err)
	}
	fontSpec, err := parts.FontRead(fontData)
	if err != nil {
		panic(err)
	}

	font := ui.FontRenderer{
		Font:    fontSpec,
		Texture: k.Assets.Texture(fontSpec.ImagePath),
	}

	fmt.Printf("%v", font)

	rootNode := e.NewNode("RootNode")
	rootNode.AddComponent(&k)

	return rootNode
}
