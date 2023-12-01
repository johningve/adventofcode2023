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
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func toDigit(c byte) uint64 {
	return uint64(c - '0')
}

func Part1(input io.Reader) uint64 {
	var sum uint64 = 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var (
			firstDigit uint64
			lastDigit  uint64
		)

		line := scanner.Text()
		for i := range line {
			if isDigit(line[i]) {
				if firstDigit == 0 {
					firstDigit = toDigit(line[i])
				}
				lastDigit = toDigit(line[i])
			}
		}

		sum += 10*firstDigit + lastDigit
	}

	return sum
}
