package main

import (
	"andurian/adventofcode/2022/util"
	"math"
	"regexp"
	"strings"

	"github.com/yourbasic/graph"
)

type DistanceMap struct {
	distances map[string]map[string]int
}

func (d *DistanceMap) Distance(from, to string) int {
	return d.distances[from][to]
}

func NewDistanceMap(adjacencyList map[string][]string, indices map[string]int) *DistanceMap {
	g := graph.New(len(indices))
	for start, targets := range adjacencyList {
		for _, target := range targets {
			g.AddCost(indices[start], indices[target], 1)
		}
	}

	getName := func(index int) string {
		for k, v := range indices {
			if v == index {
				return k
			}
		}
		panic("invalid index")
	}

	distances := make(map[string]map[string]int)
	for k, v := range indices {
		distances[k] = make(map[string]int)
		targets, dist := graph.ShortestPaths(g, v)
		for i := 0; i < len(targets); i += 1 {
			target := getName(i)
			if target != k {
				distances[k][getName(i)] = int(dist[i])
			}
		}
	}

	return &DistanceMap{distances}
}

type Network struct {
	pressures               map[string]int
	currentAchievedPressure int
	currentTime             int
	currentValve            string
	distances               *DistanceMap
}

func (n *Network) MaxRemainingFlow() int {
	sum := 0
	for _, v := range n.pressures {
		sum += n.currentTime * v
	}
	return sum
}

func (n *Network) CanOpen(valve string) bool {
	dTime := n.distances.Distance(n.currentValve, valve) + 1
	return dTime <= n.currentTime
}

func (n *Network) ReachablePressurizedValves() []string {
	ret := []string{}
	for k, v := range n.pressures {
		if v != 0 && n.CanOpen(k) {
			ret = append(ret, k)
		}
	}
	return ret
}

func (n *Network) Clone() *Network {
	return &Network{
		util.CopyMap(n.pressures),
		n.currentAchievedPressure,
		n.currentTime,
		n.currentValve,
		n.distances,
	}
}

func (n *Network) OpenValve(valve string) *Network {
	dTime := n.distances.Distance(n.currentValve, valve) + 1
	stepped := n.Clone()
	stepped.pressures[valve] = 0
	stepped.currentTime = n.currentTime - dTime
	stepped.currentAchievedPressure = n.currentAchievedPressure + stepped.currentTime*n.pressures[valve]
	stepped.currentValve = valve
	return stepped
}

func NewNetwork(pressures map[string]int, distances *DistanceMap) *Network {
	return &Network{pressures, 0, 30, "AA", distances}
}

func parse(s string) *Network {
	reValves := regexp.MustCompile(`([A-Z]{2})`)
	reFlow := regexp.MustCompile(`(\d+)`)

	connections := make(map[string][]string)
	flows := make(map[string]int)
	indices := make(map[string]int)

	for i, line := range strings.Split(s, "\n") {
		valves := reValves.FindAllString(line, -1)
		flow := util.AtoiSafe(reFlow.FindString(line))

		v := valves[0]

		connections[v] = []string{}
		connections[v] = append(connections[v], valves[1:]...)

		flows[v] = flow
		indices[v] = i
	}

	return NewNetwork(flows, NewDistanceMap(connections, indices))
}

func Task1(s string) {
	network := parse(s)

	options := []*Network{network}

	bestSoFar := []*Network{}

	for len(options) > 0 {
		newOptions := []*Network{}
		for _, option := range options {
			nextSteps := option.ReachablePressurizedValves()
			if len(nextSteps) == 0 {
				bestSoFar = append(bestSoFar, option)
				continue
			}
			for _, valve := range nextSteps {
				newOptions = append(newOptions, option.OpenValve(valve))
			}
		}
		maxCurrentFlow := math.MinInt
		for _, option := range newOptions {
			maxCurrentFlow = util.Max(maxCurrentFlow, option.currentAchievedPressure)
		}
		options = []*Network{}
		for _, option := range newOptions {
			if option.currentAchievedPressure+option.MaxRemainingFlow() >= maxCurrentFlow {
				options = append(options, option)
			}
		}
		println(len(options))
	}

	maxCurrentFlow := math.MinInt
	for _, option := range bestSoFar {
		maxCurrentFlow = util.Max(maxCurrentFlow, option.currentAchievedPressure)
	}

	println(len(bestSoFar))
	println(maxCurrentFlow)
}

func main() {
	input := util.ReadSafe("input.txt")
	Task2(input)
}
