package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

type Instruction interface {
	CycleCount() int
	Execute(cpu *CPU)
}

type Noop struct{}

func (i *Noop) CycleCount() int { return 1 }
func (i *Noop) Execute(*CPU)    {}

type AddX struct {
	amount int
}

func (i *AddX) CycleCount() int  { return 2 }
func (i *AddX) Execute(cpu *CPU) { cpu.x += i.amount }

func InstuctionFromString(s string) Instruction {
	tokens := strings.Split(strings.TrimSpace(s), " ")
	switch tokens[0] {
	case "noop":
		return &Noop{}
	case "addx":
		return &AddX{util.AtoiSafe(tokens[1])}
	}
	panic("Invalid instruction")
}
