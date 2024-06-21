package game_init

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/ui"
	"jamesraine/grl/game_dig"
	"jamesraine/grl/game_ken"
	"jamesraine/grl/game_physicstest"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type startupScene struct {
	parts.Assets
	*engine.Engine
}

func (s *startupScene) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
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

	font, _ := k.Assets.Font("robotoslab48.json")
	menu := ui.Menu{
		FontRenderer: font,
		Items: []ui.MenuItem{
			{
				MenuLabel: "KEN",
				MenuAction: func() {
					e.PushScene(game_ken.KenScene(e))
				},
			},
			{
				MenuLabel: "DIG",
				MenuAction: func() {
					e.PushScene(game_dig.DigScene(e))
				},
			},
			{
				MenuLabel: "PHYSICS",
				MenuAction: func() {
					e.PushScene(game_physicstest.PhysicsTest(e))
				},
			},
		},
	}

	rootNode := e.NewNode("RootNode - MainMenu")
	rootNode.AddComponent(&k)
	rootNode.AddComponent(&menu)

	return rootNode
}
