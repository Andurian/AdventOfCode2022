package observer

import (
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util/point"
	. "andurian/adventofcode/2022/util/point"
	"archive/zip"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	mapset "github.com/deckarep/golang-set/v2"
)

type ImageWritingObserver struct {
	topLeft     Point
	bottomRight Point

	width  int
	height int

	backgroundColor color.RGBA
	ropeColor       color.RGBA
	tailColor       color.RGBA
	pointSize       int

	writeCount int
	file       *os.File
	writer     *zip.Writer

	knots         []Point
	visitedByTail mapset.Set[Point]

	EmptyObserver
}

func (o *ImageWritingObserver) mapPointToImage(p point.Point) (x, y int) {
	y = (p.Row - o.topLeft.Row) * o.pointSize
	x = (p.Col - o.topLeft.Col) * o.pointSize
	return
}

func (o *ImageWritingObserver) paintTail(img *image.RGBA) {
	for _, p := range o.visitedByTail.ToSlice() {
		px, py := o.mapPointToImage(p)
		for y := 0; y < o.pointSize; y += 1 {
			for x := 0; x < o.pointSize; x += 1 {
				img.Set(px+x, py+y, o.tailColor)
			}
		}
	}
}

func (o *ImageWritingObserver) paintRope(img *image.RGBA) {
	for _, p := range o.knots {
		px, py := o.mapPointToImage(p)
		for y := 0; y < o.pointSize; y += 1 {
			for x := 0; x < o.pointSize; x += 1 {
				img.Set(px+x, py+y, o.ropeColor)
			}
		}
	}
}

func (o *ImageWritingObserver) createEmptyImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, o.width, o.height))
	for row := 0; row < o.height; row += 1 {
		for col := 0; col < o.width; col += 1 {
			img.Set(col, row, o.backgroundColor)
		}
	}
	return img
}

func (o *ImageWritingObserver) writeImage(filename string, includeRope bool) {
	img := o.createEmptyImage()
	o.paintTail(img)
	if includeRope {
		o.paintRope(img)
	}
	f, _ := o.writer.Create(filename)
	png.Encode(f, img)
}

func (o *ImageWritingObserver) nextFilename() string {
	ret := fmt.Sprintf("img_%06d.png", o.writeCount)
	o.writeCount += 1
	return ret
}

func (o *ImageWritingObserver) StartMoving(r rope.Rope) {
	o.knots = r.Knots()
	o.visitedByTail.Add(r.TailEnd())

	filename := o.nextFilename()
	fmt.Printf("writing: %s\n", filename)
	o.writeImage(filename, true)

}

func (o *ImageWritingObserver) StateChanged(r rope.Rope) {
	o.knots = r.Knots()
	o.visitedByTail.Add(r.TailEnd())
}

func (o *ImageWritingObserver) FinishedInstruction() {
	filename := o.nextFilename()
	fmt.Printf("writing: %s\n", filename)
	o.writeImage(filename, true)
}

func (o *ImageWritingObserver) FinishedMoving() {
	o.writeImage("tail.png", true)
	o.writer.Close()
}

func NewImageWritingObserver(topLeft, bottomRight Point, backgroundColor, ropeColor, tailColor color.RGBA, pointSize int, filename string) *ImageWritingObserver {
	height, width := point.AbsDistances(topLeft, bottomRight)
	width = (width + 1) * pointSize
	height = (height + 1) * pointSize

	file, _ := os.Create(filename)
	writer := zip.NewWriter(file)

	return &ImageWritingObserver{
		topLeft:         topLeft,
		bottomRight:     bottomRight,
		width:           width,
		height:          height,
		backgroundColor: backgroundColor,
		ropeColor:       ropeColor,
		tailColor:       tailColor,
		pointSize:       pointSize,
		writeCount:      0,
		file:            file,
		writer:          writer,
		visitedByTail:   mapset.NewSet[Point](),
		EmptyObserver:   EmptyObserver{},
	}
}
