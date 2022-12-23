package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"container/list"
	"math"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) Next() point.Next {
	switch d {
	case North:
		return point.Up
	case East:
		return point.Right
	case South:
		return point.Down
	case West:
		return point.Left
	default:
		panic("Invalid Direction")
	}
}

func DefaultDirections() *list.List {
	ret := list.New()
	ret.PushBack(North)
	ret.PushBack(South)
	ret.PushBack(West)
	ret.PushBack(East)
	return ret
}

type Elf struct {
	directionQueue *list.List
	point.Point
}

func NewElf(p point.Point) *Elf {
	return &Elf{DefaultDirections(), p}
}

func (e *Elf) ConsiderMove(field *Field) *point.Point {
	needToMove := false
	for _, p := range point.Get8Neighbors(e.Point) {
		if field.Occupied(p) {
			needToMove = true
			break
		}
	}
	if !needToMove {
		return nil
	}
	var next point.Point
	for dIter := e.directionQueue.Front(); dIter != nil; dIter = dIter.Next() {
		d := dIter.Value.(Direction)
		switch d {
		case North:
			if !field.Occupied(point.UpLeft(e.Point)) && !field.Occupied(point.Up(e.Point)) && !field.Occupied(point.UpRight(e.Point)) {
				next = point.Up(e.Point)
				return &next
			}
		case East:
			if !field.Occupied(point.UpRight(e.Point)) && !field.Occupied(point.Right(e.Point)) && !field.Occupied(point.DownRight(e.Point)) {
				next = point.Right(e.Point)
				return &next
			}
		case South:
			if !field.Occupied(point.DownRight(e.Point)) && !field.Occupied(point.Down(e.Point)) && !field.Occupied(point.DownLeft(e.Point)) {
				next = point.Down(e.Point)
				return &next
			}
		case West:
			if !field.Occupied(point.DownLeft(e.Point)) && !field.Occupied(point.Left(e.Point)) && !field.Occupied(point.UpLeft(e.Point)) {
				next = point.Left(e.Point)
				return &next
			}
		default:
			panic("")
		}
	}
	return nil
}

func (e *Elf) Move(field *Field, valid map[point.Point]bool) bool {
	p := e.ConsiderMove(field)
	e.directionQueue.MoveAfter(e.directionQueue.Front(), e.directionQueue.Back())
	if p == nil {
		return false
	} else {
		if !util.MapContainsKey(valid, *p) {
			panic("")
		}
		if valid[*p] {
			e.Point = *p
			return true
		} else {
			return false
		}
	}
}

type Field struct {
	elves map[point.Point]*Elf
}

func (f *Field) Extent() (min, max point.Point) {
	min = point.Point{Row: math.MaxInt, Col: math.MaxInt}
	max = point.Point{Row: math.MinInt, Col: math.MinInt}

	for p := range f.elves {
		min.Row = util.Min(min.Row, p.Row)
		min.Col = util.Min(min.Col, p.Col)

		max.Row = util.Max(max.Row, p.Row)
		max.Col = util.Max(max.Col, p.Col)
	}
	return
}

func (f *Field) Occupied(p point.Point) bool {
	return util.MapContainsKey(f.elves, p)
}

func (f *Field) ProposedMoves() map[point.Point]bool {
	ret := make(map[point.Point]bool)
	for _, e := range f.elves {
		p := e.ConsiderMove(f)
		if p == nil {
			continue
		}
		if !util.MapContainsKey(ret, *p) {
			ret[*p] = true
		} else {
			ret[*p] = false
		}
	}
	return ret
}

func (f *Field) Move() bool {
	valid := f.ProposedMoves()
	newField := make(map[point.Point]*Elf)
	someoneMoved := false
	for _, elf := range f.elves {
		someoneMoved = elf.Move(f, valid) || someoneMoved
		newField[elf.Point] = elf
	}
	f.elves = newField
	return someoneMoved
}

func (f *Field) StringWithFixedExtent(min, max point.Point) string {
	s := ""
	for row := min.Row; row <= max.Row; row += 1 {
		for col := min.Col; col <= max.Col; col += 1 {
			p := point.Point{Row: row, Col: col}
			if util.MapContainsKey(f.elves, p) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (f *Field) String() string {
	min, max := f.Extent()
	return f.StringWithFixedExtent(min, max)
}

func (f *Field) EmptyGround() int {
	min, max := f.Extent()
	rows := max.Row - min.Row + 1
	cols := max.Col - min.Col + 1
	return (rows * cols) - len(f.elves)
}

func FieldFromString(s string) *Field {
	ret := make(map[point.Point]*Elf)

	for row, line := range strings.Split(s, "\n") {
		for col, char := range line {
			if char == '#' {
				p := point.Point{Row: row, Col: col}
				ret[p] = NewElf(p)
			}
		}
	}

	return &Field{ret}
}

func Task1(input string) int {
	field := FieldFromString(input)
	for i := 0; i < 10; i += 1 {
		field.Move()
	}
	return field.EmptyGround()
}

func Task2(input string) int {
	field := FieldFromString(input)
	count := 0
	for {
		count += 1
		someoneMoved := field.Move()
		if !someoneMoved {
			return count
		}
	}
}

func main() {
	input := util.ReadSafe("input.txt")
	util.ExecuteTimed(23, 1, func() int { return Task1(input) })
	util.ExecuteTimed(23, 2, func() int { return Task2(input) })
}
