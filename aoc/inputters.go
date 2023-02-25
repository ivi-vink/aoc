package aoc

import (
	"bufio"
	"context"
	"strconv"
)

// Read all lines into type T
type ReadLines[T any] func(lines []string) (T, error)

func (l ReadLines[T]) Read(ctx context.Context, scanner *bufio.Scanner) (T, error) {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return l(lines)
}

// Read line by line into type T, collect in []T
type ReadByLine[T any] func(line string) (T, error)

func (l ReadByLine[T]) Read(ctx context.Context, scanner *bufio.Scanner) ([]T, error) {
	lines := []T{}
	for scanner.Scan() {
		if t, err := l(scanner.Text()); err == nil {
			lines = append(lines, t)
		} else {
			return nil, err
		}
	}
	return lines, nil
}

// Read lines into ints
var ReadInts = ReadByLine[int](func(line string) (int, error) {
	if i, err := strconv.Atoi(line); err != nil {
		return 0, err
	} else {
		return i, nil
	}
})
