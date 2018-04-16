package repository_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gossip/domain"
	. "gossip/repository"
	"gossip/testing"
	. "gossip/testing/matchers"
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
				actual, _ := repo.GetMessages()
				Expect(actual).To(HaveLen(2))
			})
		})
	})

	Describe("Update", func() {
		Context("A message has already been added", func() {
			var msg domain.Message

			BeforeEach(func() {
				var err error
				msg = testing.NewMessage()
				msg.Message = "Old message"
				msg, err = repo.AddMessage(msg)
				Expect(err).ToNot(HaveOccurred())
			})

			It("Updates the message", func() {
				msg.Message = "New message"
				err := repo.UpdateMessage(msg)
				Expect(err).ToNot(HaveOccurred())
				messages, _ := repo.GetMessages()
				Expect(messages).To(HaveLen(1))
				Expect(messages[0]).To(HaveMessage(Equal("New message")))
			})
		})

		It("Returns an error if id is invalid", func() {
			msg := testing.NewMessage()
			err := repo.UpdateMessage(msg)
			Expect(err).To(HaveOccurred())
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
				Eventually(ch).Should(Receive(HaveMessage(Equal(addedMessage.Message))))
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
					Eventually(ch).Should(Receive(HaveMessage(Equal(addedMessage.Message))))
				})

				It("Does not notify new observable", func() {
					Consistently(ch2).ShouldNot(Receive())
				})
			})
		})
	})
})
