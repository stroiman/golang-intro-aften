package repository

import (
	. "gossip/domain"
)

type MessageRepository struct {
	messages []Message
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		messages: []Message{{
			Id:      "1",
			Message: "Foo",
		}, {
			Id:      "2",
			Message: "Bar",
		}},
	}
}

func (r *MessageRepository) GetMessages() []Message {
	return r.messages
}
