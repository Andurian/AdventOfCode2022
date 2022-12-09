package rope

import (
	"andurian/adventofcode/2022/util"
	. "andurian/adventofcode/2022/util/point"
	"fmt"
)

type Rope struct {
	knots []Point
}

func (r *Rope) Reset(start Point) {
	for i := range r.knots {
		r.knots[i] = start
	}
}

func (r *Rope) Move(next Next) {
	r.knots[0] = next(r.knots[0])
	r.sanitizeTail(next)
}

func (r *Rope) Knots() []Point {
	return r.knots
}

func (r *Rope) TailEnd() Point {
	return r.knots[len(r.knots)-1]
}

func (r *Rope) sanitizeTail(lastHeadMove Next) {
	for i := 0; i < len(r.knots)-1; i += 1 {
		moveKnots(&r.knots[i], &r.knots[i+1])
	}
}

func moveKnots(head *Point, tail *Point) {
	if Are8Neighbors(*head, *tail) {
		return
	}

	dRow, dCol := Distances(*tail, *head)
	if dRow != 0 && dCol != 0 {
		if util.Abs(dRow) == 1 {
			tail.Row = head.Row
			dRow = 0
		} else if util.Abs(dCol) == 1 {
			tail.Col = head.Col
			dCol = 0
		}
	}

	if dRow < 0 {
		tail.Row += dRow + 1
	} else if dRow > 0 {
		tail.Row += dRow - 1
	}

	if dCol < 0 {
		tail.Col += dCol + 1
	} else if dCol > 0 {
		tail.Col += dCol - 1
	}
}

func (r *Rope) String() string {
	ret := fmt.Sprintf("%v", r.knots[0])
	for _, k := range r.knots[1:] {
		ret += fmt.Sprintf(" - %v", k)
	}
	return ret
}

func NewRope(tailLength int) *Rope {
	start := Point{Row: 0, Col: 0}
	r := &Rope{make([]Point, tailLength+1)}
	r.Reset(start)
	return r
}

func MoveRope(instructions []Instruction, tailLength int, observer Observer) {
	rope := NewRope(tailLength)
	observer.StartMoving(*rope)
	for _, instruction := range instructions {
		observer.AboutToExecute(instruction)
		for i := 0; i < instruction.amount; i += 1 {
			rope.Move(instruction.move())
			observer.StateChanged(*rope)
		}
		observer.FinishedInstruction()
	}
	observer.FinishedMoving()
}
