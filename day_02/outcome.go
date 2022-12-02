package main

import (
	"fmt"
)

type Outcome int

const (
	Win Outcome = iota
	Draw
	Loss
)

func (o Outcome) Score() int {
	switch o {
	case Win:
		return 6
	case Draw:
		return 3
	case Loss:
		return 0
	}
	panic("Unreachable")
}

func ToOutcome(o byte) Outcome {
	switch o {
	case 'X':
		return Loss
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic(fmt.Sprintf("Trying to convert invalid byte %q", o))
}
