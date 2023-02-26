package calories_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCalories(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calories Suite")
}
