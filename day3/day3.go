package day3

import (
	"bufio"
	_ "embed"
	"regexp"
	"strconv"
	"strings"

	"github.com/johningve/adventofcode2023"
)

//go:embed input.txt
var input string

func init() {
	adventofcode2023.AddSolutionPart1(3, func() string { return Part1(input) })
	adventofcode2023.AddSolutionPart2(3, func() string { return Part2(input) })
}

func Part1(input string) string {
	var (
		line, prevLine, nextLine string
		numberReg                = regexp.MustCompile("\\d+")
		symbolReg                = regexp.MustCompile("[^.\\d\\s]")
	)

	var sum uint64 = 0

	scanner := bufio.NewScanner(strings.NewReader(input))

	if !scanner.Scan() {
		panic("expected at least 2 lines of input")
	}
	line = scanner.Text()
	prevLine = line // HACK: we only count numbers in line, and this won't introduce any new adjacencies.

	search := func() {
		numbersFound := numberReg.FindAllStringIndex(line, -1)
		symbolsFound := symbolReg.FindAllStringIndex(prevLine, -1)
		symbolsFound = append(symbolsFound, symbolReg.FindAllStringIndex(line, -1)...)
		symbolsFound = append(symbolsFound, symbolReg.FindAllStringIndex(nextLine, -1)...)

		for _, n := range numbersFound {
			for _, s := range symbolsFound {
				if n[0]-1 < s[1] && n[1]+1 > s[0] {
					number, err := strconv.ParseUint(line[n[0]:n[1]], 10, 64)
					if err != nil {
						panic(err)
					}
					sum += number
					break
				}
			}
		}
	}

	for scanner.Scan() {
		nextLine = scanner.Text()

		search()

		prevLine, line = line, nextLine
	}

	nextLine = line // HACK again: we need to check the last line too. Same as before, this won't change the result

	search()

	return strconv.FormatUint(sum, 10)
}

func Part2(input string) string {
	return ""
}
