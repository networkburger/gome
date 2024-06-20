package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type LineStripComponent struct {
	Color    rl.Color
	Vertices []v.Vec2

	_xfv []rl.Vector2
}

func NewLineStripComponent(col rl.Color, verts []v.Vec2) LineStripComponent {
	return LineStripComponent{
		Color:    col,
		Vertices: verts,
		_xfv:     make([]rl.Vector2, len(verts)),
	}
}

func (c *LineStripComponent) Event(event engine.NodeEvent, gs *engine.Scene, node *engine.Node) {
	if event == engine.NodeEventDraw {
		nodeXf := node.Transform()
		xf := v.MatrixMultiply(nodeXf, gs.Camera.Matrix)

		for i := 0; i < len(c._xfv); i++ {
			x := c.Vertices[i].Xfm(xf)
			c._xfv[i] = rl.NewVector2(x.X, x.Y)
		}

		rl.DrawLineStrip(c._xfv, c.Color)
	}
}

type CircleComponent struct {
	Color  rl.Color
	Radius float32
}

func (c *CircleComponent) Event(event engine.NodeEvent, gs *engine.Scene, node *engine.Node) {
	if event == engine.NodeEventDraw {
		pos := gs.Camera.Transform(node.AbsolutePosition())
		rl.DrawCircle(int32(pos.X), int32(pos.Y), c.Radius, c.Color)
	}
}
