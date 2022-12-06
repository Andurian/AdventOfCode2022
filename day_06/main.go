package main

import (
	"andurian/adventofcode/2022/util"

	mapset "github.com/deckarep/golang-set/v2"
)

func findMessageStart(input string, windowSize int) int {
	for window, i := input[:windowSize], windowSize; i < len(input); window, i = input[i+1-windowSize:i+1], i+1 {
		set := mapset.NewSet([]byte(window)...)
		if set.Cardinality() == windowSize {
			return i
		}
	}
	return 0
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(6, 1, func() int { return findMessageStart(input, 4) })
	util.ExecuteTimed(6, 2, func() int { return findMessageStart(input, 14) })
}
