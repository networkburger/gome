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
	}
}

func StartupScene(e *engine.Engine) *engine.Scene {
	k := startupScene{}
	k.Engine = e
	k.Assets = parts.NewAssets("ass")

	font, _ := k.Assets.Font("robotoslab48.json")
	menu := ui.Menu{
		FontRenderer: font,
		Backout: func() {
			rl.CloseWindow()
		},
		Items: []ui.MenuItem{
			ui.NewMenuItem("KEN", func() {
				e.PushScene(game_ken.KenScene(e))
			}),
			ui.NewMenuItem("DIG", func() {
				e.PushScene(game_dig.DigScene(e))
			}),
			ui.NewMenuItem("Physics Test", func() {
				e.PushScene(game_physicstest.PhysicsTest(e))
			}),
			ui.NewSubMenu("Options", []ui.MenuItem{
				ui.NewMenuItem("OPT1", func() {}),
				ui.NewMenuItem("OPT2", func() {}),
				ui.NewSubMenu("MORE", []ui.MenuItem{
					ui.NewMenuItem("OPT2.1", func() {}),
					ui.NewMenuItem("OPT2.2", func() {}),
					ui.NewMenuItem("OPT2.3", func() {}),
				}),
				ui.NewMenuItem("OPT3", func() {}),
			}),
			ui.NewMenuItem("Quit", func() {
				rl.CloseWindow()
			}),
		},
	}

	rootNode := e.NewNode("RootNode - MainMenu")
	rootNode.AddComponent(&k)
	rootNode.AddComponent(&menu)

	return &engine.Scene{
		Node: rootNode,
	}
}
