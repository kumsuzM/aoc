package day02

import (
	"math"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) Part1(input string) (string, error) {
	reports := parseInput(input)
	numSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafe++
		}
	}

	return strconv.Itoa(numSafe), nil
}

func isSafe(report []int) bool {
	var isAsc bool = false

	for i := range report {
		// skip past first item
		if i == 0 {
			continue
		}
		diff := math.Abs(float64(report[i]) - float64(report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		currIsAsc := report[i] > report[i-1]
		if i != 1 && isAsc != currIsAsc {
			return false
		}
		isAsc = currIsAsc
	}

	return true
}

func (s *Solver) Part2(input string) (string, error) {
	return "day 2 part 1 sol", nil
}

func parseInput(input string) [][]int {
	reports := make([][]int, 0, 1000)

	// go through each line, create a report and add it to the list
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		strReport := strings.Fields(line)
		intReport := make([]int, len(strReport))
		for i, strLevel := range strReport {
			intLevel, err := strconv.Atoi(strLevel)
			if err != nil {
				panic(err)
			}
			intReport[i] = intLevel
		}
		reports = append(reports, intReport)
	}

	return reports
}
