package util

import (
	"log"
	"time"
)

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
