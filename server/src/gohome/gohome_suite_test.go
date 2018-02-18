package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGohome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gohome Suite")
}
