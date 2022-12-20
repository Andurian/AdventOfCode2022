package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/listutils"
	"container/list"
	"fmt"
	"strings"
)

func walkNaive(l *list.List, from *list.Element, offset int) *list.Element {
	offset %= l.Len()
	if offset == 0 {
		return from
	}

	var advance func(*list.Element) *list.Element
	if offset < 0 {
		advance = func(current *list.Element) (next *list.Element) {
			next = current.Prev()
			if next == nil {
				next = l.Back()
			}
			return
		}
	} else {
		advance = func(current *list.Element) (next *list.Element) {
			next = current.Next()
			if next == nil {
				next = l.Front()
			}
			return
		}
	}

	current := from

	for c := 0; c < util.Abs(offset); c += 1 {
		current = advance(current)
	}

	return current
}

func decrypt(l *list.List, count int) *list.List {
	decrypted := list.New()
	decrypted.PushBackList(l)

	decryptionQueue := make([]*list.Element, decrypted.Len())
	for elem, i := decrypted.Front(), 0; elem != nil; elem, i = elem.Next(), i+1 {
		decryptionQueue[i] = elem
	}

	for c := 0; c < count; c += 1 {
		for i, elem := range decryptionQueue {
			offset := elem.Value.(int)
			if offset == 0 {
				continue
			}

			var previous *list.Element

			if offset > 0 {
				previous = elem.Prev()
				if previous == nil {
					previous = decrypted.Back()
				}
			} else {
				previous = elem.Next()
				if previous == nil {
					previous = decrypted.Front()
				}
			}

			decrypted.Remove(elem)

			targetElement := walkNaive(decrypted, previous, offset)

			if offset > 0 {
				if targetElement == decrypted.Back() {
					decryptionQueue[i] = decrypted.PushFront(elem.Value.(int))
				} else {
					decryptionQueue[i] = decrypted.InsertAfter(elem.Value.(int), targetElement)
				}
			} else {
				if targetElement == decrypted.Front() {
					decryptionQueue[i] = decrypted.PushBack(elem.Value.(int))
				} else {
					decryptionQueue[i] = decrypted.InsertBefore(elem.Value.(int), targetElement)
				}
			}
		}
	}

	return decrypted
}

func coordinateChecksum(l *list.List) int {
	coords := [...]int{1000, 2000, 3000}
	sum := 0
	startIndex := listutils.IndexOfValue(l, 0)
	for _, c := range coords {
		element := listutils.ElementAt(l, startIndex+c)
		sum += element.Value.(int)
	}
	return sum
}

func Task1(encrypted *list.List) int {
	return coordinateChecksum(decrypt(encrypted, 1))
}

func Task2(encrypted *list.List, key int) int {
	keyApplied := list.New()
	for elem := encrypted.Front(); elem != nil; elem = elem.Next() {
		keyApplied.PushBack(elem.Value.(int) * key)
	}
	return coordinateChecksum(decrypt(keyApplied, 10))
}

func parse(input string) *list.List {
	ret := list.New()
	for _, line := range strings.Split(input, "\n") {
		ret.PushBack(util.AtoiSafe(line))
	}
	return ret
}

func printList(l *list.List) {
	s := "["
	for elem := l.Front(); elem != nil; elem = elem.Next() {
		if elem != l.Front() {
			s += ", "
		}
		s += fmt.Sprintf("%v", elem.Value)
	}
	s += "]"
	println(s)
}

func main() {
	input := util.ReadSafe("input.txt")
	encrypted := util.PreprocessTimed(func() *list.List { return parse(input) })
	util.ExecuteTimed(20, 1, func() int { return Task1(encrypted) })
	util.ExecuteTimed(20, 1, func() int { return Task2(encrypted, 811589153) })
}
