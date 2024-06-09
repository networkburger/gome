package util

import rl "github.com/gen2brain/raylib-go/raylib"

func CircleOverlapsRect(circlePos rl.Vector2, radius float32, rectangle rl.Rectangle) bool {
	// Find the closest point to the circle within the rectangle
	closestX := circlePos.X
	closestY := circlePos.Y

	if circlePos.X < rectangle.X {
		closestX = rectangle.X
	} else if circlePos.X > rectangle.X+rectangle.Width {
		closestX = rectangle.X + rectangle.Width
	}

	if circlePos.Y < rectangle.Y {
		closestY = rectangle.Y
	} else if circlePos.Y > rectangle.Y+rectangle.Height {
		closestY = rectangle.Y + rectangle.Height
	}

	// Determine collision
	dist := rl.Vector2DistanceSqr(circlePos, rl.NewVector2(closestX, closestY))
	r2 := radius * radius
	return dist < r2
}
