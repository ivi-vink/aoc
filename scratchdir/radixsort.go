package main

import (
	"fmt"
	"math/rand"
)

const (
	signedMin = -(1 << 31)
	signedMax = (1 << 31) - 1
)

func countsort(in []int32, r uint32) {
	count := make([]int32, r)

	for i := range in {
		count[in[i]]++
	}

	for i := 1; i < len(in)-1; i++ {
		count[i] += count[i-1]
	}

	var pointer int32
	for i := 0; i < len(count)-1; i++ {
		for count[i] > 0 && pointer < count[i] {
			in[pointer] = int32(i)
			pointer++
		}
	}
}

func radix() {
}

func input(r *rand.Rand, n int) []int32 {
	out := make([]int32, n)
	for i := 0; i < n; i++ {
		out[i] = int32(r.Intn(signedMax))
	}
	return out
}

func main() {
	in := input(rand.New(rand.NewSource(1)), 1<<8)
	countsort(in, signedMax)
	fmt.Println(in)
}
