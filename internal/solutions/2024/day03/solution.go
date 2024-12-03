package day03

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
	sum := 0
	i := 0
	for i < len(input) {
		res, indexInc := tryParse(input, i)
		sum += res
		i += indexInc
	}

	return strconv.Itoa(sum), nil
}

func tryParse(input string, startIdx int) (int, int) {
	//mul(X,Y)
	mIdx := startIdx
	uIdx := startIdx + 1
	lIdx := startIdx + 2
	openBrackIdx := startIdx + 3
	num1Start := startIdx + 4
	if input[mIdx] != 'm' {
		return 0, 1
	}

	if input[uIdx] != 'u' {
		return 0, 2
	}

	if input[lIdx] != 'l' {
		return 0, 3
	}

	if input[openBrackIdx] != '(' {
		return 0, 3
	}

	parsedNum1, numDigits1 := parseNum(input, num1Start)
	if input[num1Start+numDigits1] != ',' {
		return 0, 3 + numDigits1
	}
	num2start := num1Start + numDigits + 1
	parsedNum1, numDigits2 := parseNum(input, num2start)
	if input[num1Start+numDigits2] != ')' {
		return 0, 3 + numDigits1 + 1 + numDigits2
	}

}

func parseNum(input string, startIdx int) (int, int) {
	var builder strings.Builder

	i := startIdx
	for input[startIdx] >= '0' && input[startIdx] <= '9' {
		builder.WriteString(string(input[startIdx]))
		i += 1
	}
	result := builder.String()
	resultNum, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}

	return resultNum, len(result)
}

func (s *Solver) Part2(input string) (string, error) {
	return "sol 2", nil
}
