package listutils

import (
	"andurian/adventofcode/2022/util"
	"container/list"
	"fmt"
)

func CleanIndex(l *list.List, i int) int {
	i %= l.Len()
	if i < 0 {
		i += l.Len()
	}
	return i
}

func IndexOf(l *list.List, target *list.Element) int {
	for elem, i := l.Front(), 0; elem != nil; elem, i = elem.Next(), i+1 {
		if elem == target {
			return i
		}
	}
	panic("Index Of Element not found")
}

func IndexOfValue(l *list.List, target int) int {
	for elem, i := l.Front(), 0; elem != nil; elem, i = elem.Next(), i+1 {
		if elem.Value.(int) == target {
			return i
		}
	}
	panic("Index of Value Element not found")
}

func ElementAt(l *list.List, target int) *list.Element {
	target = CleanIndex(l, target)
	if target >= l.Len() {
		panic("listutils.ElementAt: Index not found")
	}
	for elem, i := l.Front(), 0; elem != nil; elem, i = elem.Next(), i+1 {
		if i == target {
			return elem
		}
	}
	panic("listutils.ElementAt: Index not found")
}

func ElementAtNaive(l *list.List, target int) *list.Element {
	if target == 0 {
		return l.Front()
	}

	var advance func(*list.Element, int) (*list.Element, int)
	if target < 0 {
		advance = func(current *list.Element, i int) (next *list.Element, j int) {
			next = current.Prev()
			j = i - 1
			if next == nil {
				next = l.Back()
				j = l.Len() - 1
			}
			return
		}
	} else {
		advance = func(current *list.Element, i int) (next *list.Element, j int) {
			next = current.Next()
			j = i + 1
			if next == nil {
				next = l.Front()
				j = 0
			}
			return
		}
	}

	current := l.Front()
	i := 0

	for c := 0; c < util.Abs(target); c += 1 {
		current, i = advance(current, i)
	}

	return current
}

func ToString(l *list.List) string {
	s := "["
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		if elem != l.Front() {
			s += ", "
		}
		s += fmt.Sprintf("%v", elem.Value)
	}
	s += "]"
	return s
}
