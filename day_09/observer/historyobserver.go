package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util/point"
)

type HistoryObserver struct {
	knotsInTime [][]point.Point
}

func (o *HistoryObserver) addState(r rope.Rope) {
	knots := make([]point.Point, len(r.Knots()))
	copy(knots, r.Knots())
	o.knotsInTime = append(o.knotsInTime, knots)
}

func (o *HistoryObserver) StartMoving(r rope.Rope) {
	o.addState(r)
}

func (o *HistoryObserver) StateChanged(r rope.Rope) {
	o.addState(r)
}
