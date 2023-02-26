package aoc

import (
	"bufio"
	"context"
	"strconv"
)

type (
	Reader[T any] func(ctx context.Context, scanner *bufio.Scanner) (T, error)
	Lines[T any]  func(lines []string) (T, error)
	ByLine[T any] func(line string) (T, error)
)

// Read all lines into type T
func ReadLines[T any](l Lines[T]) Reader[T] {
	return func(ctx context.Context, scanner *bufio.Scanner) (T, error) {
		lines := []string{}
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return l(lines)
	}
}

// Read line by line into type T, collect in []T
func ReadByLine[T any](l ByLine[T]) Reader[[]T] {
	return func(ctx context.Context, scanner *bufio.Scanner) ([]T, error) {
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
}

// Read lines into ints
var ReadInts = ReadByLine(func(line string) (int, error) {
	if i, err := strconv.Atoi(line); err != nil {
		return 0, err
	} else {
		return i, nil
	}
})
