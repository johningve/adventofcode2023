package day1_test

import (
	"testing"

	"github.com/johningve/adventofcode2023/day1"
)

var inputPart1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	want := "142"
	got := day1.Part1(inputPart1)
	if got != want {
		t.Errorf("Part1() = %s, want %s", got, want)
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
	want := "281"
	got := day1.Part2(inputPart2)
	if got != want {
		t.Errorf("Part2() = %s, want %s", got, want)
	}
}
