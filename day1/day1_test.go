package main

import (
	"strings"
	"testing"
)

var testInput = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

func TestPart1(t *testing.T) {
	var want uint64 = 142
	got := Part1(strings.NewReader(testInput))

	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}
