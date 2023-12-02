package day1_test

import (
	"testing"

	"github.com/johningve/adventofcode2023/day1"
	"github.com/johningve/adventofcode2023/testutil"
)

var inputPart1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	testutil.RunTest(t, day1.Part1, inputPart1, "142")
}

func BenchmarkPart1(b *testing.B) {
	testutil.RunBench(b, day1.Part1, inputPart1)
}

var inputPart2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPart2(t *testing.T) {
	testutil.RunTest(t, day1.Part2, inputPart2, "281")
}

func BenchmarkPart2(b *testing.B) {
	testutil.RunBench(b, day1.Part2, inputPart2)
}
