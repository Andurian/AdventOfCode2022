package main

import (
	"andurian/adventofcode/2022/util"
	"reflect"
	"testing"
)

func TestForestFromString(t *testing.T) {
	input := util.ReadSafe("input_test.txt")
	expected := map[Point]int{
		{0, 0}: 3, {0, 1}: 0, {0, 2}: 3, {0, 3}: 7, {0, 4}: 3,
		{1, 0}: 2, {1, 1}: 5, {1, 2}: 5, {1, 3}: 1, {1, 4}: 2,
		{2, 0}: 6, {2, 1}: 5, {2, 2}: 3, {2, 3}: 3, {2, 4}: 2,
		{3, 0}: 3, {3, 1}: 3, {3, 2}: 5, {3, 3}: 4, {3, 4}: 9,
		{4, 0}: 3, {4, 1}: 5, {4, 2}: 3, {4, 3}: 9, {4, 4}: 0}
	actual := ForestFromString(input).heights

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("ForestFromString(%#v) should be %#v but was %#v", input, expected, actual)
	}
}

func TestForestCountVisibleTrees(t *testing.T) {
	forest := ForestFromString(util.ReadSafe("input_test.txt"))
	expected := 21
	actual := forest.CountVisibleTrees()

	if actual != expected {
		t.Errorf("Forest{%#v}.CountVisibleTrees() was %d but should be %d", forest, expected, actual)
	}
}

func TestForestScenicScore(t *testing.T) {
	cases := []struct {
		input    Point
		expected int
	}{
		{Point{1, 2}, 4},
		{Point{3, 2}, 8},
	}

	forest := ForestFromString(util.ReadSafe("input_test.txt"))

	for _, testcase := range cases {
		actual := forest.ScenicScore(testcase.input)
		if actual != testcase.expected {
			t.Errorf("Forest{%#v}.ScenicScore(%#v) was %d but should be %d", forest, testcase.input, testcase.expected, actual)
		}
	}
}

func TestForestBestScenicScore(t *testing.T) {
	forest := ForestFromString(util.ReadSafe("input_test.txt"))
	expected := 8
	actual := forest.BestScenicScore()

	if actual != expected {
		t.Errorf("Forest{%#v}.BestScenicScore() was %d but should be %d", forest, expected, actual)
	}
}
