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
}

//go:embed input.txt
var input string

type token uint8

func Part1(input string) string {
	type cubes struct {
		red   uint
		green uint
		blue  uint
	}

	maxCubes := cubes{12, 13, 14}

	findGameID := func(input string) (gameID uint, remainder string) {
		i := strings.IndexByte(input, ':')
		if i == -1 {
			panic(fmt.Sprintf("Invalid input: missing ':' in '%s'", input))
		}
		if n, err := fmt.Sscanf(input[:i], "Game %d", &gameID); n != 1 || err != nil {
			panic(fmt.Sprintf("Failed to read game ID (error is %v)", err))
		}
		return gameID, input[i+1:]
	}

	findHandful := func(input string) (handful string, remainder string) {
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

	parseCubes := func(input string) cubes {
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

	validCubes := func(cubes, maxCubes cubes) bool {
		return cubes.red <= maxCubes.red && cubes.green <= maxCubes.green && cubes.blue <= maxCubes.blue
	}

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

			if !validCubes(cubes, maxCubes) {
				continue outer
			}
		}
		sum += int(gameID)
	}

	return strconv.FormatUint(uint64(sum), 10)
}
