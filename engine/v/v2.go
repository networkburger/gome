package v

import "math"

type Vec2 struct {
	X, Y float32
}

func V2(x, y float32) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{v.X * v2.X, v.Y * v2.Y}
}

func (v Vec2) Div(s float32) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

func (v Vec2) Dot(v2 Vec2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}
func (v Vec2) LenLen() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Dist(v2 Vec2) float32 {
	return v.Sub(v2).Len()
}

func (v Vec2) DistDist(v2 Vec2) float32 {
	return v.Sub(v2).LenLen()
}

func (v Vec2) Scl(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v Vec2) Nrm() Vec2 {
	return v.Div(v.Len())
}

func (v Vec2) Ang() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v Vec2) Rot(a float32) Vec2 {
	cos := float32(math.Cos(float64(a)))
	sin := float32(math.Sin(float64(a)))
	return Vec2{v.X*cos - v.Y*sin, v.X*sin + v.Y*cos}
}

func (v Vec2) Xfm(m Mat) Vec2 {
	return Vec2{
		v.X*m.M0 + v.Y*m.M4 + m.M12,
		v.X*m.M1 + v.Y*m.M5 + m.M13,
	}
}
