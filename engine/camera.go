package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera struct {
	Position rl.Rectangle
	Bounds   rl.Rectangle

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
}
