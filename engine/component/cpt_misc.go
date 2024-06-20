package component

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	Radius float32
	Color  rl.Color
}

func (s *CircleComponent) Event(e engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if e == engine.NodeEventDraw {
		pos := gs.Camera.Transform(n.AbsolutePosition())
		sc := n.AbsoluteScale()
		rl.DrawCircle(int32(pos.X), int32(pos.Y), s.Radius*sc, s.Color)
	}
}
