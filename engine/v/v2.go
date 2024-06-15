package v

import "math"

type V2 struct {
	X, Y float32
}

func Vec2(x, y float32) V2 {
	return V2{x, y}
}

func (v V2) Add(v2 V2) V2 {
	return V2{v.X + v2.X, v.Y + v2.Y}
}

func (v V2) Sub(v2 V2) V2 {
	return V2{v.X - v2.X, v.Y - v2.Y}
}

func (v V2) Mul(s float32) V2 {
	return V2{v.X * s, v.Y * s}
}

func (v V2) Div(s float32) V2 {
	return V2{v.X / s, v.Y / s}
}

func (v V2) Dot(v2 V2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v V2) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v V2) Normalize() V2 {
	return v.Div(v.Len())
}

func (v V2) Angle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}

func (v V2) Rotate(a float32) V2 {
	cos := float32(math.Cos(float64(a)))
	sin := float32(math.Sin(float64(a)))
	return V2{v.X*cos - v.Y*sin, v.X*sin + v.Y*cos}
}
