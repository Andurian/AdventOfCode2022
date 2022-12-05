package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
	"testing"
)

var testCargoString = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 `

var testInstructions = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestTopCrates(t *testing.T) {
	cargo := CargoFromString(testCargoString)
	actual := cargo.CratesOnTop()
	expected := "NDP"
	if actual != expected {
		t.Errorf("Cargo{%#v}.CratesOnTop should be %s but was %s", cargo, expected, actual)
	}
}

func TestApplyInstructions(t *testing.T) {
	cargo := CargoFromString(testCargoString)
	instructions := util.Transform(strings.Split(testInstructions, "\n"), InstructionFromString)

	actual := topCratesAfterInstuctions(cargo, instructions)
	expected := "CMZ"

	if actual != expected {
		t.Errorf("ApplyInstructions(...) should be %s but was %s", expected, actual)
	}

}

func TestApplyBatchInstructions(t *testing.T) {
	cargo := CargoFromString(testCargoString)
	instructions := util.Transform(strings.Split(testInstructions, "\n"), InstructionFromString)

	actual := topCratesAfterBatchInstuctions(cargo, instructions)
	expected := "MCD"

	if actual != expected {
		t.Errorf("ApplyInstructions(...) should be %s but was %s", expected, actual)
	}

}
