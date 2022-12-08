package util

func Accumulate[T Number](values []T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func AccumulateFunc[T any, N Number](arr []T, f func(T) N) N {
	var sum N
	for _, val := range arr {
		sum += f(val)
	}
	return sum
}

func AccumulateMapFunc[K, V comparable, N Number](m map[K]V, f func(V) N) N {
	var sum N
	for _, val := range m {
		sum += f(val)
	}
	return sum
}

func Transform[T any, U any](arr []T, f func(T) U) []U {
	transformed := []U{}
	for _, val := range arr {
		transformed = append(transformed, f(val))
	}
	return transformed
}

func TransformMap[K, V comparable, T any](m map[K]V, f func(V) T) map[K]T {
	transformed := make(map[K]T)
	for k, v := range m {
		transformed[k] = f(v)
	}
	return transformed
}

func Filter[T any](arr []T, f func(T) bool) []T {
	filtered := []T{}
	for _, val := range arr {
		if f(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

func Reverse[T any](s []T) []T {
	a := make([]T, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
