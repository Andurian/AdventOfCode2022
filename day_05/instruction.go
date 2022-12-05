package main

import (
	"andurian/adventofcode/2022/util"
	"fmt"
	"regexp"
)

type Instruction struct {
	amount int
	from   rune
	to     rune
}

func (i Instruction) String() string {
	return fmt.Sprintf("Move %d from %c to %c", i.amount, i.from, i.to)
}

var instructionRegex = regexp.MustCompile(`move (\d+) from (\S) to (\S)`)

func InstructionFromString(input string) Instruction {
	result := instructionRegex.FindStringSubmatch(input)
	return Instruction{util.AtoiSafe(result[1]), []rune(result[2])[0], []rune(result[3])[0]}
}
