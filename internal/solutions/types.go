package solutions

type Solver interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}
