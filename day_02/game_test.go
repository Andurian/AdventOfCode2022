package main

import (
	"testing"
)

func TestGameOutcome(t *testing.T) {
	cases := []struct {
		input    Game
		expected Outcome
	}{
		{Game{Rock, Rock}, Draw},
		{Game{Rock, Paper}, Loss},
		{Game{Rock, Scissors}, Win},
		{Game{Paper, Rock}, Win},
		{Game{Paper, Paper}, Draw},
		{Game{Paper, Scissors}, Loss},
		{Game{Scissors, Rock}, Loss},
		{Game{Scissors, Paper}, Win},
		{Game{Scissors, Scissors}, Draw},
	}

	for _, testcase := range cases {
		actual := testcase.input.Outcome()
		if actual != testcase.expected {
			t.Errorf("Outcome(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestGameScore(t *testing.T) {
	cases := []struct {
		input    Game
		expected int
	}{
		{Game{Rock, Rock}, 1 + 3},
		{Game{Rock, Paper}, 1 + 0},
		{Game{Rock, Scissors}, 1 + 6},
		{Game{Paper, Rock}, 2 + 6},
		{Game{Paper, Paper}, 2 + 3},
		{Game{Paper, Scissors}, 2 + 0},
		{Game{Scissors, Rock}, 3 + 0},
		{Game{Scissors, Paper}, 3 + 6},
		{Game{Scissors, Scissors}, 3 + 3},
	}

	for _, testcase := range cases {
		actual := testcase.input.Score()
		if actual != testcase.expected {
			t.Errorf("Score(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}
