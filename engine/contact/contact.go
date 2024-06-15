package contact

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionSurface struct {
	Normal       rl.Vector2
	ContactPoint rl.Vector2
}

type CollisionSurfaceProvider interface {
	// Not in love with this interface - puts a lot of responsibility on the implementor
	// to know the internal workings of the solver.
	// Better perhaps to provide a function for the implementor to call with each hit rectangle
	Surfaces(n *en.Node, pos rl.Vector2, radius float32, hits []CollisionSurface, nhits *int)
}

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

func GenHitsForSquare(pos rl.Vector2, radius float32, tileArea rl.Rectangle, hits []CollisionSurface, nhits *int) {
	if approxSquareCircleOverlap(pos, radius, tileArea) {
		// TOP EDGE
		collides, collisionPos := circleSegmentIntersection(radius, pos,
			rl.NewVector2(tileArea.X, tileArea.Y),
			rl.NewVector2(tileArea.X+tileArea.Width, tileArea.Y))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:       rl.NewVector2(0, -1),
				ContactPoint: collisionPos,
			}
			*nhits = (*nhits) + 1
		}

		// BOTTOM EDGE
		collides, collisionPos = circleSegmentIntersection(radius, pos,
			rl.NewVector2(tileArea.X, tileArea.Y+tileArea.Height),
			rl.NewVector2(tileArea.X+tileArea.Width, tileArea.Y+tileArea.Height))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:       rl.NewVector2(0, 1),
				ContactPoint: collisionPos,
			}
			*nhits = (*nhits) + 1
		}

		// LEFT EDGE
		collides, collisionPos = circleSegmentIntersection(radius, pos,
			rl.NewVector2(tileArea.X, tileArea.Y),
			rl.NewVector2(tileArea.X, tileArea.Y+tileArea.Height))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:       rl.NewVector2(-1, 0),
				ContactPoint: collisionPos,
			}
			*nhits = (*nhits) + 1
		}

		// LEFT EDGE
		collides, collisionPos = circleSegmentIntersection(radius, pos,
			rl.NewVector2(tileArea.X+tileArea.Width, tileArea.Y),
			rl.NewVector2(tileArea.X+tileArea.Width, tileArea.Y+tileArea.Height))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:       rl.NewVector2(1, 0),
				ContactPoint: collisionPos,
			}
			*nhits = (*nhits) + 1
		}
	}
}

func Maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Absf(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func triangleArea(a, b, c rl.Vector2) float32 {
	return Absf((a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2)
}

func circleSegmentIntersection(radius float32, o, p, q rl.Vector2) (bool, rl.Vector2) {
	oq := rl.Vector2Subtract(o, q)
	op := rl.Vector2Subtract(o, p)
	qp := rl.Vector2Subtract(q, p)
	qo := rl.Vector2Subtract(o, q)
	pq := rl.Vector2Subtract(q, p)
	oplen := rl.Vector2Length(op)
	oqlen := rl.Vector2Length(oq)
	min_dist := float32(99999999)
	max_dist := Maxf(oplen, oqlen)
	var hit rl.Vector2
	opdotqp := rl.Vector2DotProduct(op, qp)
	oqdotpq := rl.Vector2DotProduct(oq, pq)
	if opdotqp > 0 && oqdotpq > 0 {
		min_dist = float32(2) * triangleArea(o, p, q) / rl.Vector2Length(pq)
		qodotqp := rl.Vector2DotProduct(rl.Vector2Normalize(qo), rl.Vector2Normalize(qp))
		qpn := rl.Vector2Normalize(qp)
		hit = rl.Vector2Add(q, rl.Vector2Scale(qpn, qodotqp*oqlen))
	} else if oplen < oqlen {
		hit = p
		min_dist = oplen
	} else {
		hit = q
		min_dist = oqlen
	}
	if min_dist <= radius && max_dist >= radius {
		return true, hit
	} else {
		return false, hit
	}
}

func approxSquareCircleOverlap(pos rl.Vector2, radius float32, square rl.Rectangle) bool {
	sposx := square.X + square.Width*0.5
	sposy := square.Y + square.Height*0.5
	sradius := 1.5 * square.Width * 0.5
	dist := rl.Vector2DistanceSqr(pos, rl.NewVector2(sposx, sposy))
	return dist < (sradius+radius)*(sradius+radius)
}
