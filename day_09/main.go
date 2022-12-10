package main

import (
	"andurian/adventofcode/2022/day_09/observer"
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"strings"

	"github.com/muesli/gamut/palette"
)

var Black = color.RGBA{0, 0, 0, 255}
var Gray = color.RGBA{128, 128, 128, 255}
var White = color.RGBA{255, 255, 255, 255}
var Red = color.RGBA{255, 0, 0, 255}

func getColor(s string) color.RGBA {
	c, ok := palette.Wikipedia.Color(s)
	if !ok {
		panic(fmt.Sprintf("Cannot find color %q", s))
	}
	r, g, b, a := c.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func parseInstructions(input string) []rope.Instruction {
	lines := strings.Split(input, "\n")
	return util.Transform(lines, rope.InstructionFromString)
}

func calcNumVisitedByTail(instructions []rope.Instruction, tailLength int) int {
	tailObserver := observer.NewTailObserver()
	rope.MoveRope(instructions, tailLength, tailObserver)
	return tailObserver.NumVisitedByTail()
}

func getRopeExtent(instructions []rope.Instruction, tailLength int) (topLeft, bottomRight point.Point) {
	extentObserver := observer.NewExtentObserver()
	rope.MoveRope(instructions, tailLength, extentObserver)
	topLeft, bottomRight = extentObserver.Extent()
	return
}

func drawTailRouteAscii(instructions []rope.Instruction, tailLength int) {
	topLeft, bottomRight := getRopeExtent(instructions, tailLength)
	debugObserver := observer.NewDebugObserver(topLeft, bottomRight, observer.TailOnly)
	rope.MoveRope(instructions, tailLength, debugObserver)
}

func drawKnotRoutes(instructions []rope.Instruction, tailLength int, filename string) {
	topLeft, bottomRight := getRopeExtent(instructions, tailLength)
	knotFilter := make([]bool, tailLength+1)
	for i := range knotFilter {
		knotFilter[i] = true
	}
	knotFilter[len(knotFilter)-1] = true
	black := getColor("Black")
	rope1 := getColor("Magenta")
	rope2 := getColor("Dandelion")
	ropeColors := [2]color.RGBA{rope1, rope2}

	o := observer.NewImageGeneratingObserver(topLeft, bottomRight, knotFilter, black, ropeColors, 4)
	rope.MoveRope(instructions, tailLength, o)

	f, _ := os.Create(filename)
	defer f.Close()

	png.Encode(f, o.Image())
}

func drawKnotRoutesAnim(instructions []rope.Instruction, tailLengthRange [2]int, delay int, filename string) {
	topLeft, bottomRight := getRopeExtent(instructions, tailLengthRange[1])
	black := getColor("Black")
	rope1 := getColor("Magenta")
	rope2 := getColor("Dandelion")
	ropeColors := [2]color.RGBA{rope1, rope2}

	numImages := tailLengthRange[1] - tailLengthRange[0] + 1

	imgs := make([]*image.Paletted, numImages)
	delays := make([]int, numImages)

	for tailLength := tailLengthRange[0]; tailLength <= tailLengthRange[1]; tailLength += 1 {
		knotFilter := make([]bool, tailLength+1)
		for i := range knotFilter {
			knotFilter[i] = true
		}
		knotFilter[len(knotFilter)-1] = true

		o := observer.NewImageGeneratingObserver(topLeft, bottomRight, knotFilter, black, ropeColors, 2)
		rope.MoveRope(instructions, tailLength, o)

		i := tailLength - tailLengthRange[0]

		imgs[i] = o.Image()
		delays[i] = delay
	}

	f, _ := os.Create(filename)
	defer f.Close()

	gif.EncodeAll(f, &gif.GIF{
		Image: imgs,
		Delay: delays,
	})
}

func main() {
	input := util.ReadSafe("input.txt")
	instructions := util.PreprocessTimed(func() []rope.Instruction { return parseInstructions(input) })

	util.ExecuteTimed(9, 1, func() int { return calcNumVisitedByTail(instructions, 1) })
	util.ExecuteTimed(9, 1, func() int { return calcNumVisitedByTail(instructions, 9) })
}
