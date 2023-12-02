package day2

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/johningve/adventofcode2023"
)

func init() {
	adventofcode2023.AddSolutionPart1(2, func() string { return Part1(input) })
	adventofcode2023.AddSolutionPart2(2, func() string { return Part2(input) })
}

//go:embed input.txt
var input string

type cubes struct {
	red   uint
	green uint
	blue  uint
}

func (c cubes) lessThan(other cubes) bool {
	return c.red <= other.red && c.green <= other.green && c.blue <= other.blue
}

func (c cubes) max(other cubes) cubes {
	return cubes{
		red:   max(c.red, other.red),
		green: max(c.green, other.green),
		blue:  max(c.blue, other.blue),
	}
}

func (c cubes) pow() uint64 {
	return uint64(c.red) * uint64(c.green) * uint64(c.blue)
}

func findGameID(input string) (gameID uint, remainder string) {
	i := strings.IndexByte(input, ':')
	if i == -1 {
		panic(fmt.Sprintf("Invalid input: missing ':' in '%s'", input))
	}
	if n, err := fmt.Sscanf(input[:i], "Game %d", &gameID); n != 1 || err != nil {
		panic(fmt.Sprintf("Failed to read game ID (error is %v)", err))
	}
	return gameID, input[i+1:]
}

func findHandful(input string) (handful string, remainder string) {
	i := strings.IndexByte(input, ';')
	if i == -1 {
		return input, ""
	}
	return input[:i], input[i+1:]
}

var (
	redRegex   = regexp.MustCompile("(\\d+) red")
	greenRegex = regexp.MustCompile("(\\d+) green")
	blueRegex  = regexp.MustCompile("(\\d+) blue")
)

func parseCubes(input string) cubes {
	var cubes cubes
	read := func(input string, re *regexp.Regexp, cubes *uint) {
		groups := re.FindStringSubmatch(input)
		if groups != nil {
			n, err := strconv.ParseUint(groups[1], 10, 32)
			if err != nil {
				panic("Failed to read number of cubes")
			}
			*cubes = uint(n)
		}
	}
	read(input, redRegex, &cubes.red)
	read(input, greenRegex, &cubes.green)
	read(input, blueRegex, &cubes.blue)
	return cubes
}

func Part1(input string) string {
	maxCubes := cubes{12, 13, 14}

	sum := 0
	scanner := bufio.NewScanner(strings.NewReader(input))

outer:
	for scanner.Scan() {
		line := scanner.Text()

		var gameID uint
		gameID, line = findGameID(line)

		for len(line) > 0 {
			var (
				zeroCubes cubes
				handful   string
			)

			handful, line = findHandful(line)
			cubes := parseCubes(handful)

			if cubes == zeroCubes {
				break
			}

			if !cubes.lessThan(maxCubes) {
				continue outer
			}
		}
		sum += int(gameID)
	}

	return strconv.FormatUint(uint64(sum), 10)
}

func Part2(input string) string {
	var sum uint64 = 0
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		_, line = findGameID(line)

		var minCubes cubes
		for len(line) > 0 {
			var handful string

			handful, line = findHandful(line)
			cubes := parseCubes(handful)
			minCubes = cubes.max(minCubes)
		}
		sum += minCubes.pow()
	}

	return strconv.FormatUint(sum, 10)
}
