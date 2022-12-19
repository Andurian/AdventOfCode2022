package main

import (
	"andurian/adventofcode/2022/util"
	"container/list"
	"math"
	"regexp"
	"strings"
)

type Resources struct {
	ore, clay, obsidian, geodes int
}

type Robots struct {
	ore, clay, obsidian, geodes int
}

func (r Robots) AddResources(res Resources) Resources {
	ret := res
	ret.ore += r.ore
	ret.clay += r.clay
	ret.obsidian += r.obsidian
	ret.geodes += r.geodes
	return ret
}

func InitialRobots() Robots {
	ret := Robots{}
	ret.ore = 1
	return ret
}

type State struct {
	minute    int
	resources Resources
	robots    Robots
}

type TrackedState struct {
	history []State
}

func (t TrackedState) Clone() TrackedState {
	c := make([]State, len(t.history))
	copy(c, t.history)
	return TrackedState{c}
}

func (t TrackedState) Current() State {
	return t.history[len(t.history)-1]
}

func (t TrackedState) Previous() State {
	if len(t.history) == 1 {
		return t.history[0]
	}
	return t.history[len(t.history)-2]
}

func (t TrackedState) Appended(s State) TrackedState {
	c := t.Clone()
	c.history = append(c.history, s)
	return c
}

type Blueprint struct {
	id            int
	oreRobot      struct{ ore int }
	clayRobot     struct{ ore int }
	obsidianRobot struct{ ore, clay int }
	geodeRobot    struct{ ore, obsidian int }
	maxOre        int
}

type Action func(s State) State

func (b Blueprint) CanBuildOre(state State) bool {
	return state.resources.ore >= b.oreRobot.ore
}

func (b Blueprint) CanBuildClay(state State) bool {
	return state.resources.ore >= b.clayRobot.ore
}

func (b Blueprint) CanBuildObsidian(state State) bool {
	return state.resources.ore >= b.obsidianRobot.ore && state.resources.clay >= b.obsidianRobot.clay
}

func (b Blueprint) CanBuildGeode(state State) bool {
	return state.resources.ore >= b.geodeRobot.ore && state.resources.obsidian >= b.geodeRobot.obsidian
}

func DidBuildOre(current, previous State) bool {
	return current.robots.ore == previous.robots.ore+1
}

func DidBuildClay(current, previous State) bool {
	return current.robots.clay == previous.robots.clay+1
}

func DidBuildObsidian(current, previous State) bool {
	return current.robots.obsidian == previous.robots.obsidian+1
}

func DidBuildGeode(current, previous State) bool {
	return current.robots.geodes == previous.robots.geodes+1
}

func DidBuildAnything(current, previous State) bool {
	return DidBuildOre(current, previous) || DidBuildClay(current, previous) || DidBuildObsidian(current, previous) || DidBuildGeode(current, previous)
}

func (b Blueprint) PossibleActions(currentTrackedState TrackedState) []TrackedState {
	ret := []TrackedState{}

	current := currentTrackedState.Current()
	previous := currentTrackedState.Previous()

	nextMinute := current.minute + 1
	nextResources := current.robots.AddResources(current.resources)

	// If we can build a geode robot, we build it
	if b.CanBuildGeode(current) {
		nextState := State{nextMinute, nextResources, current.robots}
		nextState.robots.geodes += 1
		nextState.resources.ore -= b.geodeRobot.ore
		nextState.resources.obsidian -= b.geodeRobot.obsidian
		ret = append(ret, currentTrackedState.Appended(nextState))
		return ret
	}

	// Only build a robot if the current number of robots do not build enough to build any other robot
	if b.CanBuildOre(current) && current.robots.ore < b.maxOre && !(b.CanBuildOre(previous) && !DidBuildAnything(current, previous)) {
		nextState := State{nextMinute, nextResources, current.robots}
		nextState.robots.ore += 1
		nextState.resources.ore -= b.oreRobot.ore
		ret = append(ret, currentTrackedState.Appended(nextState))
	}

	if b.CanBuildClay(current) && current.robots.clay < b.obsidianRobot.clay && !(b.CanBuildClay(previous) && !DidBuildAnything(current, previous)) {
		nextState := State{nextMinute, nextResources, current.robots}
		nextState.robots.clay += 1
		nextState.resources.ore -= b.clayRobot.ore
		ret = append(ret, currentTrackedState.Appended(nextState))
	}

	if b.CanBuildObsidian(current) && current.robots.obsidian < b.geodeRobot.obsidian && !(b.CanBuildObsidian(previous) && !DidBuildAnything(current, previous)) {
		nextState := State{nextMinute, nextResources, current.robots}
		nextState.robots.obsidian += 1
		nextState.resources.ore -= b.obsidianRobot.ore
		nextState.resources.clay -= b.obsidianRobot.clay
		ret = append(ret, currentTrackedState.Appended(nextState))
	}

	ret = append(ret, currentTrackedState.Appended(State{nextMinute, nextResources, current.robots}))

	return ret
}

func BlueprintFromString(s string) Blueprint {
	re := regexp.MustCompile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	t := re.FindStringSubmatch(s)
	c := util.AtoiSafe
	b := Blueprint{
		id:            c(t[1]),
		oreRobot:      struct{ ore int }{c(t[2])},
		clayRobot:     struct{ ore int }{c(t[3])},
		obsidianRobot: struct{ ore, clay int }{c(t[4]), c(t[5])},
		geodeRobot:    struct{ ore, obsidian int }{c(t[6]), c(t[7])},
	}

	b.maxOre = util.Max(b.oreRobot.ore, util.Max(b.clayRobot.ore, util.Max(b.obsidianRobot.ore, b.geodeRobot.ore)))
	return b
}

func (b Blueprint) MaxProducibleGeodes(minutes int) int {
	init := TrackedState{[]State{{0, Resources{}, InitialRobots()}}}
	max := math.MinInt

	current := list.New()
	current.PushBack(init)
	for current.Len() > 0 {
		next := list.New()
		maxGeodes := math.MinInt
		for c := current.Front(); c != nil; c = c.Next() {
			maxGeodes = util.Max(maxGeodes, c.Value.(TrackedState).Current().resources.geodes)
		}
		for cIter := current.Front(); cIter != nil; cIter = cIter.Next() {
			c := cIter.Value.(TrackedState)
			if c.Current().resources.geodes < maxGeodes-3 { // Probably arbitrary pruning value?
				continue
			}
			nextStates := b.PossibleActions(c)
			for _, s := range nextStates {
				if s.Current().minute == minutes {
					max = util.Max(max, s.Current().resources.geodes)
				} else {
					next.PushBack(s)
				}
			}
		}
		current = nil
		current = next
	}
	return max
}

func parse(input string) []Blueprint {
	ret := []Blueprint{}
	for _, line := range strings.Split(input, "\n") {
		ret = append(ret, BlueprintFromString(line))
	}
	return ret
}

func Task1(blueprints []Blueprint) int {
	ret := 0
	for _, b := range blueprints {
		ret += b.id * b.MaxProducibleGeodes(24)
	}
	return ret
}

func Task2(blueprints []Blueprint) int {
	ret := 1
	for _, b := range blueprints[:3] {
		ret *= b.MaxProducibleGeodes(32)
	}
	return ret
}

func main() {
	input := util.ReadSafe("input.txt")
	blueprints := util.PreprocessTimed(func() []Blueprint { return parse(input) })
	util.ExecuteTimed(19, 1, func() int { return Task1(blueprints) })
	util.ExecuteTimed(19, 2, func() int { return Task2(blueprints) })
}
