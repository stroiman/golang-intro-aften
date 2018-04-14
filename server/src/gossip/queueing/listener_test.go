package queueing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"

	"gossip/domain"
	. "gossip/queueing"
	"gossip/testing"
)

var _ = Describe("Listener", func() {
	var (
		c        Connection
		conn     *amqp.Connection
		ch       *amqp.Channel
		messages <-chan domain.Message
	)

	BeforeSuite(func() {
		var err error
		conn, err = CreateConnection()
		Expect(err).ToNot(HaveOccurred())
		c, err = NewConnection()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterSuite(func() {
		Expect(conn.Close()).To(Succeed())
	})

	BeforeEach(func() {
		var err error
		ch, err = conn.Channel()
		Expect(err).ToNot(HaveOccurred())
		Expect(DeclareExchange(ch)).To(Succeed())
		messages, err = c.Subscribe()
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(ch.Close()).To(Succeed())
	})

	It("Starts with a failing test", func() {
		msg := testing.NewMessage()
		Expect(c.PublishMessage(msg)).To(Succeed())
		Eventually(messages).Should(Receive(Equal(msg)))
	})
})
