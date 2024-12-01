package day01

import (
	// "fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

// Implements solutions.Solver interface
func (s *Solver) Part1(input string) (string, error) {
	l, r := parseInput(input)

	sort.Ints(l)
	sort.Ints(r)

	sum := 0
	for i := 0; i < 1000; i++ {
		diff := math.Abs(float64(l[i]) - float64(r[i]))
		sum += int(diff)
	}

	return strconv.Itoa(sum), nil
}

func (s *Solver) Part2(input string) (string, error) {
	// Your solution for part 2
	return "solution to part 2", nil
}

func parseInput(input string) ([]int, []int) {
	l := make([]int, 0, 1000)
	r := make([]int, 0, 1000)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		leftNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		rightNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		l = append(l, leftNum)
		r = append(r, rightNum)
	}

	return l, r

}
