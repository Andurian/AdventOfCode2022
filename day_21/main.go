package main

import (
	"andurian/adventofcode/2022/util"
	"regexp"
	"strings"
)

type Monkey interface {
	Run()
	GetNumber() Term
}

type BaseMonkey struct {
	name string
	out  chan Term
}

func (m *BaseMonkey) GetNumber() Term {
	return <-m.out
}

type NumberMonkey struct {
	number Term
	BaseMonkey
}

func (m *NumberMonkey) Run() {
	go func() {
		m.out <- m.number
	}()
}

type Operation func(left, right int) int

func plus(left, right int) int     { return left + right }
func minus(left, right int) int    { return left - right }
func multiply(left, right int) int { return left * right }
func divide(left, right int) int   { return left / right }

func OperationFromString(s string) Operation {
	switch s {
	case "+":
		return plus
	case "-":
		return minus
	case "*":
		return multiply
	case "/":
		return divide
	default:
		panic("Invalid Op")
	}
}

type TermOperation func(left, right Term) Term

func TermOperationFromString(s string) TermOperation {
	switch s {
	case "+":
		return AddTerm
	case "-":
		return SubtractTerm
	case "*":
		return MultiplyTerm
	case "/":
		return DivideTerm
	default:
		panic("Invalid Op")
	}
}

type CalculatingMonkey struct {
	inLeft, inRight chan Term
	operation       TermOperation
	BaseMonkey
}

func (m *CalculatingMonkey) Run() {
	go func() {
		// tLeft := <-m.inLeft
		// tRight := <-m.inRight
		// res := m.operation(tLeft, tRight)
		// fmt.Printf("%s: %v o %v = %v\n", m.name, tLeft, tRight, res)
		// m.out <- res
		m.out <- m.operation(<-m.inLeft, <-m.inRight)
	}()
}

type EqualityMonkey struct {
	inLeft, inRight chan Term
	BaseMonkey
}

func (m *EqualityMonkey) Run() {
	go func() {
		tLeft := <-m.inLeft
		tRight := <-m.inRight
		if tLeft.Degree() > tRight.Degree() {
			m.out <- SubtractTerm(tLeft, tRight)
		} else {
			m.out <- SubtractTerm(tRight, tLeft)
		}
	}()

}

type TermMonkey struct {
	BaseMonkey
}

func (m *TermMonkey) Run() {
	go func() {
		m.out <- NewTermDeg1FromInt(1, 0)
	}()
}

type Arena struct {
	outChannels map[string]chan Term
	monkeys     map[string]Monkey
}

func ArenaFromString(s string, addSpecialTask2Monkeys bool) *Arena {
	outChannels := make(map[string]chan Term)
	monkeys := make(map[string]Monkey)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		name := line[:4]
		outChannels[name] = make(chan Term, 1)
	}

	reNumber := regexp.MustCompile(`(\d+)`)
	reCalc := regexp.MustCompile(`(\w+) ([+\-*/]) (\w+)`)
	for _, line := range lines {
		name := line[:4]
		base := BaseMonkey{name: name, out: outChannels[name]}
		content := line[6:]
		matchNum := reNumber.FindStringSubmatch(content)
		if len(matchNum) != 0 {
			if name == "humn" && addSpecialTask2Monkeys {
				monkeys[name] = &TermMonkey{BaseMonkey: base}
			} else {
				monkeys[name] = &NumberMonkey{number: NewTermDeg0FromInt(util.AtoiSafe(matchNum[1])), BaseMonkey: base}
			}
			continue
		}

		matchCalc := reCalc.FindStringSubmatch(content)
		if len(matchCalc) != 0 {
			nameLeft := matchCalc[1]
			nameRight := matchCalc[3]
			op := TermOperationFromString(matchCalc[2])
			if name == "root" && addSpecialTask2Monkeys {
				monkeys[name] = &EqualityMonkey{inLeft: outChannels[nameLeft], inRight: outChannels[nameRight], BaseMonkey: base}
			} else {
				monkeys[name] = &CalculatingMonkey{inLeft: outChannels[nameLeft], inRight: outChannels[nameRight], operation: op, BaseMonkey: base}
			}
			continue
		}
		panic("Could not parse line")
	}
	return &Arena{outChannels: outChannels, monkeys: monkeys}
}

func (a *Arena) Run() {
	for _, v := range a.monkeys {
		v.Run()
	}
}

func Task1(input string) int {
	arena := ArenaFromString(input, false)
	arena.Run()
	return TryAsInt(arena.monkeys["root"].GetNumber().Get(0))
}

func Task2(input string) int {
	arena := ArenaFromString(input, true)
	arena.Run()
	equation := arena.monkeys["root"].GetNumber().(*TermDeg1)
	return TryAsInt(Solve(equation))
}

func main() {
	input := util.ReadSafe("input.txt")
	util.ExecuteTimed(21, 1, func() int { return Task1(input) })
	util.ExecuteTimed(21, 2, func() int { return Task2(input) })
}
