package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gossip/domain"
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

	Describe("AddMessage", func() {
		var addedMessage domain.Message

		JustBeforeEach(func() {
			addedMessage = testing.NewMessage()
			repo.AddMessage(addedMessage)
		})

		Context("Hook has been registered", func() {
			var ch chan domain.Message

			BeforeEach(func() {
				ch = make(chan domain.Message)
				repo.AddObserver(func(m domain.Message) {
					ch <- m
				})
			})

			It("Notifies observable", func() {
				Eventually(ch).Should(Receive(Equal(addedMessage)))
			})

			Context("Another hook has been registered and removed", func() {
				var ch2 chan domain.Message

				BeforeEach(func() {
					ch2 = make(chan domain.Message)
					handle := repo.AddObserver(func(m domain.Message) {
						ch2 <- m
					})
					repo.RemoveObserver(handle)
				})

				It("Notifies original observable", func() {
					Eventually(ch).Should(Receive(Equal(addedMessage)))
				})

				It("Does not notify new observable", func() {
					Consistently(ch2).ShouldNot(Receive())
				})
			})
		})
	})
})
