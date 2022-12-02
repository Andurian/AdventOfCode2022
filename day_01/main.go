package main

import (
	"andurian/adventofcode/2022/util"

	"fmt"
	"sort"
	"strconv"
	"strings"
)

func countCaloriesPerElf(lines string) []int {
	elves := []int{}
	accumulator := 0
	for _, line := range strings.Split(lines, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			elves = append(elves, accumulator)
			accumulator = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(fmt.Sprintf("Error while converting string %s: %s", line, err))
			}
			accumulator += calories
		}
	}
	elves = append(elves, accumulator)
	return elves
}

func sumTopNCalorieCarryingElves(lines string, n int) int {
	caloriesPerElf := countCaloriesPerElf(lines)
	sort.Sort(sort.Reverse(sort.IntSlice(caloriesPerElf)))

	return util.Accumulate(caloriesPerElf[:n])
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(1, 1, func() int { return sumTopNCalorieCarryingElves(input, 1) })
	util.ExecuteTimed(1, 2, func() int { return sumTopNCalorieCarryingElves(input, 3) })

}
