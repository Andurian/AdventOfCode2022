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

func Left(p Point) Point {
	return Point{p.Row, p.Col - 1}
}

func Right(p Point) Point {
	return Point{p.Row, p.Col + 1}
}

func Up(p Point) Point {
	return Point{p.Row - 1, p.Col}
}

func Down(p Point) Point {
	return Point{p.Row + 1, p.Col}
}

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

func Are8Neighbors(p1 Point, p2 Point) bool {
	dRow, dCol := AbsDistances(p1, p2)
	return dRow <= 1 && dCol <= 1
}

func Are4Neighbors(p1 Point, p2 Point) bool {
	dRow, dCol := AbsDistances(p1, p2)

	return dRow <= 1 && dCol == 0 || dRow == 0 && dCol <= 1
}
