package v

type Rect struct {
	X, Y, W, H float32
}

func R(x, y, w, h float32) Rect {
	return Rect{x, y, w, h}
}
