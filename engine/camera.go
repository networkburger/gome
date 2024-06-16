package engine

import (
	"jamesraine/grl/engine/v"
)

type Camera struct {
	Position v.Rect
	Bounds   v.Rect

	// NOTE updated at the start of each frame
	// will NOT immediately reflect changes to position, rotation, or zoom
	// You can always call cache() to force an update
	Matrix v.Mat
}

func (c *Camera) Transform(v v.Vec2) v.Vec2 {
	return v.Xfm(c.Matrix)
}

func (c *Camera) cache() {
	c.Matrix = v.MatrixTranslate(-c.Position.X, -c.Position.Y, 0)
}
