package render

import (
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawRect(tex Texture2D, srcX, srcY, srcW, srcH, dstX, dstY, dstW, dstH float32, color v.Color) {
	rl.DrawTexturePro(
		rl.Texture2D(tex),
		rl.NewRectangle(srcX, srcY, srcW, srcH),
		rl.NewRectangle(dstX, dstY, dstW, dstH),
		rl.Vector2{},
		0,
		color.RL(),
	)
}

func DrawLine(x1, y1, x2, y2 int32, color v.Color) {
	rl.DrawLine(x1, y1, x2, y2, color.RL())
}

func DrawLineStrip(points []v.Vec2, color v.Color) {
	for i := 1; i < len(points); i++ {
		rl.DrawLine(
			int32(points[i-1].X),
			int32(points[i-1].Y),
			int32(points[i].X),
			int32(points[i].Y),
			color.RL())
	}
}

func ClearBackground(r, b, g uint8) {
	rl.ClearBackground(rl.NewColor(r, b, g, 255))
}

func DrawCircle(x, y int32, radius float32, color v.Color) {
	rl.DrawCircle(x, y, radius, color.RL())
}
