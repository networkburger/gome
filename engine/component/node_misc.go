package component

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CircleComponent struct {
	Radius float32
	Color  rl.Color
}

func (s *CircleComponent) Event(e en.NodeEvent, n *en.Node) {}

func (c *CircleComponent) Tick(gs *en.GameState, node *en.Node) {
	pos := node.AbsolutePosition()
	sc := node.AbsoluteScale()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), c.Radius*sc, c.Color)
}
