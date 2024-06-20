package v

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
