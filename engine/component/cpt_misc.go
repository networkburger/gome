package component

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
)

type CircleComponent struct {
	Radius float32
	Color  v.Color
}

func (s *CircleComponent) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventDraw {
		pos := gs.Camera.Transform(n.AbsolutePosition())
		sc := n.AbsoluteScale()
		render.DrawCircle(int32(pos.X), int32(pos.Y), s.Radius*sc, s.Color)
	}
}
