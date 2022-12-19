package main

import (
	"andurian/adventofcode/2022/util"
	"math"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Point struct {
	X, Y, Z int
}

func Add(a, b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func Subtract(a, b Point) Point {
	return Point{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

type Next func(p Point) Point

func Left(p Point) Point    { return Point{p.X - 1, p.Y + 0, p.Z + 0} }
func Right(p Point) Point   { return Point{p.X + 1, p.Y + 0, p.Z + 0} }
func Down(p Point) Point    { return Point{p.X + 0, p.Y - 1, p.Z + 0} }
func Up(p Point) Point      { return Point{p.X + 0, p.Y + 1, p.Z + 0} }
func Forward(p Point) Point { return Point{p.X + 0, p.Y + 0, p.Z - 1} }
func Back(p Point) Point    { return Point{p.X + 0, p.Y + 0, p.Z + 1} }

func Get6Neighbors(p Point) []Point {
	return []Point{
		Left(p), Right(p), Down(p), Up(p), Forward(p), Back(p),
	}
}

func Less(a, b Point) bool {
	if a.X < b.X {
		return true
	} else if a.X > b.X {
		return false
	} else {
		if a.Y < b.Y {
			return true
		} else if a.Y > b.Y {
			return false
		} else {
			return a.Z < b.Z
		}
	}
}

func Extent(points []Point) (min, max Point) {
	min = Point{math.MaxInt, math.MaxInt, math.MaxInt}
	max = Point{math.MinInt, math.MinInt, math.MinInt}

	for _, p := range points {
		min.X = util.Min(min.X, p.X)
		min.Y = util.Min(min.Y, p.Y)
		min.Z = util.Min(min.Z, p.Z)

		max.X = util.Max(max.X, p.X)
		max.Y = util.Max(max.Y, p.Y)
		max.Z = util.Max(max.Z, p.Z)
	}
	return
}

type Side struct {
	a, b Point
}

func NewSide(a, b Point) Side {
	if Less(a, b) {
		return Side{a, b}
	} else {
		return Side{b, a}
	}
}

func (p Point) Sides() []Side {
	return []Side{
		NewSide(p, Left(p)),
		NewSide(p, Right(p)),
		NewSide(p, Down(p)),
		NewSide(p, Up(p)),
		NewSide(p, Forward(p)),
		NewSide(p, Back(p)),
	}
}

func preprocess(s string) []Point {
	lines := strings.Split(s, "\n")
	ret := make([]Point, len(lines))
	for i, line := range lines {
		tokens := strings.Split(line, ",")
		ret[i] = Point{util.AtoiSafe(tokens[0]), util.AtoiSafe(tokens[1]), util.AtoiSafe(tokens[2])}
	}
	return ret
}

func Task1(points []Point) int {
	sides := make(map[Side]int)
	for _, p := range points {
		for _, s := range p.Sides() {
			if util.MapContainsKey(sides, s) {
				sides[s] += 1
			} else {
				sides[s] = 1
			}
		}
	}

	sum := 0
	for _, v := range sides {
		if v == 1 {
			sum += 1
		}
	}

	return sum
}

func Task2(points []Point) int {
	ones := Point{1, 1, 1}
	min, max := Extent(points)
	min = Subtract(min, ones)
	max = Add(max, ones)

	isInside := func(p Point) bool {
		return p.X >= min.X && p.X <= max.X &&
			p.Y >= min.Y && p.Y <= max.Y &&
			p.Z >= min.Z && p.Z <= max.Z
	}

	pointSet := mapset.NewThreadUnsafeSet(points...)
	outside := mapset.NewThreadUnsafeSet[Point]()

	var fill func(Point)
	fill = func(current Point) {
		if !isInside(current) || pointSet.Contains(current) || outside.Contains(current) {
			return
		}

		outside.Add(current)
		for _, n := range Get6Neighbors(current) {
			fill(n)
		}
	}

	fill(min)

	sum := 0
	for _, p := range points {
		for _, n := range Get6Neighbors(p) {
			if !pointSet.Contains(n) && !outside.Contains(n) {
				sum += 1
			}
		}
	}

	return Task1(points) - sum
}

func main() {
	input := util.ReadSafe("input.txt")
	points := util.PreprocessTimed(func() []Point { return preprocess(input) })
	util.ExecuteTimed(18, 1, func() int { return Task1(points) })
	util.ExecuteTimed(18, 2, func() int { return Task2(points) })
}
