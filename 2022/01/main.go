package main

import (
	"context"
	"strconv"

	"mvinkio.online/aoc/aoc"
)

func readCaloriesInts(line string) (int, error) {
	if len(line) == 0 {
		return -1, nil
	}
	if i, err := strconv.Atoi(line); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

func PartOne(ctx context.Context, data []int) ([]int, error) {
	return biggestElf(data), nil
}

func PartTwo(ctx context.Context, data []int) ([]int, error) {
	return []int{sum(biggestElf(data))}, nil
}

// Boilerplate

func main() {
	aoc.RunDay(
		context.TODO(),
		aoc.NewScanCloser("2022/01/input.txt"),
		aoc.ReadByLine(readCaloriesInts),
		PartOne,
		PartTwo,
	)
}
