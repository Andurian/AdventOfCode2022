package rope

import (
	"andurian/adventofcode/2022/util"
	. "andurian/adventofcode/2022/util/point"
	"fmt"
	"strings"
)

type Instruction struct {
	direction string
	amount    int
}

func (i Instruction) String() string {
	return fmt.Sprintf("== %s %d ==", i.direction, i.amount)
}

func (i Instruction) move() Next {
	switch i.direction {
	case "U":
		return Up
	case "R":
		return Right
	case "D":
		return Down
	case "L":
		return Left
	default:
	}
	panic("Reached invalid state")
}

func InstructionFromString(s string) Instruction {
	tokens := strings.Split(strings.TrimSpace(s), " ")
	return Instruction{tokens[0], util.AtoiSafe(tokens[1])}
}
