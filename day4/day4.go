package day4

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/johningve/adventofcode2023"
)

//go:embed input.txt
var input string

func init() {
	adventofcode2023.AddSolutionPart1(4, func() string { return Part1(input) })
	adventofcode2023.AddSolutionPart2(4, func() string { return Part2(input) })
}

func indexOrPanic(str string, chr byte) int {
	i := strings.IndexByte(str, chr)
	if i == -1 {
		panic(fmt.Sprintf("expected to find '%v' in '%s'", chr, str))
	}
	return i
}

var whitespaceReg = regexp.MustCompile("\\s+")

func readNumbers(str string) (numbers []int) {
	ss := whitespaceReg.Split(strings.TrimSpace(str), -1)
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

func Part1(input string) string {
	var sum uint64

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		i, j := indexOrPanic(line, ':'), indexOrPanic(line, '|')
		numbers := readNumbers(line[i+1 : j-1])
		winning := readNumbers(line[j+1:])
		var score uint64 = 0
		for _, n := range numbers {
			if slices.Contains(winning, n) {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
		sum += score
	}

	return strconv.FormatUint(sum, 10)
}

func Part2(input string) string {
	return ""
}
