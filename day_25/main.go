package main

import (
	"andurian/adventofcode/2022/util"
	"strings"
)

var snafuDigits = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

var snafuDigitsRev = map[int]rune{
	2:  '2',
	1:  '1',
	0:  '0',
	-1: '-',
	-2: '=',
}

func SnafuItoa(s string) int {
	ret := 0
	for i, exp := 0, 1; i < len(s); i, exp = i+1, exp*5 {
		c := rune(s[len(s)-1-i])
		ret += exp * snafuDigits[c]
	}
	return ret
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	ret := base
	for i := 1; i < exp; i += 1 {
		ret *= base
	}
	return ret
}

func SnafuAtoi(x int) string {
	digitsAndFirstDigit := func(x int) (int, rune) {
		for i, exp := 0, 1; ; i, exp = i+1, exp*5 {
			if x < exp-exp/2 {
				if x < 2*(exp/5)-(exp/5)/2 {
					return i, '1'
				}
				return i, '2'
			}
		}
	}

	nextDigit := func(x, exp int) (int, rune) {
		a := x
		for ex := exp; ex > 0; ex /= 5 {
			a += 2 * ex
		}
		b := a / exp
		c := b - 2
		ret := snafuDigitsRev[c]
		return x - c*exp, ret
	}

	digits, firstDigit := digitsAndFirstDigit(x)
	var b strings.Builder
	b.WriteRune(firstDigit)
	exp := pow(5, digits-1)
	x -= snafuDigits[firstDigit] * exp
	for exp = exp / 5; exp > 0; exp /= 5 {
		y, d := nextDigit(x, exp)
		b.WriteRune(d)
		x = y
	}
	return b.String()
}

func Task1(s string) string {
	ret := 0
	for _, line := range strings.Split(s, "\n") {
		ret += SnafuItoa(line)
	}
	return SnafuAtoi(ret)
}

func main() {
	input := util.ReadSafe("input.txt")
	util.ExecuteTimedString(25, 1, func() string { return Task1(input) })
}
