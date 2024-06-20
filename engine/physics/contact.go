package physics

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SurfaceProperties struct {
	Friction    float32
	Restitution float32
}

type CollisionSurface struct {
	Normal       v.Vec2
	ContactPoint v.Vec2
	SurfaceProperties
}

type CollisionSurfaceProvider interface {
	// Not in love with this interface - puts a lot of responsibility on the implementor
	// to know the internal workings of the solver.
	// Better perhaps to provide a function for the implementor to call with each hit rectangle
	Surfaces(n *engine.Node, pos v.Vec2, radius float32, hits []CollisionSurface, nhits *int)
}

func CircleOverlapsRect(circlePos v.Vec2, radius float32, rectangle rl.Rectangle) bool {
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
	dist := circlePos.DistDist(v.V2(closestX, closestY))
	r2 := radius * radius
	return dist < r2
}

func GenHitsForSquare(pos v.Vec2, radius float32, tileArea rl.Rectangle, surfaceProperties SurfaceProperties, hits []CollisionSurface, nhits *int) {
	if approxSquareCircleOverlap(pos, radius, tileArea) {
		// TOP EDGE
		collides, collisionPos := CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X, tileArea.Y),
			v.V2(tileArea.X+tileArea.Width, tileArea.Y))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:            v.V2(0, -1),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
			}
			*nhits = (*nhits) + 1
		}

		// BOTTOM EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X+tileArea.Width, tileArea.Y+tileArea.Height),
			v.V2(tileArea.X, tileArea.Y+tileArea.Height),
		)
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:            v.V2(0, 1),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
			}
			*nhits = (*nhits) + 1
		}

		// LEFT EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X, tileArea.Y+tileArea.Height),
			v.V2(tileArea.X, tileArea.Y),
		)
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:            v.V2(-1, 0),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
			}
			*nhits = (*nhits) + 1
		}

		// RIGHT EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X+tileArea.Width, tileArea.Y),
			v.V2(tileArea.X+tileArea.Width, tileArea.Y+tileArea.Height))
		if collides {
			hits[*nhits] = CollisionSurface{
				Normal:            v.V2(1, 0),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
			}
			*nhits = (*nhits) + 1
		}
	}
}

func triangleArea(a, b, c v.Vec2) float32 {
	ab := b.Sub(a)
	ac := c.Sub(a)
	cp := ab.X*ac.Y - ab.Y*ac.X
	return v.Absf(cp) / 2
}

// reference material:
// https://www.baeldung.com/cs/circle-line-segment-collision-detection

func CircleSegmentIntersection(radius float32, o, p, q v.Vec2) (bool, v.Vec2) {
	oq := o.Sub(q)
	op := o.Sub(p)
	qp := q.Sub(p)
	qo := q.Sub(o)
	pq := p.Sub(q)
	oplen := op.Len()
	oqlen := oq.Len()
	min_dist := float32(99999999)
	max_dist := v.Maxf(oplen, oqlen)
	var hit v.Vec2
	opdotqp := op.Dot(qp)
	oqdotpq := oq.Dot(pq)
	if opdotqp > 0 && oqdotpq > 0 {
		min_dist = float32(2) * triangleArea(o, p, q) / pq.Len()
		qodotqp := qo.Nrm().Dot(qp.Nrm())

		pqn := pq.Nrm()
		hit = q.Add(pqn.Scl(qodotqp * oqlen))
	} else if oplen < oqlen {
		hit = p
		min_dist = oplen
	} else {
		hit = q
		min_dist = oqlen
	}
	if min_dist < radius && max_dist > radius {
		return true, hit
	} else {
		return false, hit
	}
}

func approxSquareCircleOverlap(pos v.Vec2, radius float32, square rl.Rectangle) bool {
	sposx := square.X + square.Width*0.5
	sposy := square.Y + square.Height*0.5
	sradius := 1.5 * square.Width * 0.5
	dist := pos.DistDist(v.V2(sposx, sposy))
	return dist < (sradius+radius)*(sradius+radius)
}
