package calories

import (
	"context"
	"strconv"
)

type solver func(ctx context.Context, data []int) ([]int, error)

var (
	Reader  = readCaloriesInts
	Solvers = []solver{partOne, partTwo}
)

func partOne(ctx context.Context, data []int) ([]int, error) {
	return biggestElf(data), nil
}

func partTwo(ctx context.Context, data []int) ([]int, error) {
	return []int{sum(biggestElf(data))}, nil
}

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
		if data[i] > 0 {
			Elf = Elf + data[i]
		}
		if data[i] < 0 || i == len(data)-1 {
			for j := 0; j < len(topThree); j++ {
				if topThree[j] < Elf {
					push(j, Elf)
					break
				}
			}
			Elf = 0
			continue
		}
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
