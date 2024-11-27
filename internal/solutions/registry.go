package solutions

import (
	"fmt"
	"github.com/kumsuzM/aoc/internal/utils"
)

type Registry struct {
	solutions map[int]map[int]Solver
}

func NewRegistry() *Registry {
	return &Registry{
		solutions: make(map[int]map[int]Solver),
	}
}

// Register adds a solution for a specific year and day
func (r *Registry) Register(year, day int, solver Solver) {
	if r.solutions[year] == nil {
		r.solutions[year] = make(map[int]Solver)
	}
	r.solutions[year][day] = solver
}

// Get returns the solver for a specific year and day
func (r *Registry) Get(year, day int) (Solver, error) {
	yearSolutions, yearExists := r.solutions[year]
	if !yearExists {
		return nil, fmt.Errorf("no solutions registered for year %d", year)
	}

	solver, exists := yearSolutions[day]
	if !exists {
		return nil, fmt.Errorf("no solution registered for year %d day %d", year, day)
	}

	return solver, nil
}

// Solve handles solution execution for any registered year/day combination
func (r *Registry) Solve(year, day, part int) (string, error) {
	solver, err := r.Get(year, day)
	if err != nil {
		return "", err
	}

	input, err := utils.ReadInput(year, day)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	switch part {
	case 1:
		return solver.Part1(input)
	case 2:
		return solver.Part2(input)
	default:
		return "", fmt.Errorf("invalid part %d", part)
	}
}
