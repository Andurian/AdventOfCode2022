package main

import (
	mapset "github.com/deckarep/golang-set/v2"
)

type Game struct {
	player Shape
	elf    Shape
}

func (g Game) Outcome() Outcome {
	if g.player == g.elf {
		return Draw
	}

	winningCombinations := mapset.NewSet(Game{Rock, Scissors}, Game{Paper, Rock}, Game{Scissors, Paper})
	if winningCombinations.Contains(g) {
		return Win
	}

	return Loss
}

func (g Game) Score() int {
	return g.player.Score() + g.Outcome().Score()
}
