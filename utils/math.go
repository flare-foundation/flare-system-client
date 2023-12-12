package utils

import "golang.org/x/exp/constraints"

// Return max of two values
func Max[T constraints.Ordered](a, b T) T {
	if b > a {
		return b
	}
	return a
}

// Return min of two values
func Min[T constraints.Ordered](a, b T) T {
	if b < a {
		return b
	}
	return a
}

// Return the intersection of two intervals [a1, a2] and [b1, b2]
// If the intervals do not intersect, check if the returned values are in the correct order
func IntervalIntersection[T constraints.Ordered](a1, a2, b1, b2 T) (T, T) {
	return Max(a1, b1), Min(a2, b2)
}
