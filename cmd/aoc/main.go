package main

import (
	"fmt"

	"github.com/johningve/adventofcode2023"
	_ "github.com/johningve/adventofcode2023/day1"
)

func main() {
	solutions := adventofcode2023.Solutions()
	results := make([]chan string, 0, len(solutions))

	for _, solution := range solutions {
		ch := make(chan string, 1)
		results = append(results, ch)

		go func(s func() string) {
			ch <- s()
		}(solution)
	}

	for _, ch := range results {
		fmt.Println(<-ch)
	}
}
