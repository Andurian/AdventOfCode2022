package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"container/list"
	"fmt"
	"math"
	"reflect"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	default:
		panic("Invalid Direction")
	}
}

func (d Direction) Rune() rune {
	switch d {
	case North:
		return '^'
	case East:
		return '>'
	case South:
		return 'v'
	case West:
		return '<'
	default:
		panic("Invalid Direction")
	}
}

func TryDirectionFromRune(b rune) (Direction, bool) {
	switch b {
	case '^':
		return North, true
	case '>':
		return East, true
	case 'v':
		return South, true
	case '<':
		return West, true
	default:
		return North, false
	}
}

type Blizzard struct {
	direction Direction
	point.Point
}

func (b Blizzard) NextLocation() point.Point {
	return b.direction.Next()(b.Point)
}

type Valley struct {
	rows, cols int
	start, end point.Point
	grid       map[point.Point][]Blizzard
}

func ValleyFromString(s string) *Valley {
	lines := strings.Split(s, "\n")
	rows := len(lines) - 2
	cols := len(lines[0]) - 2

	startRow := -1
	endRow := len(lines) - 2
	var startCol, endCol int

	for col, c := range lines[0][1 : len(lines[0])-1] {
		if c == '.' {
			startCol = col
			break
		}
	}

	for col, c := range lines[endRow+1][1 : len(lines[endRow+1])-1] {
		if c == '.' {
			endCol = col
			break
		}
	}

	grid := make(map[point.Point][]Blizzard)
	for row, line := range lines[1 : len(lines)-1] {
		for col, c := range line[1 : len(line)-1] {
			direction, ok := TryDirectionFromRune(c)
			if ok {
				p := point.Point{Row: row, Col: col}
				grid[p] = []Blizzard{{direction, p}}
			}
		}
	}

	start := point.Point{Row: startRow, Col: startCol}
	end := point.Point{Row: endRow, Col: endCol}

	return &Valley{rows, cols, start, end, grid}
}

func (v *Valley) ExtentWithWalls() (min, max point.Point) {
	return point.Point{Row: -1, Col: -1}, point.Point{Row: v.rows, Col: v.cols}
}

func (v *Valley) String() string {
	min, max := v.ExtentWithWalls()
	var s strings.Builder
	for col := min.Col; col <= max.Col; col += 1 {
		if v.start.Col == col {
			s.WriteRune('.')
		} else {
			s.WriteRune('#')
		}
	}
	s.WriteRune('\n')

	for row := 0; row < v.rows; row += 1 {
		s.WriteRune('#')
		for col := 0; col < v.cols; col += 1 {
			p := point.Point{Row: row, Col: col}
			blizzards, ok := v.grid[p]
			if !ok {
				s.WriteRune('.')
			} else {
				if len(blizzards) > 1 {
					s.WriteString(fmt.Sprintf("%d", len(blizzards)))
				} else {
					s.WriteRune(blizzards[0].direction.Rune())
				}
			}
		}
		s.WriteRune('#')
		s.WriteRune('\n')
	}

	for col := min.Col; col <= max.Col; col += 1 {
		if v.end.Col == col {
			s.WriteRune('.')
		} else {
			s.WriteRune('#')
		}
	}
	return s.String()
}

func (v *Valley) IsInside(p point.Point) bool {
	return p.Row >= 0 && p.Row < v.rows && p.Col >= 0 && p.Col < v.cols
}

func (v *Valley) IsWalkable(p point.Point) bool {
	if p == v.start || p == v.end {
		return true
	}
	if !v.IsInside(p) {
		return false
	}
	return !util.MapContainsKey(v.grid, p)
}

func (v *Valley) Step() *Valley {
	grid := make(map[point.Point][]Blizzard)
	for _, blizzards := range v.grid {
		for _, blizzard := range blizzards {
			next := blizzard.NextLocation()
			if !v.IsInside(next) {
				switch blizzard.direction {
				case North:
					next.Row = v.rows - 1
				case East:
					next.Col = 0
				case South:
					next.Row = 0
				case West:
					next.Col = v.cols - 1
				default:
				}
			}
			if !util.MapContainsKey(grid, next) {
				grid[next] = []Blizzard{{blizzard.direction, next}}
			} else {
				grid[next] = append(grid[next], Blizzard{blizzard.direction, next})
			}
		}
	}
	return &Valley{v.rows, v.cols, v.start, v.end, grid}
}

type PathSearch struct {
	blizzardStates map[int]*Valley
	loopLength     int
}

func NewSearch(valley *Valley) *PathSearch {
	blizzardStates := make(map[int]*Valley)
	blizzardStates[0] = valley

	loopLength := 0
	next := valley.Step()
	for !reflect.DeepEqual(valley, next) {
		loopLength += 1
		blizzardStates[loopLength] = next
		next = next.Step()
	}

	return &PathSearch{blizzardStates, loopLength + 1}
}

func (s *PathSearch) ValleyStart() point.Point {
	return s.blizzardStates[0].start
}

func (s *PathSearch) ValleyEnd() point.Point {
	return s.blizzardStates[0].end
}

type SearchState struct {
	pos       point.Point
	timestamp int
}

type SearchStateWithHistory struct {
	history []point.Point
	SearchState
}

func NewSearchStateWithHistory(pos point.Point) SearchStateWithHistory {
	return SearchStateWithHistory{[]point.Point{pos}, SearchState{pos, 0}}
}

func (s *PathSearch) normalizedState(state SearchState) SearchState {
	return SearchState{state.pos, state.timestamp % s.loopLength}
}

func (s *PathSearch) candidates(state SearchState) []SearchState {
	ret := []SearchState{}
	nextTimestamp := state.timestamp + 1
	blizzards := s.blizzardStates[nextTimestamp%s.loopLength]
	if blizzards.IsWalkable(state.pos) {
		ret = append(ret, SearchState{state.pos, nextTimestamp})
	}
	for _, p := range point.Get4Neighbors(state.pos) {
		if blizzards.IsWalkable(p) {
			ret = append(ret, SearchState{p, nextTimestamp})
		}
	}
	return ret
}

func (s *PathSearch) Search(start SearchStateWithHistory, target point.Point) SearchStateWithHistory {
	visited := mapset.NewSet(s.normalizedState(start.SearchState))

	queue := list.New()
	queue.PushBack(start)

	best := SearchStateWithHistory{SearchState: SearchState{timestamp: math.MaxInt}}

	makeNextStateWithHistory := func(current SearchStateWithHistory, next SearchState) SearchStateWithHistory {
		history := make([]point.Point, len(current.history)+1)
		copy(history, current.history)
		history[len(history)-1] = next.pos
		return SearchStateWithHistory{history, next}
	}

	for queue.Len() > 0 {
		current := queue.Front()
		queue.Remove(current)
		currentState := current.Value.(SearchStateWithHistory)
		for _, nextState := range s.candidates(currentState.SearchState) {
			if nextState.timestamp > best.timestamp {
				continue
			}
			if nextState.pos == target {
				if nextState.timestamp < best.timestamp {
					best = makeNextStateWithHistory(currentState, nextState)
				}
				continue
			}
			nextStateNormalized := s.normalizedState(nextState)
			if visited.Contains(nextStateNormalized) {
				continue
			}
			visited.Add(nextStateNormalized)
			queue.PushBack(makeNextStateWithHistory(currentState, nextState))
		}
	}
	return best
}

func Task1(search *PathSearch) int {
	start := NewSearchStateWithHistory(search.ValleyStart())
	target := search.ValleyEnd()
	best := search.Search(start, target)
	return best.timestamp
}

func Task2(search *PathSearch) int {
	start := search.ValleyStart()
	target := search.ValleyEnd()

	best := search.Search(NewSearchStateWithHistory(start), target)
	best = search.Search(best, start)
	best = search.Search(best, target)

	return best.timestamp
}

func main() {
	input := util.ReadSafe("input.txt")
	search := util.PreprocessTimed(func() *PathSearch { return NewSearch(ValleyFromString(input)) })
	util.ExecuteTimed(24, 1, func() int { return Task1(search) })
	util.ExecuteTimed(24, 2, func() int { return Task2(search) })
}
