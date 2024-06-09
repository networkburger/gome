package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Position rl.Vector2
	Rotation AngleD
	Zoom     float32

	// NOTE updated at the start of each frame
	// will NOT immediately reflect changes to position, rotation, or zoom
	// You can always call cache() to force an update
	Matrix rl.Matrix
}

func (c *Camera) Transform(v rl.Vector2) rl.Vector2 {
	return rl.Vector2Transform(v, c.Matrix)
}

func (c *Camera) cache() {
	c.Matrix = rl.MatrixTranslate(-c.Position.X, -c.Position.Y, 0)
	c.Matrix = rl.MatrixMultiply(c.Matrix, rl.MatrixRotateZ(float32(c.Rotation.Rad())))
	c.Matrix = rl.MatrixMultiply(c.Matrix, rl.MatrixScale(c.Zoom, c.Zoom, 1))
}
