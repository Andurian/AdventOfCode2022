package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	. "andurian/adventofcode/2022/util/point"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

type Config int

const (
	DebugFull Config = iota
	DebugIntermediate
	TailOnly
)

type DebugObserver struct {
	topLeft     Point
	bottomRight Point

	config Config

	knots         []Point
	start         Point
	visitedByTail mapset.Set[Point]
}

func (v *DebugObserver) tryGetKnotChar(p Point, knots *[]Point) string {
	if knots == nil {
		return ""
	}

	for i, k := range *knots {
		if k == p {
			if i == 0 {
				return "H"
			} else {
				return fmt.Sprint(i)
			}
		}
	}

	return ""
}

func (v *DebugObserver) tryGetStartChar(p Point, start *Point) string {
	if start != nil && p == *start {
		return "s"
	}
	return ""
}

func (v *DebugObserver) tryGetTailChar(p Point, tailHistory *mapset.Set[Point]) string {
	if tailHistory != nil && (*tailHistory).Contains(p) {
		return "#"
	}
	return ""
}

func (v *DebugObserver) printState(knots *[]Point, start *Point, tailHistory *mapset.Set[Point]) {
	for row := v.topLeft.Row; row <= v.bottomRight.Row; row += 1 {
		for col := v.topLeft.Col; col <= v.bottomRight.Col; col += 1 {
			p := Point{Row: row, Col: col}

			charKnot := v.tryGetKnotChar(p, knots)
			charStart := v.tryGetStartChar(p, start)
			charTail := v.tryGetTailChar(p, tailHistory)

			current := charKnot
			if current == "" {
				current = charStart
			}
			if current == "" {
				current = charTail
			}
			if current == "" {
				current = "."
			}

			print(current)
		}
		println()
	}
	println()
}

func (v *DebugObserver) StartMoving(r rope.Rope) {
	v.knots = r.Knots()
	v.visitedByTail.Add(r.TailEnd())

	if v.config != TailOnly {
		println("== Initial State ==")
		println()
		v.printState(&v.knots, &v.start, nil)
	}
}

func (v *DebugObserver) AboutToExecute(i rope.Instruction) {
	if v.config != TailOnly {
		fmt.Println(i)
		println()
	}
}

func (v *DebugObserver) StateChanged(r rope.Rope) {
	v.knots = r.Knots()
	v.visitedByTail.Add(r.TailEnd())

	if v.config == DebugFull {
		v.printState(&v.knots, &v.start, nil)
	}
}

func (v *DebugObserver) FinishedInstruction() {
	if v.config == DebugIntermediate {
		v.printState(&v.knots, &v.start, nil)
	}
}

func (v *DebugObserver) FinishedMoving() {
	if v.config == TailOnly {
		v.printState(nil, nil, &v.visitedByTail)
	}
}

func NewDebugObserver(topLeft, bottomRight Point, config Config) *DebugObserver {
	return &DebugObserver{
		topLeft:       topLeft,
		bottomRight:   bottomRight,
		config:        config,
		visitedByTail: mapset.NewSet[Point](),
	}
}
