package main

import "fmt"

var sortingdata []int = []int{1, 23, 4, 234, 123, 4, 4512, 0}

func selectionSort(in []int) []int {
	end := len(in)
	for i := 0; i < end; i++ {
		min := i
		for j := i; j < end; j++ {
			if in[j] < in[min] {
				min = j
			}
		}
		in[i], in[min] = in[min], in[i]
	}
	return in
}

func main() {
	fmt.Println(selectionSort(sortingdata))
}
