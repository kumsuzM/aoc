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
	var isAsc bool = report[1] > report[0]

	for i := range report {
		// skip past first item
		if i == 0 {
			continue
		}
		diff := math.Abs(float64(report[i]) - float64(report[i-1]))
		currIsAsc := report[i] > report[i-1]
		if (diff < 1 || diff > 3) || (i > 1 && isAsc != currIsAsc) {
			return false
		}
	}

	return true
}

func (s *Solver) Part2(input string) (string, error) {
	reports := parseInput(input)
	numSafe := 0
	for _, report := range reports {
		if isSafe2(report) {
			numSafe++
		}
	}

	return strconv.Itoa(numSafe), nil
}

func isSafe2(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := range report {
		dampened := make([]int, 0, len(report)-1)
		dampened = append(dampened, report[:i]...)
		dampened = append(dampened, report[i+1:]...)

		// Check if this dampened version is safe
		if isSafe(dampened) {
			return true
		}
	}

	return false
}

func dampenable(currIsAsc bool, currIdx, before, after int) (bool, bool) {
	// 1. first check if removing me will fix diff
	diff := math.Abs(float64(before) - float64(after))
	if diff < 1 || diff > 3 {
		return false, false
	}

	// 2. second check if removing me will change isAsc
	if currIdx == 1 {
		return true, after > before
	} else {
		return currIsAsc == (after > before), currIsAsc
	}
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
