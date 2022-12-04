package util

import (
	"os"
	"strconv"

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

func AtoiSafe(s string) int {
	val, err := strconv.Atoi(s)
	Check(err)
	return val
}
