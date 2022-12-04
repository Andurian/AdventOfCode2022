package main

import (
	"strings"

	"andurian/adventofcode/2022/util"
)

func preprocess(input string) []SectionPair {
	transform := func(s string) SectionPair { return SectionPairFromString(strings.TrimSpace(s)) }
	return util.Transform(strings.Split(input, "\n"), transform)
}

func CountFullyContainingSections(pairs []SectionPair) int {
	countIf := func(p SectionPair) int { return util.Btoi(p.FullyContainsEachOther()) }
	return util.AccumulateFunc(pairs, countIf)

}

func CountOverlappingSections(pairs []SectionPair) int {
	countIf := func(p SectionPair) int { return util.Btoi(p.Overlaps()) }
	return util.AccumulateFunc(pairs, countIf)
}

func main() {
	input := util.ReadSafe("input.txt")

	pairs := util.PreprocessTimed(func() []SectionPair { return preprocess(input) })

	util.ExecuteTimed(4, 1, func() int { return CountFullyContainingSections(pairs) })
	util.ExecuteTimed(4, 1, func() int { return CountOverlappingSections(pairs) })
}
