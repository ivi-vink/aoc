package aoc

import (
	"bufio"
	"log"
	"os"
)

func NewScannerFromFile(filename string) *bufio.Scanner {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	return bufio.NewScanner(f)
}
