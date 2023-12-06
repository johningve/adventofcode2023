package day5

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/johningve/adventofcode2023"
)

//go:embed input.txt
var input string

func init() {
	adventofcode2023.AddSolutionPart1(5, func() string { return Part1(input) })
	adventofcode2023.AddSolutionPart2(5, func() string { return Part2(input) })
}

func readNumbers(str string) (numbers []int) {
	ss := strings.Split(strings.TrimSpace(str), " ")
	numbers = make([]int, 0, len(ss))
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

const (
	destRangeStart int = iota
	sourceRangeStart
	rangeLength
)

func Part1(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	if !scanner.Scan() {
		panic("no input")
	}

	line := scanner.Text()

	var (
		prevNumbers []int
		numbers     = readNumbers(line[6:])
	)

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			for i, n := range numbers {
				if n == 0 {
					numbers[i] = prevNumbers[i] // pass thru any numbers that don't match
				}
			}
			prevNumbers, numbers = numbers, make([]int, len(numbers))
			continue
		}

		if strings.HasSuffix(line, "map:") {
			continue
		}

		mapping := readNumbers(line)
		if len(mapping) != 3 {
			panic(fmt.Sprintf("Invalid mapping '%v' read from line '%s'", mapping, line))
		}

		for i, n := range prevNumbers {
			diff := n - mapping[sourceRangeStart]
			if diff >= 0 && diff < mapping[rangeLength] {
				numbers[i] = mapping[destRangeStart] + diff
			}
		}
	}

	for i, n := range numbers {
		if n == 0 {
			numbers[i] = prevNumbers[i] // pass thru any numbers that don't match
		}
	}
	prevNumbers, numbers = numbers, make([]int, len(numbers))

	return strconv.Itoa(slices.Min(prevNumbers))
}

func Part2(input string) string {
	return ""
}
