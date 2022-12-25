package main

import (
	"testing"
)

func TestSnafuIota(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}
	for _, testcase := range cases {
		actual := SnafuItoa(testcase.input)
		if actual != testcase.expected {
			t.Errorf("SnafuItoa(%#v) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestSnafuAtoi(t *testing.T) {
	cases := []struct {
		expected string
		input    int
	}{
		{"1", 1},
		{"2", 2},
		{"1=", 3},
		{"1-", 4},
		{"10", 5},
		{"11", 6},
		{"12", 7},
		{"2=", 8},
		{"2-", 9},
		{"20", 10},
		{"1=0", 15},
		{"1-0", 20},
		{"1=11-2", 2022},
		{"1-0---0", 12345},
		{"1121-1110-1=0", 314159265},
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}
	for _, testcase := range cases {
		actual := SnafuAtoi(testcase.input)
		if actual != testcase.expected {
			t.Errorf("SnafuItoa(%d) should be %s but was %s", testcase.input, testcase.expected, actual)
		}
	}
}
