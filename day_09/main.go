package main

import (
	"andurian/adventofcode/2022/day_09/observer"
	"andurian/adventofcode/2022/day_09/rope"
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"image/color"
	"strings"
)

var Black = color.RGBA{0, 0, 0, 255}
var Gray = color.RGBA{128, 128, 128, 255}
var White = color.RGBA{255, 255, 255, 255}
var Red = color.RGBA{255, 0, 0, 255}

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

func writeImages(instructions []rope.Instruction, tailLength int) {
	topLeft, bottomRight := getRopeExtent(instructions, tailLength)
	imageWritingObserver := observer.NewImageWritingObserver(topLeft, bottomRight, Black, Red, White, 2, "imgs.zip")
	rope.MoveRope(instructions, tailLength, imageWritingObserver)
}

func main() {
	input := util.ReadSafe("input.txt")
	instructions := util.PreprocessTimed(func() []rope.Instruction { return parseInstructions(input) })

	util.ExecuteTimed(9, 1, func() int { return calcNumVisitedByTail(instructions, 1) })
	util.ExecuteTimed(9, 1, func() int { return calcNumVisitedByTail(instructions, 9) })

	//drawTailRouteAscii(instructions, 1)
	writeImages(instructions, 9)
}
