package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"fmt"
	"math"
	"strings"
)

type Result int

const (
	NormalRest Result = iota
	IntoVoid
	Standstill
)

type Grid struct {
	min, max   point.Point
	entryPoint point.Point
	content    map[point.Point]rune
}

func (g *Grid) AddFloor(minCol, maxCol int) {
	row := g.max.Row + 2

	for col := minCol; col <= maxCol; col += 1 {
		g.content[point.Point{Row: row, Col: col}] = '#'
	}

	g.min = point.Point{Row: util.Min(g.min.Row, row), Col: util.Min(g.min.Col, minCol)}
	g.max = point.Point{Row: util.Max(g.max.Row, row), Col: util.Max(g.max.Col, maxCol)}
}

func (g *Grid) IsInside(p point.Point) bool {
	return p.Row >= g.min.Row && p.Row <= g.max.Row && p.Col >= g.min.Col && p.Col <= g.max.Col
}

func (g *Grid) IsOccupied(p point.Point) bool {
	return util.MapContainsKey(g.content, p)
}

func (g *Grid) DropSand() Result {
	getNext := func(p point.Point) point.Point {
		next := point.Down(p)
		if !g.IsOccupied(next) {
			return next
		}
		next = point.DownLeft(p)
		if !g.IsOccupied(next) {
			return next
		}
		next = point.DownRight(p)
		if !g.IsOccupied(next) {
			return next
		}
		return p
	}

	current := g.entryPoint
	for {
		next := getNext(current)
		if next == current {
			g.content[current] = 'o'
			if next == g.entryPoint {
				return Standstill
			}
			return NormalRest
		}

		if !g.IsInside(next) {
			return IntoVoid
		}

		current = next
	}

}

func (g *Grid) Size() (rows, cols int) {
	return g.max.Row - g.min.Row + 1, g.max.Col - g.min.Col + 1
}

func (g *Grid) String() string {
	s := ""
	for row := g.min.Row; row <= g.max.Row; row += 1 {
		for col := g.min.Col; col <= g.max.Col; col += 1 {
			key := point.Point{Row: row, Col: col}
			s += string(util.TryGetFromMap(g.content, key, '.'))
		}
		s += "\n"
	}
	return s
}

func GridFromString(s string, entryPoint point.Point) *Grid {
	grid := make(map[point.Point]rune)
	for _, line := range strings.Split(s, "\n") {
		points := util.Transform(strings.Split(line, " -> "), func(token string) point.Point {
			t := strings.Split(token, ",")
			return point.Point{Row: util.AtoiSafe(t[1]), Col: util.AtoiSafe(t[0])}
		})
		for i, current := range points[:len(points)-1] {
			next := points[i+1]
			//fmt.Printf("Add %v -> %v\n", current, next)
			if next.Row == current.Row {
				colMin, colMax := util.MinMax(current.Col, next.Col)
				for col := colMin; col <= colMax; col += 1 {
					p := point.Point{Row: current.Row, Col: col}
					//fmt.Printf("\t%v", p)
					grid[p] = '#'
				}
			} else if next.Col == current.Col {
				rowMin, rowMax := util.MinMax(current.Row, next.Row)
				for row := rowMin; row <= rowMax; row += 1 {
					p := point.Point{Row: row, Col: current.Col}
					//fmt.Printf("\t%v", p)
					grid[p] = '#'
				}
			} else {
				panic("Cannot handle diagonal wall")
			}
			//println()
		}

	}
	grid[entryPoint] = '+'

	rowMin := math.MaxInt
	rowMax := math.MinInt
	colMin := math.MaxInt
	colMax := math.MinInt

	for p := range grid {
		rowMin = util.Min(rowMin, p.Row)
		rowMax = util.Max(rowMax, p.Row)
		colMin = util.Min(colMin, p.Col)
		colMax = util.Max(colMax, p.Col)
	}

	return &Grid{
		min:        point.Point{Row: rowMin, Col: colMin},
		max:        point.Point{Row: rowMax, Col: colMax},
		entryPoint: entryPoint,
		content:    grid,
	}
}

func Task1(input string) int {
	entryPoint := point.Point{Row: 0, Col: 500}
	grid := GridFromString(input, entryPoint)

	sandCount := 0
	for {
		res := grid.DropSand()
		switch res {
		case NormalRest:
			{
				sandCount += 1
			}
		case IntoVoid:
			{
				fmt.Println(grid)
				return sandCount
			}
		case Standstill:
			{
				panic("Sandstill")
			}
		default:
			panic("...")
		}
	}
}

func Task2(input string) int {
	entryPoint := point.Point{Row: 0, Col: 500}
	grid := GridFromString(input, entryPoint)
	grid.AddFloor(0, 1000)

	sandCount := 0
	for {
		res := grid.DropSand()
		//fmt.Println(grid)
		switch res {
		case NormalRest:
			{
				sandCount += 1
			}
		case IntoVoid:
			{
				//fmt.Println(grid)
				panic("void")
			}
		case Standstill:
			{
				//fmt.Println(grid)
				return sandCount + 1
			}
		default:
			panic("...")
		}
	}
}

func main() {
	input := util.ReadSafe("input.txt")
	util.ExecuteTimed(14, 1, func() int { return Task1(input) })
	util.ExecuteTimed(14, 2, func() int { return Task2(input) })
}
