package main

import (
	"testing"
)

func TestRequiredChoice(t *testing.T) {
	cases := []struct {
		inputElfShape       Shape
		inputDesiredOutcome Outcome
		expected            Shape
	}{
		{Rock, Win, Paper},
		{Rock, Draw, Rock},
		{Rock, Loss, Scissors},
		{Paper, Win, Scissors},
		{Paper, Draw, Paper},
		{Paper, Loss, Rock},
		{Scissors, Win, Rock},
		{Scissors, Draw, Scissors},
		{Scissors, Loss, Paper},
	}
	for _, testcase := range cases {
		actual := requiredChoice(testcase.inputElfShape, testcase.inputDesiredOutcome)
		if actual != testcase.expected {
			t.Errorf("requiredChoice(%q, %q) should be %q but was %q", testcase.inputElfShape, testcase.inputDesiredOutcome, testcase.expected, actual)
		}
	}
}

func TestGameFromString(t *testing.T) {
	cases := []struct {
		input             string
		expectedByShape   Game
		expectedByOutcome Game
	}{
		{"A Y", Game{Paper, Rock}, Game{Rock, Rock}},
		{"B X", Game{Rock, Paper}, Game{Rock, Paper}},
		{"C Z", Game{Scissors, Scissors}, Game{Rock, Scissors}},
	}

	for _, testcase := range cases {
		actualByShape := gameFromString(testcase.input, translateToShape)
		if actualByShape != testcase.expectedByShape {
			t.Errorf("gamesFromString(%q, translateToShape) should be %d but was %d", testcase.input, testcase.expectedByShape, actualByShape)
		}

		actualByOutcome := gameFromString(testcase.input, translateToOutcome)
		if actualByOutcome != testcase.expectedByOutcome {
			t.Errorf("gamesFromString(%q, translateToOutcome) should be %d but was %d", testcase.input, testcase.expectedByOutcome, actualByOutcome)
		}
	}
}

func TestTotalScore(t *testing.T) {
	cases := []struct {
		input             string
		expectedByShape   int
		expectedByOutcome int
	}{
		{`A Y
		B X
		C Z`, 15, 12},
	}

	for _, testcase := range cases {
		actual := totalScore(testcase.input, translateToShape)
		if actual != testcase.expectedByShape {
			t.Errorf("totalScore(%q, translateToShape) should be %d but was %d", testcase.input, testcase.expectedByShape, actual)
		}

		actual = totalScore(testcase.input, translateToOutcome)
		if actual != testcase.expectedByOutcome {
			t.Errorf("totalScore(%q, translateToOutcome) should be %d but was %d", testcase.input, testcase.expectedByOutcome, actual)
		}
	}
}
