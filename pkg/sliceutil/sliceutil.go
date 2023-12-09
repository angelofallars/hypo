package sliceutil

// Map transforms a slice's elements based on a function.
func Map[T any, U any](s []T, fn func(T) U) []U {
	n := make([]U, 0, len(s))
	for _, v := range s {
		n = append(n, fn(v))
	}
	return n
}
