package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"math"
)

type ExtentObserver struct {
	rowMin int
	rowMax int
	colMin int
	colMax int

	EmptyObserver
}

func (o *ExtentObserver) addPoint(p point.Point) {
	o.rowMax = util.Max(o.rowMax, p.Row)
	o.rowMin = util.Min(o.rowMin, p.Row)
	o.colMax = util.Max(o.colMax, p.Col)
	o.colMin = util.Min(o.colMin, p.Col)
}

func (o *ExtentObserver) addRope(r rope.Rope) {
	for _, k := range r.Knots() {
		o.addPoint(k)
	}
}

func (o *ExtentObserver) StartMoving(r rope.Rope) {
	o.addRope(r)
}

func (o *ExtentObserver) StateChanged(r rope.Rope) {
	o.addRope(r)
}

func (o *ExtentObserver) Extent() (topLeft, bottomRight point.Point) {
	topLeft = point.Point{Row: o.rowMin, Col: o.colMin}
	bottomRight = point.Point{Row: o.rowMax, Col: o.colMax}
	return
}

func NewExtentObserver() *ExtentObserver {
	return &ExtentObserver{
		rowMin: math.MaxInt,
		rowMax: math.MinInt,
		colMin: math.MaxInt,
		colMax: math.MinInt,
	}
}
