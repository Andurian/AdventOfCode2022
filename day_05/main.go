package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

func preprocess(input string) (Cargo, []Instruction) {
	splits := strings.Split(input, "\n\n")

	cargo := CargoFromString(splits[0])
	instructions := util.Transform(strings.Split(splits[1], "\n"), InstructionFromString)

	return cargo, instructions
}

func topCratesAfterInstuctions(cargo Cargo, instructions []Instruction) string {
	for _, i := range instructions {
		cargo.Apply(i)
	}
	return cargo.CratesOnTop()
}

func topCratesAfterBatchInstuctions(cargo Cargo, instructions []Instruction) string {
	for _, i := range instructions {
		cargo.ApplyBatch(i)
	}
	return cargo.CratesOnTop()
}

func main() {
	input := util.ReadSafe("input.txt")

	cargo, instructions := util.PreprocessTimedPair(func() (Cargo, []Instruction) { return preprocess(input) })

	util.ExecuteTimedString(5, 1, func() string { return topCratesAfterInstuctions(cargo.Copy(), instructions) })
	util.ExecuteTimedString(5, 1, func() string { return topCratesAfterBatchInstuctions(cargo.Copy(), instructions) })
}
