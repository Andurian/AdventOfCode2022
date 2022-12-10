package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

func makeInstructions(input string) []Instruction {
	return util.Transform(strings.Split(input, "\n"), InstuctionFromString)
}

func Task1(instructions []Instruction) int {
	cpu := NewCPU()

	sum := 0
	callback := func(cycle, x int) {
		sum += cycle * x
	}
	cpu.AddCallback(CallbackId{20, 40}, callback)
	cpu.ExecuteMultiple(instructions)

	return sum
}

func Task2(rows, cols int, instructions []Instruction) string {
	cpu := NewCPU()
	ret := ""
	callback := func(cycle, x int) {
		if cycle > rows*cols {
			return
		}

		row := (cycle - 1) / cols
		col := (cycle - 1) % rows

		if col == 0 && row != 0 {
			ret += "\n"
		}

		if col == x-1 || col == x || col == x+1 {
			ret += "#"
		} else {
			ret += "."
		}
	}
	cpu.AddCallback(CallbackId{1, 1}, callback)
	cpu.ExecuteMultiple(instructions)

	return ret
}

func main() {
	input := util.ReadSafe("input.txt")

	instructions := util.PreprocessTimed(func() []Instruction { return makeInstructions(input) })

	util.ExecuteTimed(10, 1, func() int { return Task1(instructions) })
	util.ExecuteTimedStringMultiline(10, 2, func() string { return Task2(40, 6, instructions) })
}
