package aoc

import (
	"bufio"
	"context"
)

type Reader[T any] interface {
	Read(ctx context.Context, scanner *bufio.Scanner) (T, error)
}

type Solver[T any, R any] interface {
	Solve(ctx context.Context, data T) (R, error)
}

func RunDay[T any, R any](ctx context.Context, in Reader[T], solvers ...Solver[T, R]) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
}
