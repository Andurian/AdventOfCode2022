package main

import (
	"testing"
)

func TestScoreOutcome(t *testing.T) {
	cases := []struct {
		input    Outcome
		expected int
	}{
		{Win, 6},
		{Draw, 3},
		{Loss, 0},
	}

	for _, testcase := range cases {
		actual := testcase.input.Score()
		if actual != testcase.expected {
			t.Errorf("score(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}
