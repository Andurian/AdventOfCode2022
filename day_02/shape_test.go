package main

import (
	"testing"
)

func TestToChoice(t *testing.T) {
	cases := []struct {
		input    byte
		expected Shape
	}{
		{'A', Rock},
		{'B', Paper},
		{'C', Scissors},
		{'X', Rock},
		{'Y', Paper},
		{'Z', Scissors},
	}

	for _, testcase := range cases {
		actual := ToShape(testcase.input)
		if actual != testcase.expected {
			t.Errorf("toChoice(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestScoreChoice(t *testing.T) {
	cases := []struct {
		input    Shape
		expected int
	}{
		{Rock, 1},
		{Paper, 2},
		{Scissors, 3},
	}

	for _, testcase := range cases {
		actual := testcase.input.Score()
		if actual != testcase.expected {
			t.Errorf("score(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}
