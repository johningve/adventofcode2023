package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(Part1(strings.NewReader(input)))
	fmt.Println(Part2(strings.NewReader(input)))
}

func solve(input io.Reader, getDigit func(str string, i int) (uint64, bool)) uint64 {
	var sum uint64 = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			firstDigit uint64
			lastDigit  uint64
		)

		line := scanner.Text()
		for i := range line {
			digit, ok := getDigit(line, i)
			if !ok {
				continue
			}
			if firstDigit == 0 {
				firstDigit = digit
			}
			lastDigit = digit
		}

		sum += 10*firstDigit + lastDigit
	}

	return sum
}

func getDigit(str string, i int) (uint64, bool) {
	if str[i] >= '1' && str[i] <= '9' {
		return uint64(str[i] - '0'), true
	}
	return 0, false
}

func Part1(input io.Reader) uint64 {
	return solve(input, getDigit)
}

var digits = map[string]uint64{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getSpelledDigit(str string, i int) (uint64, bool) {

	for k, v := range digits {
		if len(str)-i < len(k) {
			continue
		}
		if str[i:i+len(k)] == k {
			return v, true
		}
	}
	return 0, false
}

func Part2(input io.Reader) uint64 {

	return solve(input, func(str string, i int) (uint64, bool) {
		d, ok := getDigit(str, i)
		if ok {
			return d, ok
		}
		return getSpelledDigit(str, i)
	})
}
