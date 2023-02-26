package aoc

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

type ScanCloser struct {
	*bufio.Scanner
	io.Closer
}

func NewScanCloser(file string) *ScanCloser {
	f, err := os.Open(file)
	if err != nil {
		log.Panic(err)
	}
	return &ScanCloser{
		Scanner: bufio.NewScanner(f),
		Closer:  f,
	}
}

type Solvers[T any, R any] []func(ctx context.Context, data T) (R, error)

func RunDay[T any, R any](ctx context.Context, s *ScanCloser, in Reader[T], solvers Solvers[T, R]) error {
	defer s.Closer.Close()

	data, err := in(ctx, s.Scanner)
	if err != nil {
		return err
	}

	results := make([]R, 0)
	for _, solver := range solvers {
		out, err := solver(ctx, data)
		if err != nil {
			return fmt.Errorf("")
		}
		results = append(results, out)
	}

	for _, result := range results {
		fmt.Println(result)
	}
	return nil
}
