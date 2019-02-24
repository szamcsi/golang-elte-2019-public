// Package lines deals with lines in files.
package lines

import (
	"bufio"
	"math"
	"os"
)

func load(path string) ([]string, error) {
	// TODO: move the loading of strings here from Count().
	return nil, nil
}

// MinMaxCount represents line statistics information.
type MinMaxCount struct {
	Min, Max, Count int
}

func minMaxCount(lines []string) *MinMaxCount {
	// TODO: calculate the minimum and maximum line length while counting the lines.
	return &MinMaxCount{}
}

// Count counts the lines in a file and determines the minimal and maximal line length.
// It returns an error, if the file was not readable.
func Count(path string) (*MinMaxCount, error) {
	lines, err := load(path)
	if err != nil {
		return nil, err
	}
	return minMaxCount(lines), nil
}
