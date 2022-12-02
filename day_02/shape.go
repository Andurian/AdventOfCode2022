package main

import (
	"fmt"
)

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

func (c Shape) Score() int {
	switch c {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic("Unreachable")
}

func ToShape(c byte) Shape {
	switch c {
	case 'A':
		return Rock
	case 'X':
		return Rock
	case 'B':
		return Paper
	case 'Y':
		return Paper
	case 'C':
		return Scissors
	case 'Z':
		return Scissors
	}

	panic(fmt.Sprintf("Trying to convert invalid byte %q", c))
}
