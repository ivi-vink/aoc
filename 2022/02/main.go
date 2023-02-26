package main

import (
	"context"

	"mvinkio.online/aoc/2022/02/jankenpon"
	"mvinkio.online/aoc/aoc"
)

// Boilerplate

func main() {
	aoc.RunDay(
		context.TODO(),
		aoc.NewScanCloser("2022/02/input.txt"),
		aoc.ReadLines(jankenpon.Reader),
		jankenpon.Solvers,
	)
}
