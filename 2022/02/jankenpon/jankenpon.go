package jankenpon

import (
	"context"
)

var (
	Reader  = readJankenpon
	Solvers = []func(ctx context.Context, data any) (any, error){
		partOne, partTwo,
	}
)

func partOne(ctx context.Context, data any) (any, error) {
	return nil, nil
}

func partTwo(ctx context.Context, data any) (any, error) {
	return nil, nil
}

func readJankenpon(line []string) (any, error) {
	return nil, nil
}
