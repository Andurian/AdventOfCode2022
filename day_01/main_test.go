package main

import (
	"testing"
)

const testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestSumTopNCalorieCarryingElves(t *testing.T) {
	cases := []struct {
		input    string
		n        int
		expected int
	}{
		{testInput, 1, 24000},
		{testInput, 3, 45000},
	}

	for _, testcase := range cases {
		actual := sumTopNCalorieCarryingElves(testcase.input, testcase.n)
		if actual != testcase.expected {
			t.Errorf("sumTopNCalorieCarryingElves(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}
