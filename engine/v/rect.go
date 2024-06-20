package v

type Rect struct {
	X, Y, W, H float32
}

type IntRect struct {
	X, Y, W, H int32
}

func (r IntRect) Contains(x, y int32) bool {
	return x >= r.X && x <= r.X+r.W && y >= r.Y && y <= r.Y+r.H
}

func R(x, y, w, h float32) Rect {
	return Rect{x, y, w, h}
}
func Ri(x, y, w, h int32) IntRect {
	return IntRect{x, y, w, h}
}

func (r Rect) Round() IntRect {
	return IntRect{
		int32(r.X),
		int32(r.Y),
		int32(r.W),
		int32(r.H),
	}
}

func (r Rect) Contains(p Vec2) bool {
	return p.X >= r.X && p.X <= r.X+r.W && p.Y >= r.Y && p.Y <= r.Y+r.H
}

func (r Rect) Overlaps(r2 Rect) bool {
	return r.X < r2.X+r2.W && r.X+r.W > r2.X && r.Y < r2.Y+r2.H && r.Y+r.H > r2.Y
}

// Clean returns a rect with the width and height always positive.
func (r Rect) Clean() Rect {
	return R(
		r.Left(),
		r.Top(),
		r.Right()-r.Left(),
		r.Bottom()-r.Top(),
	)
}

func (r Rect) Left() float32 {
	return Minf(r.X, r.X+r.W)
}

func (r Rect) Right() float32 {
	return Maxf(r.X, r.X+r.W)
}

func (r Rect) Top() float32 {
	return Minf(r.Y, r.Y+r.H)
}

func (r Rect) Bottom() float32 {
	return Maxf(r.Y, r.Y+r.H)
}

func (r Rect) Center() Vec2 {
	return V2((r.Left()+r.Right())*0.5, (r.Top()+r.Bottom())*0.5)
}

func (r Rect) TopLeft() Vec2 {
	return V2(r.Left(), r.Top())
}

func (r Rect) BottomRight() Vec2 {
	return V2(r.Right(), r.Bottom())
}

func (r Rect) TopRight() Vec2 {
	return V2(r.Right(), r.Top())
}

func (r Rect) BottomLeft() Vec2 {
	return V2(r.Left(), r.Bottom())
}
