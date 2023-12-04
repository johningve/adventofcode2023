package main

import (
	"fmt"
	"time"

	"github.com/johningve/adventofcode2023"
	_ "github.com/johningve/adventofcode2023/day1"
	_ "github.com/johningve/adventofcode2023/day2"
	_ "github.com/johningve/adventofcode2023/day3"
	_ "github.com/johningve/adventofcode2023/day4"
)

func main() {
	solutions := adventofcode2023.Solutions()
	results := make([]chan string, 0, len(solutions))

	start := time.Now()

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

	fmt.Printf("Ran %d solutions in %v\n", len(solutions), time.Since(start))
}
