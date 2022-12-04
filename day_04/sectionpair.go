package main

import (
	"regexp"

	"andurian/adventofcode/2022/util"
)

type SectionPair struct {
	first  Section
	second Section
}

func (p SectionPair) Overlaps() bool {
	return p.first.Overlaps(p.second)
}

func (p SectionPair) FullyContainsEachOther() bool {
	return p.first.Contains(p.second) || p.second.Contains(p.first)
}

var sectionRegex = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func SectionPairFromString(s string) SectionPair {
	result := sectionRegex.FindStringSubmatch(s)
	first := Section{util.AtoiSafe(result[1]), util.AtoiSafe(result[2])}
	second := Section{util.AtoiSafe(result[3]), util.AtoiSafe(result[4])}
	return SectionPair{first, second}
}
