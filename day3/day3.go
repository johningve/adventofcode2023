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

func overlaps(a, b []int) bool {
	return a[0] < b[1] && a[1] > b[0]
}

func Part1(input string) string {
	var (
		line, prevLine, nextLine string
		sum                      uint64
		numberReg                = regexp.MustCompile("\\d+")
		symbolReg                = regexp.MustCompile("[^.\\d\\s]")
	)

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
				if overlaps([]int{n[0] - 1, n[1] + 1}, s) {
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
	var (
		line, prevLine, nextLine string
		sum                      uint64
		numberReg                = regexp.MustCompile("\\d+")
		asteriskReg              = regexp.MustCompile("\\*")
	)

	scanner := bufio.NewScanner(strings.NewReader(input))

	if !scanner.Scan() {
		panic("expected at least 2 lines of input")
	}
	line = scanner.Text()
	prevLine = line // HACK: we only count numbers in line, and this won't introduce any new adjacencies.

	search := func() {
		asterisksFound := asteriskReg.FindAllStringIndex(line, -1)

		lines := []string{prevLine, line, nextLine}
		for _, a := range asterisksFound {
			var (
				c int
				p uint64 = 1
			)
			for _, l := range lines {
				numbersFound := numberReg.FindAllStringIndex(l, -1)
				for _, n := range numbersFound {
					if overlaps([]int{n[0] - 1, n[1] + 1}, a) {
						number, err := strconv.ParseUint(l[n[0]:n[1]], 10, 64)
						if err != nil {
							panic(err)
						}
						c++
						p *= number
					}
				}
			}
			if c == 2 {
				sum += p
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
