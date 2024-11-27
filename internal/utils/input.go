package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadInput(year, day int) (string, error) {
	path := filepath.Join("inputs", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d.txt", day))

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read input file: %w", err)
	}

	return string(data), nil
}
