package repository_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/types"

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

	Describe("Update", func() {
		It("Updates the message", func() {
			msg := testing.NewMessage()
			msg.Message = "Old message"
			repo.AddMessage(msg)
			newMessage := repo.GetMessages()[0]
			newMessage.Message = "New message"
			err := repo.UpdateMessage(newMessage)
			Expect(err).ToNot(HaveOccurred())
			messages := repo.GetMessages()
			Expect(messages).To(HaveLen(1))
			Expect(messages[0]).To(HaveMessage(Equal("New message")))
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

type MessageMatcher struct {
	message domain.Message
	matcher GomegaMatcher
}

func HaveMessage(matcher GomegaMatcher) *MessageMatcher { return &MessageMatcher{matcher: matcher} }

func (m *MessageMatcher) Match(actual interface{}) (success bool, err error) {
	if message, ok := actual.(domain.Message); ok {
		m.message = message
		return m.matcher.Match(message.Message)
	}
	return false, fmt.Errorf("MessageMatcher expects a message")
}

func (m *MessageMatcher) FailureMessage(actual interface{}) string {
	return m.matcher.FailureMessage(m.message.Message)
}

func (m *MessageMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.matcher.NegatedFailureMessage(m.message.Message)
}
