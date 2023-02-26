package main

import (
	"context"

	"mvinkio.online/aoc/2022/01/calories"
	"mvinkio.online/aoc/aoc"
)

// Boilerplate

func main() {
	aoc.RunDay(
		context.TODO(),
		aoc.NewScanCloser("2022/01/input.txt"),
		aoc.ReadByLine(calories.Reader),
		calories.Solvers,
	)
}
