package dataaccess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gossip/dataaccess"
)

var url = "postgres://gossip:gossip@127.0.0.1/gossip?sslmode=disable"

var _ = Describe("Dataaccess", func() {
	It("Initializes", func() {
		_, err := NewConnection(url)
		Expect(err).ToNot(HaveOccurred())
	})
})
