package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

func requiredChoice(elfChoice Shape, desiredOutcome Outcome) Shape {
	if desiredOutcome == Draw {
		return elfChoice
	}

	if desiredOutcome == Win {
		switch elfChoice {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	}

	switch elfChoice {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	}
	panic("Unreachable")
}

type InstructionTranslator func(byte, Shape) Shape

func gameFromString(input string, translateInstruction InstructionTranslator) Game {
	elfShape := ToShape(input[0])
	playerShape := translateInstruction(input[2], elfShape)
	return Game{playerShape, elfShape}
}

func translateToShape(c byte, _ Shape) Shape {
	return ToShape(c)
}

func translateToOutcome(c byte, elfChoice Shape) Shape {
	return requiredChoice(elfChoice, ToOutcome(c))
}

func gamesFromString(input string, translateInstruction InstructionTranslator) []Game {
	games := []Game{}
	for _, line := range strings.Split(input, "\n") {
		games = append(games, gameFromString(strings.TrimSpace(line), translateInstruction))
	}
	return games
}

func totalScore(input string, translateInstruction InstructionTranslator) int {
	score := 0
	for _, game := range gamesFromString(input, translateInstruction) {
		score += game.Score()
	}
	return score
}

func main() {
	input := util.ReadSafe("input.txt")

	util.ExecuteTimed(2, 1, func() int { return totalScore(input, translateToShape) })
	util.ExecuteTimed(2, 2, func() int { return totalScore(input, translateToOutcome) })
}
