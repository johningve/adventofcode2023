package day3_test

import (
	"testing"

	"github.com/johningve/adventofcode2023/day3"
	"github.com/johningve/adventofcode2023/testutil"
)

var input = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	testutil.RunTest(t, day3.Part1, input, "4361")
}

func BenchmarkPart1(b *testing.B) {
	testutil.RunBench(b, day3.Part1, input)
}

func TestPart2(t *testing.T) {
	testutil.RunTest(t, day3.Part2, input, "467835")
}

func BenchmarkPart2(b *testing.B) {
	testutil.RunBench(b, day3.Part2, input)
}
