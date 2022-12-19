package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"math"

	mapset "github.com/deckarep/golang-set/v2"
)

type LoopFindingChamber struct {
	width        int
	windProvider *WindProvider
	shapes       []PlacedShape
	shapesLoop   []PlacedShape
	occupied     mapset.Set[point.Point]
	topRow       int
}

func NewLoopFindingChamber(width int, windProvider *WindProvider) *LoopFindingChamber {
	return &LoopFindingChamber{
		width, windProvider, []PlacedShape{}, []PlacedShape{}, mapset.NewSet[point.Point](), -1,
	}
}

func (c *LoopFindingChamber) EntryRow() int {
	return c.topRow + 4
}

func Height(shapes []PlacedShape) int {
	min := math.MaxInt
	max := math.MinInt

	for _, s := range shapes {
		for p := range s.Occupied().Iter() {
			min = util.Min(min, p.Row)
			max = util.Max(max, p.Row)
		}
	}
	return max - min + 1
}

type LoopInfo struct {
	baseLength, baseHeight int
	loopLength, loopHeight int
	loop                   []PlacedShape
}

func (c *LoopFindingChamber) DropShape(s Shape) (bool, LoopInfo) {
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

	looping := func(p PlacedShape) (bool, bool) {
		if c.windProvider.currentIndex != 1 {
			return false, false
		}
		if p.shapeType != c.shapesLoop[0].shapeType {
			return true, false
		}
		return true, p.offset.Col == c.shapesLoop[0].offset.Col
	}

	windLooped, shapeLooped := looping(p)

	if shapeLooped {
		totalLength := len(c.shapes)
		loopLength := len(c.shapesLoop)
		baseLength := totalLength - loopLength
		totalHeight := c.topRow
		loopHeight := Height(c.shapesLoop)
		baseHeight := totalHeight - loopHeight
		return true, LoopInfo{baseLength, baseHeight, loopLength, loopHeight, c.shapesLoop}
	}

	if windLooped {
		c.shapesLoop = []PlacedShape{}
	}

	c.shapes = append(c.shapes, p)
	c.shapesLoop = append(c.shapesLoop, p)
	for x := range p.Occupied().Iter() {
		c.topRow = util.Max(c.topRow, x.Row)
		c.occupied.Add(x)
	}

	return false, LoopInfo{}
}

func Task2(input string) int {
	target := 1000000000000
	chamber := NewLoopFindingChamber(7, NewWindProvider(input))
	provider := NewShapeProvider()
	for {
		loopFound, loopInfo := chamber.DropShape(provider.NextShape())
		if loopFound {
			target -= loopInfo.baseLength
			// Why -2 and +1?
			totalHeight := loopInfo.baseHeight + (target/loopInfo.loopLength)*(loopInfo.loopHeight-2) + Height(loopInfo.loop[:target%loopInfo.loopLength]) + 1
			return totalHeight
		}
	}
}
