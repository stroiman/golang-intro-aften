package application_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gossip/application"
)

var _ = Describe("Application", func() {
	Describe("GetMessages", func() {
		It("Returns the messages from the data access layer", func() {
			app := NewApplication()
			_, err := app.GetMessages()
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
