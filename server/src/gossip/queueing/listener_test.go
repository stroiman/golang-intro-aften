package queueing_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/streadway/amqp"

	"gossip/domain"
	. "gossip/queueing"
	"gossip/testing"
)

func DeclareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare("gossip-messages",
		"fanout" /* exchange kind */, true, /* durable */
		false /* autodelete */, false, /* internal */
		false /* nowait */, nil /* args */)
}

func subscribe(ch *amqp.Channel) (res <-chan domain.Message, err error) {
	var queue amqp.Queue
	if queue, err = ch.QueueDeclare("", /* name, rabbit creates a unique name when empty */
		false /* durable */, false, /* autoDelete */
		true /* exclusive */, false, /* noWait */
		nil /* args */); err != nil {
		return
	}
	if err = ch.QueueBind(queue.Name, "", "gossip-messages", false, nil); err != nil {
		return
	}
	var deliveries <-chan amqp.Delivery
	if deliveries, err = ch.Consume(queue.Name, "", false /* autoAck */, false, /* exclusive */
		false /* noLocal */, false /* noWait */, nil); err != nil {
		return
	}
	output := make(chan domain.Message)
	res = output
	go func() {
		for d := range deliveries {
			msg := domain.Message{}
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				fmt.Printf("Error receiving message, cannot parse json")
			} else {
				output <- msg
			}
		}
	}()
	return
}

func publishMessage(ch *amqp.Channel, msg domain.Message) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(msg); err != nil {
		return
	}
	pub := amqp.Publishing{Body: bytes}
	return ch.Publish(
		"gossip-messages", /* exchange to publish to */
		"",                /* routing key, N/A for fanout exch */
		false,             /* mandatory */
		false,             /* immediate */
		pub)
}

var _ = Describe("Listener", func() {
	var (
		conn     *amqp.Connection
		ch       *amqp.Channel
		messages <-chan domain.Message
	)

	BeforeSuite(func() {
		var err error
		conn, err = CreateConnection()
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		var err error
		ch, err = conn.Channel()
		Expect(err).ToNot(HaveOccurred())
		Expect(DeclareExchange(ch)).To(Succeed())
		messages, err = subscribe(ch)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		Expect(ch.Close()).To(Succeed())
	})

	It("Starts with a failing test", func() {
		msg := testing.NewMessage()
		Expect(publishMessage(ch, msg)).To(Succeed())
		Eventually(messages).Should(Receive(Equal(msg)))
	})
})
