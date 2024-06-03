package engine

func SliceIndexOf[T comparable](a []T, b T) int {
	for i, v := range a {
		if v == b {
			return i
		}
	}
	return -1
}

func SliceRemove[T comparable](a []T, b T) []T {
	index := SliceIndexOf(a, b)
	if index != -1 {
		return SliceRemoveIndex(a, index)
	}
	return a
}

func SliceIndexWhere[T comparable](a []T, match func(int, T) bool) int {
	for i, v := range a {
		if match(i, v) {
			return i
		}
	}
	return -1
}

func SliceRemoveFirst[T comparable](a []T, match func(int, T) bool) []T {
	index := SliceIndexWhere(a, match)
	if index != -1 {
		return SliceRemoveIndex(a, index)
	}
	return a
}
func SliceRemoveAll[T comparable](a []T, match func(int, T) bool) []T {
	for {
		index := SliceIndexWhere(a, match)
		if index == -1 {
			break
		}
		a = SliceRemoveIndex(a, index)
	}
	return a
}

func SliceRemoveIndex[T any](a []T, index int) []T {
	b := make([]T, 0)
	b = append(b, a[:index]...)
	b = append(b, a[index+1:]...)
	return b
}
