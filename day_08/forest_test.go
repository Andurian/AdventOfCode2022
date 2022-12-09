package main

import (
	"andurian/adventofcode/2022/util"
	. "andurian/adventofcode/2022/util/point"
	"reflect"
	"testing"
)

func TestForestFromString(t *testing.T) {
	input := util.ReadSafe("input_test.txt")
	expected := map[Point]int{
		{Row: 0, Col: 0}: 3, {Row: 0, Col: 1}: 0, {Row: 0, Col: 2}: 3, {Row: 0, Col: 3}: 7, {Row: 0, Col: 4}: 3,
		{Row: 1, Col: 0}: 2, {Row: 1, Col: 1}: 5, {Row: 1, Col: 2}: 5, {Row: 1, Col: 3}: 1, {Row: 1, Col: 4}: 2,
		{Row: 2, Col: 0}: 6, {Row: 2, Col: 1}: 5, {Row: 2, Col: 2}: 3, {Row: 2, Col: 3}: 3, {Row: 2, Col: 4}: 2,
		{Row: 3, Col: 0}: 3, {Row: 3, Col: 1}: 3, {Row: 3, Col: 2}: 5, {Row: 3, Col: 3}: 4, {Row: 3, Col: 4}: 9,
		{Row: 4, Col: 0}: 3, {Row: 4, Col: 1}: 5, {Row: 4, Col: 2}: 3, {Row: 4, Col: 3}: 9, {Row: 4, Col: 4}: 0}
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
		{Point{Row: 1, Col: 2}, 4},
		{Point{Row: 3, Col: 2}, 8},
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
