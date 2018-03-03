package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gossip/repository"
	"gossip/testing"
)

var _ = Describe("MessageRepository", func() {
	var (
		repo *MessageRepository
	)

	BeforeEach(func() {
		repo = NewMessageRepository()
	})

	Describe("GetMessages", func() {
		Context("With two messages added", func() {
			BeforeEach(func() {
				repo.AddMessage(testing.NewMessage())
				repo.AddMessage(testing.NewMessage())
			})

			It("Returns two messages", func() {
				actual := repo.GetMessages()
				Expect(actual).To(HaveLen(2))
			})
		})
	})
})
