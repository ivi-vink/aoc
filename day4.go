package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Use min and max to check if fully contained, since we have sorted sections
type section struct {
	min int
	max int
}

func sectionFromString(sectionRange string) section {
	minmax := strings.Split(sectionRange, "-")
	min, errMin := strconv.Atoi(minmax[0])
	max, errMax := strconv.Atoi(minmax[1])
	if errMin != nil || errMax != nil {
		log.Fatal("Got unexpected input", sectionRange)
	}
	return section{
		min: min,
		max: max,
	}
}

func (s section) isInside(o section) bool {
	if o.min <= s.min && s.max <= o.max {
		return true
	}
	return false
}

// Use negative cases because it is easier to think about when ranges don't overlap
func (s section) isOverlapping(o section) bool {
	if s.max < o.min {
		return false
	}
	if o.max < s.min {
		return false
	}
	return true
}

func main() {
	f, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal("Could not open input file")
	}

	// part 1
	scanner := bufio.NewScanner(f)
	runningSum := 0
	for scanner.Scan() {
		sectionStrings := strings.Split(scanner.Text(), ",")
		sectionA, sectionB := sectionFromString(sectionStrings[0]), sectionFromString(sectionStrings[1])
		if sectionA.isInside(sectionB) || sectionB.isInside(sectionA) {
			runningSum += 1
		}
	}
	fmt.Println("Part 1:", runningSum)

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)

	// Part 2
	runningSum = 0
	for scanner.Scan() {
		sectionStrings := strings.Split(scanner.Text(), ",")
		sectionA, sectionB := sectionFromString(sectionStrings[0]), sectionFromString(sectionStrings[1])
		if sectionA.isOverlapping(sectionB) {
			runningSum += 1
		}
	}
	fmt.Println("Part 2:", runningSum)
}
