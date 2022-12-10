package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util/point"
	"image"
	"image/color"

	"github.com/muesli/gamut"
)

type ImageGeneratingObserver struct {
	topLeft     point.Point
	bottomRight point.Point

	width  int
	height int

	palette        color.Palette
	ropeColorRange [2]color.RGBA
	knotFilter     []bool
	pointSize      int

	HistoryObserver
	EmptyObserver
}

func (o *ImageGeneratingObserver) mapPointToImage(p point.Point) (x, y int) {
	y = (p.Row - o.topLeft.Row) * o.pointSize
	x = (p.Col - o.topLeft.Col) * o.pointSize
	return
}

func (o *ImageGeneratingObserver) paintRope(img *image.Paletted, knots []point.Point) {
	for i, p := range knots {
		if !o.knotFilter[i] {
			continue
		}
		px, py := o.mapPointToImage(p)
		for y := 0; y < o.pointSize; y += 1 {
			for x := 0; x < o.pointSize; x += 1 {
				img.SetColorIndex(px+x, py+y, uint8(i)+1)
			}
		}
	}
}

func (o *ImageGeneratingObserver) createEmptyImage() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, o.width, o.height), o.palette)
	for row := 0; row < o.height; row += 1 {
		for col := 0; col < o.width; col += 1 {
			img.SetColorIndex(col, row, 0)
		}
	}
	return img
}

func (o *ImageGeneratingObserver) StartMoving(r rope.Rope) {
	colors := gamut.Blends(o.ropeColorRange[0], o.ropeColorRange[1], len(r.Knots()))
	o.palette = append(o.palette, colors...)

	for len(o.knotFilter) < len(r.Knots()) {
		o.knotFilter = append(o.knotFilter, false)
	}

	o.HistoryObserver.StartMoving(r)
}

func (o *ImageGeneratingObserver) StateChanged(r rope.Rope) {
	o.HistoryObserver.StateChanged(r)
}

func (o *ImageGeneratingObserver) FinishedInstruction() {}

func (o *ImageGeneratingObserver) FinishedMoving() {}

func (o *ImageGeneratingObserver) Image() *image.Paletted {
	img := o.createEmptyImage()

	for _, r := range o.knotsInTime {
		o.paintRope(img, r)
	}

	return img
}

func NewImageGeneratingObserver(topLeft, bottomRight point.Point, knotFilter []bool, backgroundColor color.RGBA, ropeColorRange [2]color.RGBA, pointSize int) *ImageGeneratingObserver {
	height, width := point.AbsDistances(topLeft, bottomRight)
	width = (width + 1) * pointSize
	height = (height + 1) * pointSize

	return &ImageGeneratingObserver{
		topLeft:         topLeft,
		bottomRight:     bottomRight,
		width:           width,
		height:          height,
		ropeColorRange:  ropeColorRange,
		palette:         color.Palette{backgroundColor},
		knotFilter:      knotFilter,
		pointSize:       pointSize,
		HistoryObserver: HistoryObserver{},
		EmptyObserver:   EmptyObserver{},
	}
}
