package component

import (
	"jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	Radius float32
	Color  rl.Color
}

func (s *CircleComponent) Event(e engine.NodeEvent, n *engine.Node) {}

func (c *CircleComponent) Tick(gs *engine.GameState, node *engine.Node) {
	pos := gs.Camera.Transform(node.AbsolutePosition())
	sc := node.AbsoluteScale()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), c.Radius*sc, c.Color)
}
