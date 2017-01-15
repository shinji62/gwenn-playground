package founding_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFounding(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Founding Suite")
}
