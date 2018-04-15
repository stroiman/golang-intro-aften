package application_test

import (
	. "gossip/application"
	. "gossip/application/mock_application"
	"gossip/domain"
	. "gossip/testing/matchers"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MessageHub", func() {
	var (
		ctrl         *gomock.Controller
		exchangeMock *MockMessageExchange
		publishedMsg domain.Message
		hub          MessageHub
		msgChan      chan domain.Message
	)

	BeforeEach(func() {
		publishedMsg = domain.Message{}
		ctrl = gomock.NewController(GinkgoT())
		exchangeMock = NewMockMessageExchange(ctrl)
		msgChan = make(chan domain.Message)
		exchangeMock.EXPECT().Subscribe().Return(msgChan, nil)
	})

	AfterEach(func() {
		close(msgChan)
	})

	JustBeforeEach(func() {
		hub.Listen(exchangeMock)
		msgChan <- publishedMsg
		// hub.Notify(publishedMsg)
	})

	Context("Hook has been registered", func() {
		var ch chan domain.Message

		BeforeEach(func() {
			ch = make(chan domain.Message)
			hub.AddObserver(func(m domain.Message) {
				ch <- m
			})
		})

		It("Notifies observable", func() {
			Eventually(ch).Should(Receive(HaveMessage(Equal(publishedMsg.Message))))
		})

		Context("Another hook has been registered and removed", func() {
			var ch2 chan domain.Message

			BeforeEach(func() {
				ch2 = make(chan domain.Message)
				handle := hub.AddObserver(func(m domain.Message) {
					ch2 <- m
				})
				hub.RemoveObserver(handle)
			})

			It("Notifies original observable", func() {
				Eventually(ch).Should(Receive(HaveMessage(Equal(publishedMsg.Message))))
			})

			It("Does not notify new observable", func() {
				Consistently(ch2).ShouldNot(Receive())
			})
		})
	})
})
