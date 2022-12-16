package main

import (
	"andurian/adventofcode/2022/util"
	"math"
	"regexp"
	"strings"
)

type NetworkWithElephant struct {
	pressures               map[string]int
	currentAchievedPressure int
	currentTimeHuman        int
	currentTimeElephant     int
	currentValveHuman       string
	currentValveElephant    string
	distances               *DistanceMap
}

type id struct {
	currentAchievedPressure int
	currentTimeHuman        int
	currentTimeElephant     int
	currentValveHuman       string
	currentValveElephant    string
}

func (n *NetworkWithElephant) Id() id {
	return id{n.currentAchievedPressure, n.currentTimeHuman, n.currentTimeElephant, n.currentValveHuman, n.currentValveElephant}
}

func (n *NetworkWithElephant) IdMirrored() id {
	return id{n.currentAchievedPressure, n.currentTimeElephant, n.currentTimeHuman, n.currentValveElephant, n.currentValveHuman}
}

func (n *NetworkWithElephant) MaxRemainingFlow() int {
	sum := 0
	t := util.Max(n.currentTimeElephant, n.currentTimeHuman)
	for _, v := range n.pressures {
		sum += t * v
	}
	return sum
}

func (n *NetworkWithElephant) CanOpenHuman(valve string) bool {
	dTime := n.distances.Distance(n.currentValveHuman, valve) + 1
	return dTime <= n.currentTimeHuman
}

func (n *NetworkWithElephant) CanOpenElephant(valve string) bool {
	dTime := n.distances.Distance(n.currentValveElephant, valve) + 1
	return dTime <= n.currentTimeElephant
}

func (n *NetworkWithElephant) ReachablePressurizedValvesHuman() []string {
	ret := []string{}
	for k, v := range n.pressures {
		if v != 0 && n.CanOpenHuman(k) {
			ret = append(ret, k)
		}
	}
	return ret
}

func (n *NetworkWithElephant) ReachablePressurizedValvesElephant() []string {
	ret := []string{}
	for k, v := range n.pressures {
		if v != 0 && n.CanOpenElephant(k) {
			ret = append(ret, k)
		}
	}
	return ret
}

func (n *NetworkWithElephant) Clone() *NetworkWithElephant {
	return &NetworkWithElephant{
		util.CopyMap(n.pressures),
		n.currentAchievedPressure,
		n.currentTimeHuman,
		n.currentTimeElephant,
		n.currentValveHuman,
		n.currentValveElephant,
		n.distances,
	}
}

func (n *NetworkWithElephant) OpenValveHuman(valve string) *NetworkWithElephant {
	dTime := n.distances.Distance(n.currentValveHuman, valve) + 1
	stepped := n.Clone()
	stepped.pressures[valve] = 0
	stepped.currentTimeHuman = n.currentTimeHuman - dTime
	stepped.currentAchievedPressure = n.currentAchievedPressure + stepped.currentTimeHuman*n.pressures[valve]
	stepped.currentValveHuman = valve
	return stepped
}

func (n *NetworkWithElephant) OpenValveElephant(valve string) *NetworkWithElephant {
	dTime := n.distances.Distance(n.currentValveElephant, valve) + 1
	stepped := n.Clone()
	stepped.pressures[valve] = 0
	stepped.currentTimeElephant = n.currentTimeElephant - dTime
	stepped.currentAchievedPressure = n.currentAchievedPressure + stepped.currentTimeElephant*n.pressures[valve]
	stepped.currentValveElephant = valve
	return stepped
}

func NewNetworkWithElephant(pressures map[string]int, distances *DistanceMap) *NetworkWithElephant {
	return &NetworkWithElephant{pressures, 0, 26, 26, "AA", "AA", distances}
}

func parseWithElephant(s string) *NetworkWithElephant {
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

	return NewNetworkWithElephant(flows, NewDistanceMap(connections, indices))
}

func Task2(s string) {
	network := parseWithElephant(s)

	options := []*NetworkWithElephant{network}

	bestSoFar := []*NetworkWithElephant{}

	allKeys := make(map[id]bool)

	for len(options) > 0 {
		newOptions := []*NetworkWithElephant{}

		for _, option := range options {
			nextSteps := option.ReachablePressurizedValvesHuman()
			if len(nextSteps) == 0 {
				bestSoFar = append(bestSoFar, option)
				continue
			}
			for _, valve := range nextSteps {
				newOption := option.OpenValveHuman(valve)
				if _, value := allKeys[newOption.Id()]; !value {
					allKeys[newOption.Id()] = true
					allKeys[newOption.IdMirrored()] = true
					newOptions = append(newOptions, newOption)
				}
			}

			nextSteps = option.ReachablePressurizedValvesElephant()
			if len(nextSteps) == 0 {
				bestSoFar = append(bestSoFar, option)
				continue
			}
			for _, valve := range nextSteps {
				newOption := option.OpenValveElephant(valve)
				if _, value := allKeys[newOption.Id()]; !value {
					allKeys[newOption.Id()] = true
					allKeys[newOption.IdMirrored()] = true
					newOptions = append(newOptions, newOption)
				}
			}
		}

		maxCurrentFlow := math.MinInt
		for _, option := range newOptions {
			maxCurrentFlow = util.Max(maxCurrentFlow, option.currentAchievedPressure)
		}

		options = []*NetworkWithElephant{}
		for _, option := range newOptions {
			if option.currentAchievedPressure+option.MaxRemainingFlow() >= maxCurrentFlow {
				options = append(options, option)
			}
		}

		//options = removeDuplicates(options)
		println(len(options))
	}

	maxCurrentFlow := math.MinInt
	for _, option := range bestSoFar {
		maxCurrentFlow = util.Max(maxCurrentFlow, option.currentAchievedPressure)
	}

	println(len(bestSoFar))
	println(maxCurrentFlow)
}
