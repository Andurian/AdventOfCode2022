package util

import (
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/constraints"
)

func SlicesEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}

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

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadSafe(filename string) string {
	data, err := os.ReadFile(filename)
	Check(err)
	return string(data)
}

func PreprocessTimed[T any](f func() T) T {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Preprocessing done (%s)", elapsed)
	return result
}

func ExecuteTimed(day int, task int, f func() int) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Day %02d-%d: %d (%s)", day, task, result, elapsed)

}

func AtoiSafe(s string) int {
	val, err := strconv.Atoi(s)
	Check(err)
	return val
}
