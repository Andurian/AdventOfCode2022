package main

import "andurian/adventofcode/2022/util/point"

type GridConfig interface {
	SideId(point.Point) int
	SideOffset(id int) point.Point
	Side2Global(sideId int, p point.Point) point.Point
	Global2Side(p point.Point) (int, point.Point)
	NextSide(p point.Point, d Direction) (point.Point, Direction)
}

type TestGridConfig struct {
	sideSize int
}

func (g *TestGridConfig) SideId(p point.Point) int {
	sRow := p.Row / g.sideSize
	sCol := p.Col / g.sideSize

	if sRow == 0 && sCol == 2 {
		return 1
	}
	if sRow == 1 && sCol == 0 {
		return 2
	}
	if sRow == 1 && sCol == 1 {
		return 3
	}
	if sRow == 1 && sCol == 2 {
		return 4
	}
	if sRow == 2 && sCol == 2 {
		return 5
	}
	if sRow == 2 && sCol == 3 {
		return 6
	}
	panic("Invalid Side")
}

func (g *TestGridConfig) SideOffset(id int) point.Point {
	s := g.sideSize
	switch id {
	case 1:
		return point.Point{Row: 0 * s, Col: 2 * s}
	case 2:
		return point.Point{Row: 1 * s, Col: 0 * s}
	case 3:
		return point.Point{Row: 1 * s, Col: 1 * s}
	case 4:
		return point.Point{Row: 1 * s, Col: 2 * s}
	case 5:
		return point.Point{Row: 2 * s, Col: 2 * s}
	case 6:
		return point.Point{Row: 2 * s, Col: 3 * s}
	default:
		panic("")
	}
}

func (g *TestGridConfig) Side2Global(sideId int, p point.Point) point.Point {
	return point.Add(p, g.SideOffset(sideId))
}

func (g *TestGridConfig) Global2Side(p point.Point) (int, point.Point) {
	return g.SideId(p), point.Point{Row: p.Row % g.sideSize, Col: p.Col % g.sideSize}
}

func (g *TestGridConfig) NextSide(p point.Point, d Direction) (point.Point, Direction) {
	currentSide, local := g.Global2Side(p)
	max := g.sideSize - 1
	switch currentSide {
	case 1:
		switch d {
		case North:
			return g.Side2Global(2, point.Point{Row: 0, Col: max - local.Col}), South
		case East:
			return g.Side2Global(6, point.Point{Row: max - local.Row, Col: max}), West
		case South:
			panic("")
		case West:
			return g.Side2Global(3, point.Point{Row: 0, Col: local.Row}), South
		default:
			panic("")
		}
	case 2:
		switch d {
		case North:
			return g.Side2Global(1, point.Point{Row: 0, Col: max - local.Col}), South
		case East:
			panic("")
		case South:
			return g.Side2Global(5, point.Point{Row: max, Col: max - local.Col}), North
		case West:
			return g.Side2Global(6, point.Point{Row: max, Col: max - local.Row}), North
		default:
			panic("")
		}
	case 3:
		switch d {
		case North:
			return g.Side2Global(1, point.Point{Row: local.Col, Col: 0}), East
		case East:
			panic("")
		case South:
			return g.Side2Global(5, point.Point{Row: max - local.Col, Col: 0}), East
		case West:
			panic("")
		default:
			panic("")
		}
	case 4:
		switch d {
		case North:
			panic("")
		case East:
			return g.Side2Global(6, point.Point{Row: 0, Col: max - local.Row}), South
		case South:
			panic("")
		case West:
			panic("")
		default:
			panic("")
		}
	case 5:
		switch d {
		case North:
			panic("")
		case East:
			panic("")
		case South:
			return g.Side2Global(2, point.Point{Row: max, Col: max - local.Col}), North
		case West:
			return g.Side2Global(3, point.Point{Row: max, Col: max - local.Row}), North
		default:
			panic("")
		}
	case 6:
		switch d {
		case North:
			return g.Side2Global(4, point.Point{Row: max - local.Col, Col: max}), West
		case East:
			return g.Side2Global(1, point.Point{Row: max - local.Row, Col: max}), West
		case South:
			return g.Side2Global(2, point.Point{Row: max - local.Col, Col: 0}), East
		case West:
			panic("")
		default:
			panic("")
		}
	default:
		panic("")
	}
}

type InputGridConfig struct {
	sideSize int
}

func (g *InputGridConfig) SideId(p point.Point) int {
	sRow := p.Row / g.sideSize
	sCol := p.Col / g.sideSize

	if sRow == 0 && sCol == 1 {
		return 1
	}
	if sRow == 0 && sCol == 2 {
		return 2
	}
	if sRow == 1 && sCol == 1 {
		return 3
	}
	if sRow == 2 && sCol == 0 {
		return 5
	}
	if sRow == 2 && sCol == 1 {
		return 4
	}
	if sRow == 3 && sCol == 0 {
		return 6
	}
	panic("Invalid Side")
}

func (g *InputGridConfig) SideOffset(id int) point.Point {
	s := g.sideSize
	switch id {
	case 1:
		return point.Point{Row: 0 * s, Col: 1 * s}
	case 2:
		return point.Point{Row: 0 * s, Col: 2 * s}
	case 3:
		return point.Point{Row: 1 * s, Col: 1 * s}
	case 4:
		return point.Point{Row: 2 * s, Col: 1 * s}
	case 5:
		return point.Point{Row: 2 * s, Col: 0 * s}
	case 6:
		return point.Point{Row: 3 * s, Col: 0 * s}
	default:
		panic("")
	}
}

func (g *InputGridConfig) Side2Global(sideId int, p point.Point) point.Point {
	return point.Add(p, g.SideOffset(sideId))
}

func (g *InputGridConfig) Global2Side(p point.Point) (int, point.Point) {
	return g.SideId(p), point.Point{Row: p.Row % g.sideSize, Col: p.Col % g.sideSize}
}

func (g *InputGridConfig) NextSide(p point.Point, d Direction) (point.Point, Direction) {
	currentSide, local := g.Global2Side(p)
	max := g.sideSize - 1
	switch currentSide {
	case 1:
		switch d {
		case North:
			return g.Side2Global(6, point.Point{Row: local.Col, Col: 0}), East
		case East:
			panic("")
		case South:
			panic("")
		case West:
			return g.Side2Global(5, point.Point{Row: max - local.Row, Col: 0}), East
		default:
			panic("")
		}
	case 2:
		switch d {
		case North:
			return g.Side2Global(6, point.Point{Row: max, Col: local.Col}), North
		case East:
			return g.Side2Global(4, point.Point{Row: max - local.Row, Col: max}), West
		case South:
			return g.Side2Global(3, point.Point{Row: local.Col, Col: max}), West
		case West:
			panic("")
		default:
			panic("")
		}
	case 3:
		switch d {
		case North:
			panic("")
		case East:
			return g.Side2Global(2, point.Point{Row: max, Col: local.Row}), North
		case South:
			panic("")
		case West:
			return g.Side2Global(5, point.Point{Row: 0, Col: local.Row}), South
		default:
			panic("")
		}
	case 4:
		switch d {
		case North:
			panic("")
		case East:
			return g.Side2Global(2, point.Point{Row: max - local.Row, Col: max}), West
		case South:
			return g.Side2Global(6, point.Point{Row: local.Col, Col: max}), West
		case West:
			panic("")
		default:
			panic("")
		}
	case 5:
		switch d {
		case North:
			return g.Side2Global(3, point.Point{Row: local.Col, Col: 0}), East
		case East:
			panic("")
		case South:
			panic("")
		case West:
			return g.Side2Global(1, point.Point{Row: max - local.Row, Col: 0}), East
		default:
			panic("")
		}
	case 6:
		switch d {
		case North:
			panic("")
		case East:
			return g.Side2Global(4, point.Point{Row: max, Col: local.Row}), North
		case South:
			return g.Side2Global(2, point.Point{Row: 0, Col: local.Col}), South
		case West:
			return g.Side2Global(1, point.Point{Row: 0, Col: local.Row}), South
		default:
			panic("")
		}
	default:
		panic("")
	}
}
