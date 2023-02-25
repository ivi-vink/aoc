package main

import (
	"bufio"
	"fmt"
	"strconv"

	"mvinkio.online/aoc/aoc"
)

func biggestElf(data []int) []int {
	topThree := make([]int, 3)
	push := func(i int, a int) {
		tmp := topThree[i]
		topThree[i] = a
		carry := tmp
		for i++; i < len(topThree); i++ {
			tmp = topThree[i]
			topThree[i] = carry
			carry = tmp
		}
	}
	Elf := 0
	for i := 0; i < len(data); i++ {
		if data[i] < 0 {
			for j := 0; j < len(topThree); j++ {
				if topThree[j] < Elf {
					push(j, Elf)
					break
				}
			}
			Elf = 0
			continue
		}
		Elf = Elf + data[i]
	}
	return topThree
}

func sum(array []int) int {
	s := 0
	for _, val := range array {
		s += val
	}
	return s
}

// Boilerplate
func main() {
	s := aoc.NewScannerFromFile("2022/01/input.txt")

	data := make([]int, 0)
	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			data = append(data, -1)
		} else if num, err := strconv.Atoi(line); err == nil {
			data = append(data, num)
		}
	}
	fmt.Println(biggestElf(data))
	fmt.Println(sum(biggestElf(data)))
}
