package wunderground_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestWunderground(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wunderground Suite")
}
