package engine

// You gave a float, it's an angle.
// Is it degrees or radians? WHO KNOWS, right?
// Well now, you do. You're welcome.

type AngleD float32
type AngleR float32

func (a AngleD) Rad() AngleR {
	return AngleR(float32(a) * 0.01745)
}

func (a AngleR) Deg() AngleD {
	return AngleD(float32(a) * 57.30659)
}
