package util

import (
	"log"
	"os"
	"time"

	"golang.org/x/exp/constraints"
)

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

func ExecuteTimed(day int, task int, f func() int) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Day %02d-%d: %d (%s)", day, task, result, elapsed)

}
