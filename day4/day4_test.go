package day4_test

import (
	"testing"

	"github.com/johningve/adventofcode2023/day4"
	"github.com/johningve/adventofcode2023/testutil"
)

var input = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestPart1(t *testing.T) {
	testutil.RunTest(t, day4.Part1, input, "13")
}

func BenchmarkPart1(b *testing.B) {
	testutil.RunBench(b, day4.Part1, input)
}

func TestPart2(t *testing.T) {
	testutil.RunTest(t, day4.Part2, input, "30")
}

func BenchmarkPart2(b *testing.B) {
	testutil.RunBench(b, day4.Part2, input)
}
