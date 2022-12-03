package main

import (
	"testing"
)

func TestToPriority(t *testing.T) {
	cases := []struct {
		input    byte
		expected int
	}{
		{'a', 1},
		{'b', 2},
		{'c', 3},
		{'d', 4},
		{'x', 24},
		{'y', 25},
		{'z', 26},
		{'A', 27},
		{'B', 28},
		{'C', 29},
		{'X', 50},
		{'Y', 51},
		{'Z', 52},
	}
	for _, testcase := range cases {
		actual := toPriority(testcase.input)
		if actual != testcase.expected {
			t.Errorf("toPriority(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestFindWrongItem(t *testing.T) {
	cases := []struct {
		input    []byte
		expected int
	}{
		{[]byte("vJrwpWtwJgWrhcsFMMfFFhFp"), toPriority('p')},
		{[]byte("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), toPriority('L')},
		{[]byte("PmmdzqPrVvPwwTWBwg"), toPriority('P')},
		{[]byte("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"), toPriority('v')},
		{[]byte("ttgJtRGJQctTZtZT"), toPriority('t')},
		{[]byte("CrZsJsPPZsGzwwsLwLmpwMDw"), toPriority('s')},
	}
	for _, testcase := range cases {
		actual := findWrongItem(testcase.input)
		if actual != testcase.expected {
			t.Errorf("findWrongItem(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestAllWrongItems(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{`vJrwpWtwJgWrhcsFMMfFFhFp
		jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
		PmmdzqPrVvPwwTWBwg
		wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
		ttgJtRGJQctTZtZT
		CrZsJsPPZsGzwwsLwLmpwMDw`, 157},
	}
	for _, testcase := range cases {
		actual := findAllWrongItems(preprocess(testcase.input))
		if actual != testcase.expected {
			t.Errorf("findAllWrongItems(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestFindBadge(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{`vJrwpWtwJgWrhcsFMMfFFhFp
		jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
		PmmdzqPrVvPwwTWBwg`, toPriority('r')},
		{`wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
		ttgJtRGJQctTZtZT
		CrZsJsPPZsGzwwsLwLmpwMDw`, toPriority('Z')},
	}
	for _, testcase := range cases {
		actual := findBadge(preprocess(testcase.input), 3)
		if actual != testcase.expected {
			t.Errorf("findBadge(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}

func TestFindAllBadges(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{`vJrwpWtwJgWrhcsFMMfFFhFp
		jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
		PmmdzqPrVvPwwTWBwg
		wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
		ttgJtRGJQctTZtZT
		CrZsJsPPZsGzwwsLwLmpwMDw`, 70},
	}
	for _, testcase := range cases {
		actual := findAllBadges(preprocess(testcase.input), 3)
		if actual != testcase.expected {
			t.Errorf("findAllBadges(%q) should be %d but was %d", testcase.input, testcase.expected, actual)
		}
	}
}
