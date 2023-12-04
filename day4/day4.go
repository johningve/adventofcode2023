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
				score++
			}
		}
		sum += 1 << (score - 1)
	}

	return strconv.FormatUint(sum, 10)
}

func Part2(input string) string {
	cardNumber := 0 // I don't care what the input says.
	var copiesPerCard []int
	addCopies := func(card, amount int) {
		if len(copiesPerCard)-1 < card {
			copiesPerCard = append(copiesPerCard, amount)
		} else {
			copiesPerCard[card] += amount
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		addCopies(cardNumber, 1)
		line := scanner.Text()
		i, j := indexOrPanic(line, ':'), indexOrPanic(line, '|')
		numbers := readNumbers(line[i+1 : j-1])
		winning := readNumbers(line[j+1:])
		score := 0
		copies := copiesPerCard[cardNumber]
		for _, n := range numbers {
			if slices.Contains(winning, n) {
				score++
				addCopies(cardNumber+score, copies)
			}
		}
		cardNumber++
	}

	var sum uint64
	for _, copies := range copiesPerCard {
		sum += uint64(copies)
	}

	return strconv.FormatUint(sum, 10)
}
