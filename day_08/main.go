package main

import (
	"andurian/adventofcode/2022/util"
)

func main() {
	input := util.ReadSafe("input.txt")
	forest := util.PreprocessTimed(func() Forest { return ForestFromString(input) })

	util.ExecuteTimed(8, 1, func() int { return forest.CountVisibleTrees() })
	util.ExecuteTimed(8, 2, func() int { return forest.BestScenicScore() })
}
