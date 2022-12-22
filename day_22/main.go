package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"math"
	"strings"
)

type Tile int

const (
	Empty Tile = iota
	Wall
	NorthTile
	EastTile
	SouthTile
	WestTile
)

func (t Tile) String() string {
	switch t {
	case Empty:
		return "."
	case Wall:
		return "#"
	case NorthTile:
		return "^"
	case EastTile:
		return ">"
	case SouthTile:
		return "v"
	case WestTile:
		return "<"
	default:
		panic("Invalid tile")
	}
}

type Grid struct {
	field                              map[point.Point]Tile
	rowStart, colStart, rowEnd, colEnd map[int]point.Point
	rows, cols                         int
	config                             GridConfig
}

func GridFromString(s string, useRealInput bool) *Grid {
	field := make(map[point.Point]Tile)
	rowStart := make(map[int]point.Point)
	colStart := make(map[int]point.Point)
	rowEnd := make(map[int]point.Point)
	colEnd := make(map[int]point.Point)

	lines := strings.Split(s, "\n")

	cols := math.MinInt
	for row, line := range lines {
		for col, c := range line {
			cols = util.Max(cols, col)
			if c == ' ' {
				continue
			}
			p := point.Point{Row: row, Col: col}
			if !util.MapContainsKey(rowStart, row) {
				rowStart[row] = p
			}
			if !util.MapContainsKey(rowEnd, row) || rowEnd[row].Col < col {
				rowEnd[row] = p
			}
			if !util.MapContainsKey(colStart, col) {
				colStart[col] = p
			}
			if !util.MapContainsKey(colEnd, col) || colEnd[col].Row < row {
				colEnd[col] = p
			}
			if c == '.' {
				field[p] = Empty
			} else if c == '#' {
				field[p] = Wall
			}
		}
	}

	var config GridConfig
	if useRealInput {
		config = &InputGridConfig{50}
	} else {
		config = &TestGridConfig{4}
	}

	return &Grid{field, rowStart, colStart, rowEnd, colEnd, len(lines), cols + 1, config}
}

func (g *Grid) At(p point.Point) (Tile, bool) {
	tile, ok := g.field[p]
	return tile, ok
}

func (g *Grid) Next(p point.Point, d Direction) point.Point {
	nextPos := d.Next()(p)
	if !util.MapContainsKey(g.field, nextPos) {
		switch d {
		case North:
			nextPos = g.colEnd[nextPos.Col]
		case East:
			nextPos = g.rowStart[nextPos.Row]
		case South:
			nextPos = g.colStart[nextPos.Col]
		case West:
			nextPos = g.rowEnd[nextPos.Row]
		default:
			panic("Invalid Dir")
		}

		if !util.MapContainsKey(g.field, nextPos) {
			panic("Something went terribly wrong")
		}
	}
	return nextPos
}

func (g *Grid) NextOnCube(p point.Point, d Direction) (point.Point, Direction) {
	nextPos := d.Next()(p)
	if !util.MapContainsKey(g.field, nextPos) {
		return g.config.NextSide(p, d)
	}
	return nextPos, d
}

func (g *Grid) String() string {
	s := ""
	for row := 0; row < g.rows; row += 1 {
		for col := 0; col < g.cols; col += 1 {
			c, ok := g.field[point.Point{Row: row, Col: col}]
			if !ok {
				s += " "
			} else {
				s += c.String()
			}
		}
		s += "\n"
	}
	return s
}

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

func (d Direction) Value() int {
	switch d {
	case North:
		return 3
	case East:
		return 0
	case South:
		return 1
	case West:
		return 2
	default:
		panic("Invalid Direction")
	}
}

type Rotation int

const (
	Left Rotation = iota
	Right
)

type State struct {
	pos point.Point
	dir Direction
}

func (s State) Password() int {
	return 1000*(s.pos.Row+1) + 4*(s.pos.Col+1) + s.dir.Value()
}

type Instruction interface {
	Move(State, *Grid) State
}

type CubeMoveInstruction struct {
	amount int
}

func NewCubeMoveInstruction(amount int) *CubeMoveInstruction {
	return &CubeMoveInstruction{amount: amount}
}

func (i *CubeMoveInstruction) Move(s State, g *Grid) State {
	pos := s.pos
	dir := s.dir
	for c := 0; c < i.amount; c += 1 {
		nextPos, nextDir := g.NextOnCube(pos, dir)
		tile, ok := g.At(nextPos)
		if !ok {
			panic("Moving went wrong")
		}
		if tile == Wall {
			return State{pos, dir}
		}
		switch dir {
		case North:
			g.field[pos] = NorthTile
		case East:
			g.field[pos] = EastTile
		case South:
			g.field[pos] = SouthTile
		case West:
			g.field[pos] = WestTile
		}
		pos = nextPos
		dir = nextDir
	}
	return State{pos, dir}
}

type MoveInstruction struct {
	amount int
}

func NewMoveInstruction(amount int) *MoveInstruction {
	return &MoveInstruction{amount: amount}
}

func (i *MoveInstruction) Move(s State, g *Grid) State {
	next := s.dir.Next()
	pos := s.pos
	for c := 0; c < i.amount; c += 1 {
		nextPos := next(pos)
		tile, ok := g.At(nextPos)
		if !ok {
			switch s.dir {
			case North:
				nextPos = g.colEnd[nextPos.Col]
			case East:
				nextPos = g.rowStart[nextPos.Row]
			case South:
				nextPos = g.colStart[nextPos.Col]
			case West:
				nextPos = g.rowEnd[nextPos.Row]
			default:
				panic("Invalid Dir")
			}
			tile, ok = g.At(nextPos)
			if !ok {
				panic("Something went terribly wrong")
			}
		}
		if tile == Wall {
			return State{pos, s.dir}
		}
		pos = nextPos
	}
	return State{pos, s.dir}
}

type RotationInstruction struct {
	dir Rotation
}

func NewRotationInstruction(s string) *RotationInstruction {
	var dir Rotation
	if s == "R" {
		dir = Right
	} else {
		dir = Left
	}
	return &RotationInstruction{dir}
}

func (i *RotationInstruction) Move(s State, g *Grid) State {
	nextDir := func(d Direction, r Rotation) Direction {
		switch r {
		case Left:
			switch d {
			case North:
				return West
			case East:
				return North
			case South:
				return East
			case West:
				return South
			default:
				panic("...")
			}
		case Right:
			switch d {
			case North:
				return East
			case East:
				return South
			case South:
				return West
			case West:
				return North
			default:
				panic("...")
			}
		default:
			panic("...")
		}
	}
	return State{s.pos, nextDir(s.dir, i.dir)}
}

func parseInstructions(s string, useCubeMove bool) []Instruction {
	ret := []Instruction{}
	for len(s) > 0 {
		current := string(s[0])
		s = s[1:]
		if current == "R" || current == "L" {
			ret = append(ret, NewRotationInstruction(current))
		} else {
			var next string
			if len(s) > 0 {
				next = string(s[0])
			}
			for next != "R" && next != "L" && len(s) > 0 {
				s = s[1:]
				current += next
				if len(s) > 0 {
					next = string(s[0])
				}
			}
			if useCubeMove {
				ret = append(ret, NewCubeMoveInstruction(util.AtoiSafe(current)))
			} else {
				ret = append(ret, NewMoveInstruction(util.AtoiSafe(current)))
			}
		}
	}
	return ret
}

func parse(input string, useCubeMove, useRealInput bool) (*Grid, []Instruction) {
	parts := strings.Split(input, "\n\n")
	return GridFromString(parts[0], useRealInput), parseInstructions(parts[1], useCubeMove)
}

func CalcPassword(input string, useCubeMove, useRealInput bool) int {
	grid, instructions := util.PreprocessTimedPair(func() (*Grid, []Instruction) { return parse(input, useCubeMove, useRealInput) })
	current := State{grid.rowStart[0], East}
	for _, i := range instructions {
		current = i.Move(current, grid)
	}
	return current.Password()
}

func main() {
	useRealInput := true
	var input string
	if useRealInput {
		input = util.ReadSafe("input.txt")
	} else {
		input = util.ReadSafe("input_test.txt")

	}

	util.ExecuteTimed(22, 1, func() int { return CalcPassword(input, false, useRealInput) })
	util.ExecuteTimed(22, 1, func() int { return CalcPassword(input, true, useRealInput) })

}
