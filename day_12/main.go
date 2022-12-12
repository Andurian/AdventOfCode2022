package main

import (
	"andurian/adventofcode/2022/util"
	"andurian/adventofcode/2022/util/point"
	"sort"
	"strings"

	"github.com/yourbasic/graph"
)

type HeightMap struct {
	graph       *graph.Mutable
	heightMap   [][]rune
	nodeIndices map[point.Point]int
	start       point.Point
	target      point.Point
}

func (h *HeightMap) ShortestPath() int {
	_, dist := graph.ShortestPath(h.graph, h.nodeIndices[h.start], h.nodeIndices[h.target])
	return int(dist)
}

func (h *HeightMap) ShortestPathFromBestStartingPos() int {
	distances := []int{}
	for k, v := range h.nodeIndices {
		height := h.heightMap[k.Row][k.Col]
		if height == 'S' || height == 'a' {
			_, dist := graph.ShortestPath(h.graph, v, h.nodeIndices[h.target])
			if dist != -1 {
				distances = append(distances, int(dist))
			}
		}
	}
	sort.IntSlice(distances).Sort()
	return distances[0]
}

func parseInput(s string) *HeightMap {
	lines := strings.Split(s, "\n")

	rows := len(lines)
	cols := len(strings.TrimSpace(lines[0]))

	heightMap := [][]rune{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		lineAsArray := make([]rune, len(line))
		for col, c := range line {
			lineAsArray[col] = c
		}
		heightMap = append(heightMap, lineAsArray)
	}
	heightAt := func(p point.Point) rune {
		return heightMap[p.Row][p.Col]
	}

	isInside := func(p point.Point) bool { return p.Row >= 0 && p.Row < rows && p.Col >= 0 && p.Col < cols }
	neighbors := func(p point.Point) []point.Point {
		ret := []point.Point{}
		for _, p := range point.Get4Neighbors(p) {
			if isInside(p) {
				ret = append(ret, p)
			}
		}
		return ret
	}

	g := graph.New(rows * cols)

	nodeIndices := make(map[point.Point]int)
	indexCounter := 0
	indexOf := func(p point.Point) int {
		if !util.MapContainsKey(nodeIndices, p) {
			nodeIndices[p] = indexCounter
			indexCounter += 1
		}
		return nodeIndices[p]
	}

	distance := func(current, next rune) int {
		if current == 'S' {
			if next == 'a' || next == 'b' {
				return 1
			}
			return -1
		}

		if next == 'E' {
			if current == 'z' || current == 'y' {
				return 1
			}
			return -1
		}

		if next-current <= 1 {
			return 1
		}
		return -1
	}

	var start, end point.Point
	for row := 0; row < rows; row += 1 {
		for col := 0; col < cols; col += 1 {
			currentPoint := point.Point{Row: row, Col: col}
			currentHeight := heightAt(currentPoint)
			for _, nextPoint := range neighbors(currentPoint) {
				nextHeight := heightAt(nextPoint)
				d := distance(currentHeight, nextHeight)
				if d != -1 {
					g.AddCost(indexOf(currentPoint), indexOf(nextPoint), int64(d))
				}
			}

			if currentHeight == 'S' {
				start = currentPoint
			}
			if currentHeight == 'E' {
				end = currentPoint
			}
		}
	}

	return &HeightMap{g, heightMap, nodeIndices, start, end}
}

func main() {
	input := util.ReadSafe("input.txt")

	heightMap := util.PreprocessTimed(func() *HeightMap { return parseInput(input) })
	util.ExecuteTimed(12, 1, func() int { return heightMap.ShortestPath() })
	util.ExecuteTimed(12, 2, func() int { return heightMap.ShortestPathFromBestStartingPos() })
}
