package aoc_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aoc Suite")
}
