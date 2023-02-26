package calories

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Elfs", func() {
})

var data = make([]int, 1000000)

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum(data)
	}
}
