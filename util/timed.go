package util

import (
	"log"
	"time"
)

func PreprocessTimed[T any](f func() *T) *T {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Preprocessing done (%s)", elapsed)
	return result
}

func PreprocessTimedPair[T any, U any](f func() (T, U)) (T, U) {
	start := time.Now()
	resultT, resultU := f()
	elapsed := time.Since(start)
	log.Printf("Preprocessing done (%s)", elapsed)
	return resultT, resultU
}

func ExecuteTimed(day int, task int, f func() int) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Day %02d-%d: %d (%s)", day, task, result, elapsed)
}

func ExecuteTimedString(day int, task int, f func() string) {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	log.Printf("Day %02d-%d: %s (%s)", day, task, result, elapsed)
}
