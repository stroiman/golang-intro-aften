package dataaccess_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDataaccess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dataaccess Suite")
}
