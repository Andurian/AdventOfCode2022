package main

import (
	"andurian/adventofcode/2022/util"
	"fmt"
	"sort"
	"strings"
)

type Order int

const (
	Before Order = iota
	Same
	After
)

type Packet interface {
	String() string
	OrderTo(Packet) Order
}

type Value struct {
	val int
}

func (v *Value) String() string {
	return fmt.Sprintf("%d", v.val)
}

func (v *Value) Wrapped() *List {
	l := NewList()
	l.AddPacket(v)
	return l
}

func (v *Value) OrderTo(p Packet) Order {
	switch w := p.(type) {
	case *Value:
		{
			if v.val == w.val {
				return Same
			} else if v.val < w.val {
				return Before
			}
			return After
		}
	case *List:
		{
			return v.Wrapped().OrderTo(p)
		}
	default:
		panic("Unknown type")
	}
}

func NewValue(s string) *Value {
	return &Value{util.AtoiSafe(s)}
}

type List struct {
	content []Packet
}

func (l *List) AddPacket(p Packet) {
	l.content = append(l.content, p)
}

func (l *List) String() string {
	s := "["
	for i, p := range l.content {
		if i != 0 {
			s += ","
		}
		s += p.String()
	}
	s += "]"
	return s
}

func (v *List) OrderTo(p Packet) Order {
	switch w := p.(type) {
	case *Value:
		{
			return v.OrderTo(w.Wrapped())
		}
	case *List:
		{
			lenLeft := len(v.content)
			lenRight := len(w.content)
			for i := 0; i < util.Min(lenLeft, lenRight); i += 1 {
				result := v.content[i].OrderTo(w.content[i])
				if result != Same {
					return result
				}
			}
			if lenLeft == lenRight {
				return Same
			} else if lenLeft < lenRight {
				return Before
			}
			return After
		}
	default:
		panic("Invalid Type")
	}
}

func NewList() *List {
	return &List{[]Packet{}}
}

func parsePacket(line string) Packet {
	root := NewList()
	stack := []*List{root}

	currentNumber := ""

	currentList := func() *List { return stack[len(stack)-1] }

	tryAddCurrentNumber := func() {
		if currentNumber != "" {
			currentList().AddPacket(NewValue(currentNumber))
			currentNumber = ""
		}
	}

	push := func() {
		temp := NewList()
		currentList().AddPacket(temp)
		stack = append(stack, temp)
	}

	pop := func() {
		stack = stack[:len(stack)-1]
	}

	for _, c := range line[1:] {
		if c == '[' {
			push()
		} else if c == ']' {
			tryAddCurrentNumber()
			pop()
		} else if c == ',' {
			tryAddCurrentNumber()
		} else {
			currentNumber += string(c)
		}
	}

	return root
}

func Task1(input string) int {
	pairs := strings.Split(input, "\n\n")
	sum := 0
	for i, pair := range pairs {
		lines := strings.Split(pair, "\n")
		p1 := parsePacket(lines[0])
		p2 := parsePacket(lines[1])
		if p1.OrderTo(p2) == Before {
			sum += i + 1
		}
	}
	return sum
}

func Task2(input string) int {
	packets := []Packet{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		packets = append(packets, parsePacket(line))
	}

	div1 := parsePacket("[[2]]")
	div2 := parsePacket("[[6]]")

	packets = append(packets, div1)
	packets = append(packets, div2)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].OrderTo(packets[j]) == Before
	})

	var p1, p2 int
	for i, p := range packets {
		if p.OrderTo(div1) == Same {
			p1 = i + 1
		} else if p.OrderTo(div2) == Same {
			p2 = i + 1
		}
	}
	return p1 * p2
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(13, 1, func() int { return Task1(input) })
	util.ExecuteTimed(13, 1, func() int { return Task2(input) })
}
