package main

type Point struct {
	row int
	col int
}

type Next func(p Point) Point

func left(p Point) Point {
	return Point{p.row, p.col - 1}
}

func right(p Point) Point {
	return Point{p.row, p.col + 1}
}

func up(p Point) Point {
	return Point{p.row - 1, p.col}
}

func down(p Point) Point {
	return Point{p.row + 1, p.col}
}
