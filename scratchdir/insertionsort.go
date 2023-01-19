package main

import "fmt"

var sortingdata []int = []int{1, 23, 4, 234, 123, 4, 4512, 0}

func insert(in []int) {
	for i := len(in) - 1; i > 0; i-- {
		if in[i-1] >= in[i] {
			in[i], in[i-1] = in[i-1], in[i]
		}
	}
}

func insertionSort(in []int) []int {
	for i := 0; i < len(in); i++ {
		insert(in[:i+1])
	}
	return in
}

func main() {
	fmt.Println(insertionSort(sortingdata))
}
