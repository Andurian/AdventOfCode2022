package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"

	mapset "github.com/deckarep/golang-set/v2"
)

type ShapeType int

const (
	HorizontalLine ShapeType = iota
	Cross
	Corner
	VerticalLine
	Square
)

type Shape struct {
	bricks    mapset.Set[point.Point]
	shapeType ShapeType
}

type ShapeProvider struct {
	horizontalLine, cross, corner, verticalLine, square Shape
	order                                               []*Shape
	currentIndex                                        int
}

func NewShapeProvider() *ShapeProvider {
	p := &ShapeProvider{}
	p.horizontalLine = Shape{
		mapset.NewSet(
			point.Point{Row: 0, Col: 0},
			point.Point{Row: 0, Col: 1},
			point.Point{Row: 0, Col: 2},
			point.Point{Row: 0, Col: 3},
		),
		HorizontalLine,
	}
	p.cross = Shape{
		mapset.NewSet(
			point.Point{Row: 0, Col: 1},
			point.Point{Row: 1, Col: 0},
			point.Point{Row: 1, Col: 1},
			point.Point{Row: 1, Col: 2},
			point.Point{Row: 2, Col: 1},
		),
		Cross,
	}
	p.corner = Shape{
		mapset.NewSet(
			point.Point{Row: 0, Col: 0},
			point.Point{Row: 0, Col: 1},
			point.Point{Row: 0, Col: 2},
			point.Point{Row: 1, Col: 2},
			point.Point{Row: 2, Col: 2},
		),
		Corner,
	}
	p.verticalLine = Shape{
		mapset.NewSet(
			point.Point{Row: 0, Col: 0},
			point.Point{Row: 1, Col: 0},
			point.Point{Row: 2, Col: 0},
			point.Point{Row: 3, Col: 0},
		),
		VerticalLine,
	}
	p.square = Shape{
		mapset.NewSet(
			point.Point{Row: 0, Col: 0},
			point.Point{Row: 0, Col: 1},
			point.Point{Row: 1, Col: 0},
			point.Point{Row: 1, Col: 1},
		),
		Square,
	}
	p.order = []*Shape{&p.horizontalLine, &p.cross, &p.corner, &p.verticalLine, &p.square}
	p.currentIndex = 0
	return p
}

func (p *ShapeProvider) NextShape() Shape {
	ret := *(p.order[p.currentIndex])
	p.currentIndex = (p.currentIndex + 1) % len(p.order)
	return ret
}

type PlacedShape struct {
	offset point.Point
	Shape
}

func (s PlacedShape) Moved(n point.Next) PlacedShape {
	return PlacedShape{n(s.offset), s.Shape}
}

func (s PlacedShape) Occupied() mapset.Set[point.Point] {
	ret := mapset.NewSet[point.Point]()
	for p := range s.bricks.Iter() {
		ret.Add(point.Add(s.offset, p))
	}
	return ret
}

type WindProvider struct {
	order        string
	currentIndex int
}

func NewWindProvider(s string) *WindProvider {
	return &WindProvider{s, 0}
}

func (w *WindProvider) Next() point.Next {
	move := w.order[w.currentIndex]
	//println(string(move))
	var ret point.Next
	if move == '<' {
		ret = point.Left
	} else {
		ret = point.Right
	}
	w.currentIndex = (w.currentIndex + 1) % len(w.order)
	return ret
}

type Chamber struct {
	width        int
	windProvider *WindProvider
	shapes       []PlacedShape
	occupied     mapset.Set[point.Point]
	topRow       int
}

func NewChamber(width int, windProvider *WindProvider) *Chamber {
	return &Chamber{
		width, windProvider, []PlacedShape{}, mapset.NewSet[point.Point](), -1,
	}
}

func (c *Chamber) EntryRow() int {
	return c.topRow + 4
}

func (c *Chamber) DropShape(s Shape) {
	p := PlacedShape{
		point.Point{Row: c.EntryRow(), Col: 2},
		s,
	}

	tryDrop := func(p PlacedShape) (PlacedShape, bool) {
		moved := p.Moved(point.Up)
		occupiedByShape := moved.Occupied()
		if occupiedByShape.Intersect(c.occupied).Cardinality() != 0 {
			return p, false
		}

		for x := range occupiedByShape.Iter() {
			if x.Row < 0 {
				return p, false
			}
		}

		return moved, true
	}

	tryWindMove := func(p PlacedShape) (PlacedShape, bool) {
		moved := p.Moved(c.windProvider.Next())
		occupiedByShape := moved.Occupied()
		if occupiedByShape.Intersect(c.occupied).Cardinality() != 0 {
			return p, false
		}

		for x := range occupiedByShape.Iter() {
			//fmt.Println(x)
			if x.Col < 0 || x.Col >= c.width {
				return p, false
			}
		}

		return moved, true
	}

	for {
		var hasMoved bool
		p, _ = tryWindMove(p)
		p, hasMoved = tryDrop(p)

		if !hasMoved {
			break
		}

	}

	c.shapes = append(c.shapes, p)
	for x := range p.Occupied().Iter() {
		c.topRow = util.Max(c.topRow, x.Row)
		c.occupied.Add(x)
	}
}

func (c *Chamber) String() string {
	s := ""
	for row := c.EntryRow(); row >= 0; row -= 1 {
		s += "|"
		for col := 0; col < c.width; col += 1 {
			p := point.Point{Row: row, Col: col}
			if c.occupied.Contains(p) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "|\n"
	}
	s += "+"
	for col := 0; col < c.width; col += 1 {
		s += "-"
	}
	s += "+"
	return s
}

func Task1(s string) int {
	provider := NewShapeProvider()
	chamber := NewChamber(7, NewWindProvider(s))
	for i := 0; i < 2022; i += 1 {
		chamber.DropShape(provider.NextShape())
	}
	return chamber.topRow + 1
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(17, 1, func() int { return Task1(input) })
	util.ExecuteTimed(17, 1, func() int { return Task2(input) })

}
