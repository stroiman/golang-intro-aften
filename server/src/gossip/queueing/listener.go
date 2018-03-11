// Package queueing provides support for sending and retrieving messages from
// RabbitMQ/AMQP
package queueing

import (
	"github.com/streadway/amqp"
)

func CreateConnection() (ch *amqp.Connection, err error) {
	return amqp.Dial("amqp://localhost/")
}

func CreateChannel() (ch *amqp.Channel, err error) {
	var conn *amqp.Connection
	if conn, err = CreateConnection(); err == nil {
		ch, err = conn.Channel()
	}
	return
}
