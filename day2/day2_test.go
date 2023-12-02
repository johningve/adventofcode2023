package day2_test

import (
	"testing"

	"github.com/johningve/adventofcode2023/day2"
	"github.com/johningve/adventofcode2023/testutil"
)

var input = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestPart1(t *testing.T) {
	testutil.RunTest(t, day2.Part1, input, "8")
}

func BenchmarkPart1(b *testing.B) {
	testutil.RunBench(b, day2.Part1, input)
}

func TestPart2(t *testing.T) {
	testutil.RunTest(t, day2.Part2, input, "2286")
}

func BenchmarkPart2(b *testing.B) {
	testutil.RunBench(b, day2.Part2, input)
}
