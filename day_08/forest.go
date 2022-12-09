package main

import (
	"andurian/adventofcode/2022/util"
	"strings"

	. "andurian/adventofcode/2022/util/point"
)

type Visibility map[Point]bool

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Forest struct {
	rows    int
	cols    int
	heights map[Point]int
}

func (f Forest) ScenicScore(pos Point) int {
	viewDistance := func(p Point, next Next) int {
		distance := 0
		for {
			p = next(p)
			if !f.Contains(p) {
				return distance
			}
			distance += 1
			if f.heights[p] >= f.heights[pos] {
				return distance
			}
		}
	}

	return viewDistance(pos, Up) * viewDistance(pos, Right) * viewDistance(pos, Down) * viewDistance(pos, Left)
}

func (f Forest) BestScenicScore() int {
	bestScore := 0
	for key := range f.heights {
		score := f.ScenicScore(key)
		bestScore = util.Max(bestScore, score)
	}
	return bestScore
}

func (f Forest) VisibilityFrom(pos Direction) Visibility {
	observerTraits := map[Direction]struct {
		start        Point
		advanceOuter Next
		advanceInner Next
	}{
		North: {Point{Row: 0, Col: 0}, Right, Down},
		East:  {Point{Row: 0, Col: f.cols - 1}, Down, Left},
		South: {Point{Row: f.rows - 1, Col: f.cols - 1}, Left, Up},
		West:  {Point{Row: f.rows - 1, Col: 0}, Up, Right},
	}

	adjustVisibility := func(visibility Visibility, start Point, next Next) {
		for current := start; f.Contains(current); current = next(current) {
			if !visibility[current] {
				continue
			}
			for further := next(current); f.Contains(further); further = next(further) {
				if f.heights[further] <= f.heights[current] {
					visibility[further] = false
				}
			}
		}
	}

	visibleTrees := util.TransformMap(f.heights, func(_ int) bool { return true })
	traits := observerTraits[pos]

	for start := traits.start; f.Contains(start); start = traits.advanceOuter(start) {
		adjustVisibility(visibleTrees, start, traits.advanceInner)
	}

	return visibleTrees
}

func (f Forest) Visibility() Visibility {
	vNorth := f.VisibilityFrom(North)
	vEast := f.VisibilityFrom(East)
	vSouth := f.VisibilityFrom(South)
	vWest := f.VisibilityFrom(West)

	visibleTrees := make(Visibility)
	for p := range f.heights {
		visibleTrees[p] = vNorth[p] || vEast[p] || vSouth[p] || vWest[p]
	}
	return visibleTrees
}

func (f Forest) CountVisibleTrees() int {
	return util.AccumulateMapFunc(f.Visibility(), util.Btoi)
}

func (f Forest) Contains(p Point) bool {
	return p.Row >= 0 && p.Row < f.rows &&
		p.Col >= 0 && p.Col < f.cols
}

func ForestFromString(input string) Forest {
	heights := make(map[Point]int)

	lines := strings.Split(input, "\n")

	rows := len(lines)
	cols := len(strings.TrimSpace(lines[0]))

	for row, line := range lines {
		line = strings.TrimSpace(line)
		for col, char := range line {
			heights[Point{Row: row, Col: col}] = util.AtoiSafe(string(char))
		}
	}

	return Forest{rows, cols, heights}
}
