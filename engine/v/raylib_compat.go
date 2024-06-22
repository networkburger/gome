package v

import rl "github.com/gen2brain/raylib-go/raylib"

func (r Rect) RL() rl.Rectangle {
	return rl.NewRectangle(r.X, r.Y, r.W, r.H)
}

func (v Vec2) RL() rl.Vector2 {
	return rl.NewVector2(v.X, v.Y)
}

func (c Color) RL() rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}
