package main

import (
	"andurian/adventofcode/2022/util"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Item struct {
	start            uint64
	remainderClasses map[uint64]uint64
}

type Monkey struct {
	id      int
	items   []*Item
	op      func(uint64) uint64
	testDiv uint64
	onTrue  int
	onFalse int

	numInspected int
}

type Arena struct {
	monkeys []*Monkey
}

func (m *Monkey) InspectNext() (nextMonkexId int, item *Item) {
	item = m.items[0]
	m.items = m.items[1:]

	for k, v := range item.remainderClasses {
		item.remainderClasses[k] = m.op(v)
	}

	item.remainderClasses[m.testDiv] %= m.testDiv

	item.start = m.op(item.start)

	if item.start%m.testDiv == 0 {
		nextMonkexId = m.onTrue
	} else {
		nextMonkexId = m.onFalse
	}
	m.numInspected += 1
	return
}

func (m *Monkey) CanInspectNext() bool {
	return len(m.items) > 0
}

func (m *Monkey) CatchItem(item *Item) {
	m.items = append(m.items, item)
}

func (a *Arena) TryCleanup() {
	total := uint64(1)
	for _, m := range a.monkeys {
		total *= m.testDiv
	}
	for _, m := range a.monkeys {
		for i := range m.items {
			//m.items[i].start %= total
			if m.items[i].start%total == 0 {
				m.items[i].start = 0
			}
			// divsAll := true
			// for _, m2 := range a.monkeys {
			// 	if item.current%m2.testDiv != 0 {
			// 		divsAll = false
			// 		break
			// 	}
			// }
			// if divsAll {
			// 	m.items[i].current = m.items[i].start
			// }
		}
	}
}

func (a *Arena) Turn(m *Monkey) {
	for m.CanInspectNext() {
		nextMonkey, item := m.InspectNext()
		a.monkeys[nextMonkey].CatchItem(item)
	}
}

func (a *Arena) Round() {
	for _, m := range a.monkeys {
		a.Turn(m)
	}
	a.TryCleanup()
}

func NewArena(monkeys []*Monkey) *Arena {
	for _, m := range monkeys {
		for _, i := range m.items {
			for _, m2 := range monkeys {
				i.remainderClasses[m2.testDiv] = i.start
			}
		}
	}
	return &Arena{monkeys}
}

func (a *Arena) MonkeyBusiness() int {
	nums := make([]int, len(a.monkeys))
	for i, m := range a.monkeys {
		nums[i] = m.numInspected
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	return nums[0] * nums[1]
}

var expNumber = regexp.MustCompile(`[^\d](\d+)`)
var expOp = regexp.MustCompile(`[^\d*+]([*+]) (\d+)`)

func numberFromLine(line string) int {
	matches := expNumber.FindStringSubmatch(line)
	return util.AtoiSafe(matches[1])
}
func parseMonkey(s string) *Monkey {
	lines := strings.Split(s, "\n")

	matchesItems := expNumber.FindAllStringSubmatch(lines[1], -1)
	items := make([]*Item, len(matchesItems))
	for i, m := range matchesItems {
		x := uint64(util.AtoiSafe(m[1]))
		items[i] = &Item{x, make(map[uint64]uint64)}
	}

	matchesOp := expOp.FindStringSubmatch(lines[2])

	var op func(uint64) uint64
	if len(matchesOp) == 0 {
		op = func(item uint64) uint64 { return item * item } // small hack...
	} else {
		y := uint64(util.AtoiSafe(matchesOp[2]))
		if matchesOp[1] == "*" {
			op = func(item uint64) uint64 { return item * y }
		} else if matchesOp[1] == "+" {
			op = func(item uint64) uint64 { return item + y }
		} else {
			panic("invalid op")
		}
	}

	return &Monkey{id: numberFromLine(lines[0]),
		items:        items,
		op:           op,
		testDiv:      uint64(numberFromLine(lines[3])),
		onTrue:       numberFromLine(lines[4]),
		onFalse:      numberFromLine(lines[5]),
		numInspected: 0,
	}
}

func parse(s string) *Arena {
	ret := []*Monkey{}
	for _, c := range strings.Split(s, "\n\n") {
		monkey := parseMonkey(c)
		ret = append(ret, monkey)
	}
	return NewArena(ret)
}

func Task1(input string) int {
	arena := parse(input)
	for i := 0; i < 10000; i += 1 {
		arena.Round()
	}
	return arena.MonkeyBusiness()
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(11, 1, func() int { return Task1(input) })
}
