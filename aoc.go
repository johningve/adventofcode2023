package adventofcode2023

import "fmt"

var solutions []func() string

func addSolution(day, part int, run func() string) {
	solutions = append(solutions, func() string {
		return fmt.Sprintf("Day%d part%d = %s", day, part, run())
	})
}

func AddSolutionPart1(day int, run func() string) {
	addSolution(day, 1, run)
}

func AddSolutionPart2(day int, run func() string) {
	addSolution(day, 2, run)
}

func Solutions() []func() string {
	return solutions
}
