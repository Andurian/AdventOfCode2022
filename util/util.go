package util

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinMax(x, y int) (min, max int) {
	if x < y {
		return x, y
	}
	return y, x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadSafe(filename string) string {
	data, err := os.ReadFile(filename)
	Check(err)
	return strings.ReplaceAll(string(data), "\r\n", "\n")
}

func AtoiSafe(s string) int {
	val, err := strconv.Atoi(s)
	Check(err)
	return val
}

func IsEmptyString(s string) bool {
	return s == ""
}

func IsNotEmptyString(s string) bool {
	return !IsEmptyString(s)
}

func CopyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func MapContainsKey[K, V comparable](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}

func TryGetFromMap[K, V comparable](m map[K]V, k K, defaultValue V) V {
	val, ok := m[k]
	if ok {
		return val
	}
	return defaultValue
}

func IndexOf[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}
