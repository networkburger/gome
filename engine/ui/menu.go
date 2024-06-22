package ui

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/io"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/v"
)

type MenuLabel string

type MenuItem struct {
	MenuLabel
	MenuAction engine.DeferredAction
	Subitems   []MenuItem
}

func NewMenuItem(label string, fn engine.DeferredAction) MenuItem {
	return MenuItem{
		MenuLabel:  MenuLabel(label),
		MenuAction: fn,
	}
}

func NewSubMenu(label string, subitems []MenuItem) MenuItem {
	return MenuItem{
		MenuLabel: MenuLabel(label),
		Subitems:  subitems,
	}
}

type Menu struct {
	Items   []MenuItem
	Backout engine.DeferredAction
	parts.FontRenderer
	Cursor []int
}

func (m *Menu) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	switch event {
	case engine.NodeEventLoad:
		m.Cursor = []int{0}
	case engine.NodeEventDraw:
		items := m.Items
		selected := m.Cursor[len(m.Cursor)-1]
		for i := 0; i < len(m.Cursor)-1; i++ {
			items = items[m.Cursor[i]].Subitems
		}
		io.ProcessInputs(menuActions, func(action io.ActionID, power float32) {
			switch action {
			case MenuNext:
				selected++
				if selected >= len(items) {
					selected = 0
				}
			case MenuPrev:
				selected--
				if selected < 0 {
					selected = len(items) - 1
				}

			case MenuBack:
				if len(m.Cursor) > 1 {
					m.Cursor = m.Cursor[:len(m.Cursor)-1]
					items = m.Items
					for i := 0; i < len(m.Cursor)-1; i++ {
						items = items[m.Cursor[i]].Subitems
					}
					selected = m.Cursor[len(m.Cursor)-1]
				} else if m.Backout != nil {
					gs.Engine.Enqueue(m.Backout)
				}

			case MenuSelect:
				if len(items[selected].Subitems) > 0 {
					m.Cursor = append(m.Cursor, 0)
					items = items[selected].Subitems
					selected = 0
				} else if items[selected].MenuAction != nil {
					gs.Engine.Enqueue(items[selected].MenuAction)
				}
			}
		})

		m.Cursor[len(m.Cursor)-1] = selected

		y := int32(float32(gs.Engine.WindowPixelHeight) * 0.1)
		center := gs.Engine.WindowPixelWidth / 2
		for i, item := range items {
			w, h := m.FontRenderer.MeasureText(string(item.MenuLabel))
			x := center - w/2
			col := v.White
			if i == selected {
				col = v.Red
			}
			m.FontRenderer.TextAt(x, y, col, string(item.MenuLabel))
			y += h + 10
		}
	}
}
