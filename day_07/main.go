package main

import (
	"andurian/adventofcode/2022/util"
)

func FindDirs(root *Directory, predicate func(d *Directory) bool) []*Directory {
	foundDirs := []*Directory{}

	root.traverse(func(dInner *Directory) {
		if predicate(dInner) {
			foundDirs = append(foundDirs, dInner)
		}
	})

	return foundDirs
}

func Task1(root *Directory) int {
	dirs := FindDirs(root, func(d *Directory) bool { return d.calculateSize() <= 100000 })
	return util.AccumulateFunc(dirs, func(d *Directory) int { return d.calculateSize() })
}

func Task2(root *Directory) int {
	totalSize := 70000000
	updateSize := 30000000
	availableSize := totalSize - root.size
	requiredSize := updateSize - availableSize

	bestCandidate := root
	root.traverse(func(d *Directory) {
		if d.size >= requiredSize && d.size < bestCandidate.size {
			bestCandidate = d
		}
	})
	return bestCandidate.size
}

func main() {
	input := util.ReadSafe("input.txt")
	root := util.PreprocessTimed(func() *Directory { return MakeDirectory(input) })

	util.ExecuteTimed(7, 1, func() int { return Task1(root) })
	util.ExecuteTimed(7, 2, func() int { return Task2(root) })
}
