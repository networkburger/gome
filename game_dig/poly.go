package game_dig

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LineStripComponent struct {
	Color    rl.Color
	Vertices []rl.Vector2

	_xfv []rl.Vector2
}

func NewLineStripComponent(col rl.Color, verts []rl.Vector2) LineStripComponent {
	return LineStripComponent{
		Color:    col,
		Vertices: verts,
		_xfv:     make([]rl.Vector2, len(verts)),
	}
}

func (s *LineStripComponent) Event(e en.NodeEvent, n *en.Node) {}

func (c *LineStripComponent) Tick(gs *en.GameState, node *en.Node) {
	xf := node.Transform()

	for i := 0; i < len(c._xfv); i++ {
		c._xfv[i] = rl.Vector2Transform(c.Vertices[i], xf)
	}

	rl.DrawLineStrip(c._xfv, c.Color)
}

type CircleComponent struct {
	Color  rl.Color
	Radius float32
}

func (c *CircleComponent) Tick(gs *en.GameState, node *en.Node) {
	pos := node.AbsolutePosition()
	rl.DrawCircle(int32(pos.X), int32(pos.Y), c.Radius, c.Color)
}
