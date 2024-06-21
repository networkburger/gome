package game_shared

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func ShowPauseMenu(gs *engine.Scene, assets *parts.Assets) {

	gs.Paused = true

	font, _ := assets.Font("robotoslab48.json")
	menu := ui.Menu{
		FontRenderer: font,
		Items: []ui.MenuItem{
			{
				MenuLabel: "RESUME",
			},
			{
				MenuLabel: "MAIN MENU",
				MenuAction: func() {
					gs.Paused = false
					gs.G.PopScene()
				},
			},
			{
				MenuLabel: "QUIT TO DESKTOP",
				MenuAction: func() {
					rl.CloseWindow()
				},
			},
		},
	}
	menu.Items[0].MenuAction = func() {
		gs.G.RemoveComponentFromNode(gs.RootNode, &menu)
		gs.Paused = false
	}
	gs.G.Enqueue(func() {
		gs.RootNode.AddComponent(&menu)
	})
}
