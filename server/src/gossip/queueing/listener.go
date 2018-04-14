// Package queueing provides support for sending and retrieving messages from
// RabbitMQ/AMQP
package queueing

import (
	"encoding/json"
	"errors"
	"fmt"
	"gossip/domain"

	"github.com/streadway/amqp"
)

type Connection struct {
	conn *amqp.Connection
	ch   *amqp.Channel // Channed for publishing messages
}

func CreateConnection() (ch *amqp.Connection, err error) {
	return amqp.Dial("amqp://guest:guest@localhost")
}

func CreateChannel() (ch *amqp.Channel, err error) {
	var conn *amqp.Connection
	if conn, err = CreateConnection(); err == nil {
		ch, err = conn.Channel()
	}
	return
}

func DeclareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare("gossip-messages",
		"fanout" /* exchange kind */, true, /* durable */
		false /* autodelete */, false, /* internal */
		false /* nowait */, nil /* args */)
}

func NewConnection() (result Connection, err error) {
	result.conn, err = amqp.Dial("amqp://localhost/")
	if err == nil {
		result.ch, err = result.conn.Channel()
	}
	if err == nil {
		DeclareExchange(result.ch)
	}
	return
}

func (c Connection) Close() error {
	if c.conn == nil {
		return errors.New("Connection not initialized")
	}
	return c.conn.Close()
}

func (c Connection) PublishMessage(msg domain.Message) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(msg); err != nil {
		return
	}
	pub := amqp.Publishing{Body: bytes}
	fmt.Println("**** ch", c.ch)
	return c.ch.Publish(
		"gossip-messages", /* exchange to publish to */
		"",                /* routing key, N/A for fanout exch */
		false,             /* mandatory */
		false,             /* immediate */
		pub)
}

func (c Connection) subscribeDeliveries() (res <-chan amqp.Delivery, err error) {
	var ch *amqp.Channel
	ch, err = c.conn.Channel()
	if err != nil {
		return
	}
	var queue amqp.Queue
	if queue, err = ch.QueueDeclare("", /* name, rabbit creates a unique name when empty */
		false /* durable */, false, /* autoDelete */
		false /* exclusive */, false, /* noWait */
		nil /* args */); err != nil {
		return
	}
	if err = ch.QueueBind(queue.Name, "", "gossip-messages", false, nil); err != nil {
		return
	}
	return ch.Consume(queue.Name, "", false /* autoAck */, false, /* exclusive */
		false /* noLocal */, false /* noWait */, nil)
}

func (c Connection) Subscribe() (res <-chan domain.Message, err error) {
	var (
		deliveries <-chan amqp.Delivery
		output     = make(chan domain.Message)
	)
	if deliveries, err = c.subscribeDeliveries(); err == nil {
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
	}
	res = output
	return
}
