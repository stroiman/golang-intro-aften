package queueing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueueing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queueing Suite")
}
