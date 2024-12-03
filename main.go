package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kumsuzM/aoc/internal/solutions"
	"github.com/kumsuzM/aoc/internal/solutions/2024/day01"
	"github.com/kumsuzM/aoc/internal/solutions/2024/day02"
	"github.com/kumsuzM/aoc/internal/solutions/2024/day03"
)

var registry = solutions.NewRegistry()

func init() {
	// Register solutions for each year/day
	registry.Register(2024, 1, day01.New())
	registry.Register(2024, 2, day02.New())
	registry.Register(2024, 3, day03.New())
	// register more days here...
}

func main() {
	year := flag.Int("year", 2024, "year of puzzle to run")
	day := flag.Int("day", 1, "day of puzzle to run")
	part := flag.Int("part", 1, "part of puzzle to run (1 or 2)")

	flag.Parse()

	solution, err := registry.Solve(*year, *day, *part)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Year %d Day %d Part %d: %s\n", *year, *day, *part, solution)
}
