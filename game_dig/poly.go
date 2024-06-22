package game_dig

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/v"
)

type LineStripComponent struct {
	Color    v.Color
	Vertices []v.Vec2

	_xfv []v.Vec2
}

func NewLineStripComponent(col v.Color, verts []v.Vec2) LineStripComponent {
	return LineStripComponent{
		Color:    col,
		Vertices: verts,
		_xfv:     make([]v.Vec2, len(verts)),
	}
}

func (c *LineStripComponent) Event(event engine.NodeEvent, gs *engine.Scene, node *engine.Node) {
	if event == engine.NodeEventDraw {
		nodeXf := node.Transform()
		xf := v.MatrixMultiply(nodeXf, gs.Camera.Matrix)

		for i := 0; i < len(c._xfv); i++ {
			x := c.Vertices[i].Xfm(xf)
			c._xfv[i] = v.V2(x.X, x.Y)
		}

		render.DrawLineStrip(c._xfv, c.Color)
	}
}

type CircleComponent struct {
	Color  v.Color
	Radius float32
}

func (c *CircleComponent) Event(event engine.NodeEvent, gs *engine.Scene, node *engine.Node) {
	if event == engine.NodeEventDraw {
		pos := gs.Camera.Transform(node.AbsolutePosition())
		render.DrawCircle(int32(pos.X), int32(pos.Y), c.Radius, c.Color)
	}
}
