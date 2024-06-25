package game_shared

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/ui"
	"jamesraine/grl/engine/window"
)

type SceneLauncher func(*engine.Engine) *engine.Scene

var launchKen, launchDig, launchPhys SceneLauncher

func InitSceneLaunchers(ken, dig, phys SceneLauncher) {
	launchKen = ken
	launchDig = dig
	launchPhys = phys
}

type startupScene struct {
	parts.Assets
	*engine.Engine
}

func (s *startupScene) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventDraw:
		render.ClearBackground(18, 65, 68)
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
			window.CloseWindow()
		},
		Items: []ui.MenuItem{
			ui.NewMenuItem("KEN", func() {
				e.SetScene(launchKen(e))
			}),
			ui.NewMenuItem("DIG", func() {
				e.SetScene(launchDig(e))
			}),
			ui.NewMenuItem("Physics Test", func() {
				e.SetScene(launchPhys(e))
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
				window.CloseWindow()
			}),
		},
	}

	rootNode := e.NewNode("RootNode - MainMenu")
	rootNode.AddComponent(&k)
	rootNode.AddComponent(&menu)

	return &engine.Scene{
		Node:            rootNode,
		TargetFramerate: 15,
	}
}
