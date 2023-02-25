package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Returns two slices of the input array representing both backpacks
func ruckSacks(inputline []byte) ([]byte, []byte) {
	split := len(inputline) / 2
	return inputline[:split], inputline[split:]
}

// Gets the priority based on the character byte
func itemPriority(item byte) byte {
	if 'a' <= item && item <= 'z' {
		return (item - 'a')
	}
	if 'A' <= item && item <= 'Z' {
		return (item - 'A' + 26)
	}
	return 0
}

// Makes an array of 52 flags indicating with 0 or 1 if an item is in a rucksack
func insideFlags(ruckSack []byte) []byte {
	flags := make([]byte, 52)
	for _, item := range ruckSack {
		if flags[item] != 1 {
			flags[item] = 1
		}
	}
	return flags
}

// Checks if one of the items in a rucksacks is flagged in another rucksack
func inBoth(ruckSack []byte, insideFlags []byte) byte {
	for _, item := range ruckSack {
		if insideFlags[item] == 1 {
			return item
		}
	}
	return 0
}

// Finds the item that is in all rucksacks
func inAll(elfGroup []byte, groupSize int) int {
	for i, itemCount := range elfGroup {
		if int(itemCount) == groupSize {
			return i + 1
		}
	}
	return 0
}

// Solves the rucksack problem in linear time (since we don't have nested loops to check for the presence in the rucksacks),
// could be more efficient with memory i think? But that would involve more bookkeeping
func main() {
	f, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal("Could not open inputfile")
	}
	// Part 1
	runningSum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		priorities := make([]byte, len(scanner.Bytes()))
		for i, item := range scanner.Bytes() {
			priorities[i] = itemPriority(item)
		}
		r1, r2 := ruckSacks(priorities)
		insideR2 := insideFlags(r2)
		duplicate := inBoth(r1, insideR2)
		runningSum += int(duplicate) + 1
	}
	fmt.Println("part 1:", runningSum)

	// Part 2
	runningSum = 0
	elfGroup := make([]byte, 52)
	groupSize := 3
	elf := 0
	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		priorities := make([]byte, len(scanner.Bytes()))
		for i, item := range scanner.Bytes() {
			priorities[i] = itemPriority(item)
		}
		for i, flag := range insideFlags(priorities) {
			elfGroup[i] += flag
		}
		if elf == (groupSize - 1) {
			runningSum += inAll(elfGroup, groupSize)
			elfGroup = make([]byte, 52)
		}
		elf = (elf + 1) % groupSize
	}
	fmt.Println("part 2:", runningSum)
}
