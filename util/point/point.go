package point

import (
	"andurian/adventofcode/2022/util"
	"fmt"
)

type Point struct {
	Row int
	Col int
}

func (p Point) String() string {
	return fmt.Sprintf("[%d, %d]", p.Row, p.Col)
}

type Next func(p Point) Point

func Left(p Point) Point      { return Point{p.Row + 0, p.Col - 1} }
func Right(p Point) Point     { return Point{p.Row + 0, p.Col + 1} }
func Up(p Point) Point        { return Point{p.Row - 1, p.Col + 0} }
func Down(p Point) Point      { return Point{p.Row + 1, p.Col + 0} }
func DownLeft(p Point) Point  { return Point{p.Row + 1, p.Col - 1} }
func DownRight(p Point) Point { return Point{p.Row + 1, p.Col + 1} }
func UpLeft(p Point) Point    { return Point{p.Row - 1, p.Col - 1} }
func UpRight(p Point) Point   { return Point{p.Row - 1, p.Col + 1} }

func AbsDistances(p1 Point, p2 Point) (dRow int, dCol int) {
	dRow = util.Abs(p2.Row - p1.Row)
	dCol = util.Abs(p2.Col - p1.Col)
	return
}

func Distances(p1 Point, p2 Point) (dRow int, dCol int) {
	dRow = p2.Row - p1.Row
	dCol = p2.Col - p1.Col
	return
}

func ManhattanDistance(p1, p2 Point) int {
	dRow, dCol := AbsDistances(p1, p2)
	return dRow + dCol
}

func Are8Neighbors(p1 Point, p2 Point) bool {
	dRow, dCol := AbsDistances(p1, p2)
	return dRow <= 1 && dCol <= 1
}

func Are4Neighbors(p1 Point, p2 Point) bool {
	dRow, dCol := AbsDistances(p1, p2)

	return dRow <= 1 && dCol == 0 || dRow == 0 && dCol <= 1
}

func Get4Neighbors(p Point) []Point {
	return util.Transform([]func(Point) Point{Up, Right, Down, Left}, func(f func(Point) Point) Point { return f(p) })
}

func Get8Neighbors(p Point) []Point {
	return util.Transform([]func(Point) Point{UpLeft, Up, UpRight, Right, DownRight, Down, DownLeft, Left}, func(f func(Point) Point) Point { return f(p) })
}

func Add(a, b Point) Point {
	return Point{a.Row + b.Row, a.Col + b.Col}
}

func Subtract(a, b Point) Point {
	return Point{a.Row - b.Row, a.Col - b.Col}
}
