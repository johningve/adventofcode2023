package main

import (
	"strings"
	"testing"
)

var inputPart1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	var want uint64 = 142
	got := Part1(strings.NewReader(inputPart1))
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

var inputPart2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPart2(t *testing.T) {
	var want uint64 = 281
	got := Part2(strings.NewReader(inputPart2))
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}
