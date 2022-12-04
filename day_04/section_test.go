package main

import (
	"testing"
)

func TestSectionsFullyContainEachOther(t *testing.T) {
	cases := []struct {
		input1   Section
		input2   Section
		expected bool
	}{
		{Section{2, 4}, Section{6, 8}, false},
		{Section{2, 3}, Section{4, 5}, false},
		{Section{5, 7}, Section{7, 9}, false},
		{Section{2, 8}, Section{3, 7}, true},
		{Section{6, 6}, Section{4, 6}, true},
		{Section{2, 6}, Section{4, 8}, false},
	}
	for _, testcase := range cases {
		actual := testcase.input1.Contains(testcase.input2) || testcase.input2.Contains(testcase.input1)
		if actual != testcase.expected {
			t.Errorf("SectionsFullyContainEachOther(%#v, %#v) should be %#v but was %#v", testcase.input1, testcase.input2, testcase.expected, actual)
		}
	}
}

func TestSectionsOverlap(t *testing.T) {
	cases := []struct {
		input1   Section
		input2   Section
		expected bool
	}{
		{Section{2, 4}, Section{6, 8}, false},
		{Section{2, 3}, Section{4, 5}, false},
		{Section{5, 7}, Section{7, 9}, true},
		{Section{2, 8}, Section{3, 7}, true},
		{Section{6, 6}, Section{4, 6}, true},
		{Section{2, 6}, Section{4, 8}, true},
	}
	for _, testcase := range cases {
		actual1 := testcase.input1.Overlaps(testcase.input2)
		actual2 := testcase.input2.Overlaps(testcase.input1)
		if actual1 != testcase.expected {
			t.Errorf("Overlaps(%#v, %#v) should be %#v but was %#v", testcase.input1, testcase.input2, testcase.expected, actual1)
		}
		if actual2 != testcase.expected {
			t.Errorf("Overlaps(%#v, %#v) should be %#v but was %#v", testcase.input2, testcase.input1, testcase.expected, actual1)
		}
	}
}
