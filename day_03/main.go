package main

import (
	"andurian/adventofcode/2022/util"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func newSet(content []byte) mapset.Set[byte] {
	ret := mapset.NewSet[byte]()
	for _, val := range content {
		ret.Add(val)
	}
	return ret
}

func toPriority(item byte) int {
	A := byte('A')
	Z := byte('Z')
	a := byte('a')
	z := byte('z')
	if item >= A && item <= Z {
		return int(item-A) + 27
	}
	if item >= a && item <= z {
		return int(item-a) + 1
	}
	panic(fmt.Sprintf("Encountered invalid byte %d", item))
}

func findWrongItem(line []byte) int {
	length := len(line)
	if length%2 != 0 {
		panic(fmt.Sprintf("%q has length %d which is not a multiple of 2", line, length))
	}

	leftCompartment := newSet(line[:length/2])
	rightCompartment := newSet(line[length/2:])

	wrongItems := leftCompartment.Intersect(rightCompartment)

	if wrongItems.Cardinality() != 1 {
		panic("Found more than a single wrong Item")
	}

	return toPriority(wrongItems.ToSlice()[0])
}

func findAllWrongItems(input [][]byte) int {
	total := 0
	for _, line := range input {
		total += findWrongItem(line)
	}
	return total
}

func findBadge(input [][]byte, n int) int {
	set := newSet(input[0])
	for i := 1; i < n; i++ {
		set = set.Intersect(newSet(input[i]))
	}

	if set.Cardinality() != 1 {
		panic("Found more than a single potential item")
	}

	return toPriority(set.ToSlice()[0])
}

func findAllBadges(input [][]byte, n int) int {
	total := 0
	for i := 0; i < len(input)/n; i++ {
		group := input[i*n : (i+1)*n]
		total += findBadge(group, n)
	}
	return total
}

func preprocess(input string) [][]byte {
	var ret [][]byte
	for _, line := range strings.Split(input, "\n") {
		ret = append(ret, []byte(strings.TrimSpace(line)))
	}
	return ret
}

func main() {
	input := preprocess(util.ReadSafe("input.txt"))

	util.ExecuteTimed(3, 1, func() int { return findAllWrongItems(input) })
	util.ExecuteTimed(3, 2, func() int { return findAllBadges(input, 3) })
}
