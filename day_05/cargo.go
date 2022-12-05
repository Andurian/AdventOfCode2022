package main

import (
	"andurian/adventofcode/2022/util"
	"fmt"
	"strings"
)

type Cargo struct {
	crates     map[rune][]rune
	sortedKeys []rune
}

func (c Cargo) Move(from rune, to rune) {
	stackFrom := c.crates[from]
	stackTo := c.crates[to]
	crate := stackFrom[len(stackFrom)-1]
	c.crates[from] = stackFrom[:len(stackFrom)-1]
	c.crates[to] = append(stackTo, crate)
}

func (c Cargo) Apply(instruction Instruction) {
	for i := 0; i < instruction.amount; i++ {
		c.Move(instruction.from, instruction.to)
	}
}

func (c Cargo) ApplyBatch(instruction Instruction) {
	stackFrom := c.crates[instruction.from]
	stackTo := c.crates[instruction.to]
	crates := stackFrom[len(stackFrom)-instruction.amount:]
	c.crates[instruction.from] = stackFrom[:len(stackFrom)-instruction.amount]
	c.crates[instruction.to] = append(stackTo, crates...)
}

func (c Cargo) String() string {
	ret := " "
	for _, key := range c.sortedKeys {
		ret += fmt.Sprintf("%c   ", key)
	}
	for stackHeight := 0; ; stackHeight++ {
		s := ""
		for _, key := range c.sortedKeys {
			stack := c.crates[key]
			if len(stack) <= stackHeight {
				s += "    "
			} else {
				s += fmt.Sprintf("[%c] ", stack[stackHeight])
			}
		}
		if strings.TrimSpace(s) != "" {
			ret += "\n" + s
		} else {
			break
		}
	}
	return ret
}

func (c Cargo) CratesOnTop() string {
	ret := ""
	for _, key := range c.sortedKeys {
		stack := c.crates[key]
		if len(stack) == 0 {
			ret += " "
		} else {
			ret += fmt.Sprintf("%c", stack[len(stack)-1])
		}
	}
	return ret
}

func (c Cargo) Copy() Cargo {
	sortedKeys := make([]rune, len(c.sortedKeys))
	copy(sortedKeys, c.sortedKeys)

	crates := make(map[rune][]rune)
	for key, value := range c.crates {
		crates[key] = make([]rune, len(value))
		copy(crates[key], value)
	}

	return Cargo{crates, sortedKeys}
}

func CargoFromString(input string) Cargo {
	startConfig := util.Reverse(strings.Split(input, "\n"))

	ids := make(map[int]rune)
	crates := make(map[rune][]rune)
	sortedKeys := []rune{}

	for index, stackId := range startConfig[0] {
		if stackId == ' ' {
			continue
		}
		sortedKeys = append(sortedKeys, stackId)
		ids[index] = stackId
		crates[stackId] = []rune{}
	}

	for _, line := range startConfig[1:] {
		runes := []rune(line)
		for stringIndex, stackId := range ids {
			crate := runes[stringIndex]
			if crate == ' ' {
				continue
			}
			crates[stackId] = append(crates[stackId], crate)

		}
	}

	return Cargo{crates, sortedKeys}
}
