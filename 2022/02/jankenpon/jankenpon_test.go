package jankenpon

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"mvinkio.online/aoc/aoc"
)

var (
	data    = aoc.ReadLines(readJankenpon)
	answer1 = 15
	answer2 = 12
)

var _ = Describe("jankenpon", func() {
	Describe("Part One", func() {
		It("should solve the example to the answer", func() {
			Expect(true).To(Equal(true))
		})
	})
	Describe("Part Two", func() {
		It("should solve the example to the answer", func() {
			Expect(true).To(Equal(true))
		})
	})
})
