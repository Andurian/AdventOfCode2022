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

func Transform[T any, U any](arr []T, f func(T) U) []U {
	transformed := []U{}
	for _, val := range arr {
		transformed = append(transformed, f(val))
	}
	return transformed
}
