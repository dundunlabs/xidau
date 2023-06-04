package slices

func Map[E1 any, E2 any](s []E1, f func(E1) E2) []E2 {
	result := make([]E2, len(s))
	for _, e := range s {
		result = append(result, f(e))
	}
	return result
}
