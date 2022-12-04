package main

import (
	"testing"
)

func TestSectionPairFromString(t *testing.T) {
	cases := []struct {
		input    string
		expected SectionPair
	}{
		{"2-4,6-8", SectionPair{Section{2, 4}, Section{6, 8}}},
		{"2-3,4-5", SectionPair{Section{2, 3}, Section{4, 5}}},
		{"5-7,7-9", SectionPair{Section{5, 7}, Section{7, 9}}},
		{"2-8,3-7", SectionPair{Section{2, 8}, Section{3, 7}}},
		{"6-6,4-6", SectionPair{Section{6, 6}, Section{4, 6}}},
		{"2-6,4-8", SectionPair{Section{2, 6}, Section{4, 8}}},
	}
	for _, testcase := range cases {
		actual := SectionPairFromString(testcase.input)
		if actual != testcase.expected {
			t.Errorf("SectionFromString(%#v) should be %#v but was %#v", testcase.input, testcase.expected, actual)
		}
	}
}

func TestSectionPairFullyContainEachOther(t *testing.T) {
	cases := []struct {
		input    SectionPair
		expected bool
	}{
		{SectionPair{Section{2, 4}, Section{6, 8}}, false},
		{SectionPair{Section{2, 3}, Section{4, 5}}, false},
		{SectionPair{Section{5, 7}, Section{7, 9}}, false},
		{SectionPair{Section{2, 8}, Section{3, 7}}, true},
		{SectionPair{Section{6, 6}, Section{4, 6}}, true},
		{SectionPair{Section{2, 6}, Section{4, 8}}, false},
	}
	for _, testcase := range cases {
		actual := testcase.input.FullyContainsEachOther()
		if actual != testcase.expected {
			t.Errorf("SectionPair{%#v}.FullyContainEachOther() should be %#v but was %#v", testcase.input, testcase.expected, actual)
		}
	}
}

func TestSectionPairOverlap(t *testing.T) {
	cases := []struct {
		input    SectionPair
		expected bool
	}{
		{SectionPair{Section{2, 4}, Section{6, 8}}, false},
		{SectionPair{Section{2, 3}, Section{4, 5}}, false},
		{SectionPair{Section{5, 7}, Section{7, 9}}, true},
		{SectionPair{Section{2, 8}, Section{3, 7}}, true},
		{SectionPair{Section{6, 6}, Section{4, 6}}, true},
		{SectionPair{Section{2, 6}, Section{4, 8}}, true},
	}
	for _, testcase := range cases {
		actual := testcase.input.Overlaps()
		if actual != testcase.expected {
			t.Errorf("SectionPair{%#v}.Overlaps() should be %#v but was %#v", testcase.input, testcase.expected, actual)
		}
	}
}
