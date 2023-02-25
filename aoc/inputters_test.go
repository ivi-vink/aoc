package aoc_test

import (
	"bufio"
	"context"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"mvinkio.online/aoc/aoc"
)

func newScanner(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

var _ = Describe("Inputters", func() {
	Describe("Bylines Inputter", func() {
		in := aoc.ReadByLine[int](func(line string) (int, error) {
			if i, err := strconv.Atoi(line); err != nil {
				return 0, err
			} else {
				return i, nil
			}
		})
		It("should map each line in the input to the expected type", func() {
			data, err := in.Read(context.TODO(), newScanner("1\n1\n8\n3\n5\n8"))
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]int{1, 1, 8, 3, 5, 8}))
		})
	})

	Describe("Lines Inputter", func() {
		in := aoc.ReadLines[[]int](func(lines []string) ([]int, error) {
			ints := make([]int, len(lines))
			for i, l := range lines {
				t, err := strconv.Atoi(l)
				if err != nil {
					return nil, err
				}
				ints[i] = t
			}
			return ints, nil
		})
		It("Should correct map to type given all lines", func() {
			data, err := in.Read(context.TODO(), newScanner("1\n1\n8\n3\n5\n8"))
			Expect(err).To(BeNil())
			Expect(data).To(Equal([]int{1, 1, 8, 3, 5, 8}))
		})
	})
})
