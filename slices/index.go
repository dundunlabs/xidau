package slices

func Index[E comparable](s []E, v E) int {
	for i, e := range s {
		if e == v {
			return i
		}
	}
	return -1
}

func IndexFunc[E any](s []E, f func(E) bool) int {
	for i, e := range s {
		if f(e) {
			return i
		}
	}
	return -1
}
