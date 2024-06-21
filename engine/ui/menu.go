package ui

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/parts"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuLabel string

type MenuItem struct {
	MenuLabel
	MenuAction engine.DeferredAction
}

type Menu struct {
	Items []MenuItem
	parts.FontRenderer
	Selected int
}

func (m *Menu) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventDraw:
		if rl.IsKeyPressed(rl.KeyDown) {
			m.Selected++
			if m.Selected >= len(m.Items) {
				m.Selected = 0
			}
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			m.Selected--
			if m.Selected < 0 {
				m.Selected = len(m.Items) - 1
			}
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			gs.Engine.Enqueue(m.Items[m.Selected].MenuAction)
		}

		y := int(float32(gs.Engine.WindowPixelHeight) * 0.1)
		center := gs.Engine.WindowPixelWidth / 2
		for i, item := range m.Items {
			w, h := m.FontRenderer.MeasureText(string(item.MenuLabel))
			x := center - w/2
			col := White
			if i == m.Selected {
				col = Red
			}
			m.FontRenderer.TextAt(x, y, col.RL(), string(item.MenuLabel))
			y += h + 10
		}
	}
}
