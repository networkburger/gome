package game_shared

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/ui"
	"jamesraine/grl/engine/window"
)

func ShowPauseMenu(gs *engine.Scene) {
	gs.Paused = true

	font, _ := gs.Engine.Assets.Font("robotoslab48.json")
	menu := ui.Menu{
		FontRenderer: font,
	}
	resume := func() {
		gs.Engine.RemoveComponentFromNode(gs.Node, &menu)
		gs.Paused = false
	}
	menu.Backout = resume
	menu.Items = []ui.MenuItem{
		{
			MenuLabel:  "RESUME",
			MenuAction: resume,
		},
		{
			MenuLabel: "MAIN MENU",
			MenuAction: func() {
				gs.Paused = false
				gs.Engine.PopScene()
			},
		},
		{
			MenuLabel: "QUIT TO DESKTOP",
			MenuAction: func() {
				window.CloseWindow()
			},
		},
	}
	gs.Engine.Enqueue(func() {
		gs.Node.AddComponent(&menu)
	})
}
