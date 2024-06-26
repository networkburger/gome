package physics

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/v"
)

type SurfaceProperties struct {
	Friction    float32
	Restitution float32
}

type CollisionSurface struct {
	Normal       v.Vec2
	ContactPoint v.Vec2
	Context      interface{}
	SurfaceProperties
}

type CollisionBuffferFunc func(CollisionSurface, *engine.Node)
type CollisionSurfaceProvider interface {
	Surfaces(providerNode *engine.Node, pos v.Vec2, radius float32, enqueue CollisionBuffferFunc)
}

func CircleOverlapsRect(circlePos v.Vec2, radius float32, rectangle v.Rect) bool {
	// Find the closest point to the circle within the rectangle
	closestX := circlePos.X
	closestY := circlePos.Y

	if circlePos.X < rectangle.X {
		closestX = rectangle.X
	} else if circlePos.X > rectangle.X+rectangle.W {
		closestX = rectangle.X + rectangle.W
	}

	if circlePos.Y < rectangle.Y {
		closestY = rectangle.Y
	} else if circlePos.Y > rectangle.Y+rectangle.H {
		closestY = rectangle.Y + rectangle.H
	}

	// Determine collision
	dist := circlePos.DistDist(v.V2(closestX, closestY))
	r2 := radius * radius
	return dist < r2
}

func GenHitsForSquare(pos v.Vec2, radius float32, tileArea v.Rect, surfaceProperties SurfaceProperties, srcNode *engine.Node, log CollisionBuffferFunc, context interface{}) {
	if approxSquareCircleOverlap(pos, radius, tileArea) {
		// TOP EDGE
		collides, collisionPos := CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X, tileArea.Y),
			v.V2(tileArea.X+tileArea.W, tileArea.Y))
		if collides {
			log(CollisionSurface{
				Normal:            v.V2(0, -1),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
				Context:           context,
			}, srcNode)
		}

		// BOTTOM EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X+tileArea.W, tileArea.Y+tileArea.H),
			v.V2(tileArea.X, tileArea.Y+tileArea.H),
		)
		if collides {
			log(CollisionSurface{
				Normal:            v.V2(0, 1),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
				Context:           context,
			}, srcNode)
		}

		// LEFT EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X, tileArea.Y+tileArea.H),
			v.V2(tileArea.X, tileArea.Y),
		)
		if collides {
			log(CollisionSurface{
				Normal:            v.V2(-1, 0),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
				Context:           context,
			}, srcNode)
		}

		// RIGHT EDGE
		collides, collisionPos = CircleSegmentIntersection(radius, pos,
			v.V2(tileArea.X+tileArea.W, tileArea.Y),
			v.V2(tileArea.X+tileArea.W, tileArea.Y+tileArea.H))
		if collides {
			log(CollisionSurface{
				Normal:            v.V2(1, 0),
				ContactPoint:      collisionPos,
				SurfaceProperties: surfaceProperties,
				Context:           context,
			}, srcNode)
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

func approxSquareCircleOverlap(pos v.Vec2, radius float32, square v.Rect) bool {
	sposx := square.X + square.W*0.5
	sposy := square.Y + square.H*0.5
	sradius := 1.5 * square.W * 0.5
	dist := pos.DistDist(v.V2(sposx, sposy))
	return dist < (sradius+radius)*(sradius+radius)
}
