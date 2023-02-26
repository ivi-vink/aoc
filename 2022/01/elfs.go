package main

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
