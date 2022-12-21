package main

import (
	"testing"
)

func TestAddFraction(t *testing.T) {
	cases := []struct {
		inputLeft, inputRight Fraction
		expected              Fraction
	}{
		{Fraction{1, 2}, Fraction{1, 2}, Fraction{1, 1}},
	}

	for _, testcase := range cases {
		actual := Add(testcase.inputLeft, testcase.inputRight)
		if actual != testcase.expected {
			t.Errorf("Add(%v, %v) should be %v but was %v", testcase.inputLeft, testcase.inputRight, testcase.expected, actual)
		}
	}
}
