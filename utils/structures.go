package utils

import "fmt"

func Map[T, V any](ins []T, f func(T) V) []V {
	result := make([]V, len(ins))
	for i, el := range ins {
		result[i] = f(el)
	}
	return result
}

// Create a map from array with kf providing keys, values are array elements
func ArrayToMap[T, K comparable](ts []T, kf func(T) K) map[K]T {
	result := make(map[K]T)
	for _, t := range ts {
		result[kf(t)] = t
	}
	return result
}

// Create a map from array with kf providing keys, values are pointers to array elements
func ArrayToPtrMap[T, K comparable](ts []T, kf func(T) K) map[K]*T {
	result := make(map[K]*T)
	for _, t := range ts {
		result[kf(t)] = &t
	}
	return result
}

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func CastArray[T, V any](a []V) ([]T, error) {
	result := make([]T, len(a))
	for i, e := range a {
		if te, ok := any(e).(T); ok {
			result[i] = te
		} else {
			return nil, fmt.Errorf("error casting item %d", i)
		}
	}
	return result, nil
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
