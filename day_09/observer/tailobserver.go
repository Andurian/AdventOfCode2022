package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util/point"

	mapset "github.com/deckarep/golang-set/v2"
)

type TailObserver struct {
	visitedByTail mapset.Set[point.Point]
	EmptyObserver
}

func (o *TailObserver) StartMoving(r rope.Rope) {
	o.visitedByTail.Add(r.TailEnd())
}

func (o *TailObserver) StateChanged(r rope.Rope) {
	o.visitedByTail.Add(r.TailEnd())
}

func (o *TailObserver) NumVisitedByTail() int {
	return o.visitedByTail.Cardinality()
}

func NewTailObserver() *TailObserver {
	return &TailObserver{mapset.NewSet[point.Point](), EmptyObserver{}}
}
