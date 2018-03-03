package repository

import (
	"gossip/domain"
)

type MessageObserver func(domain.Message)

type MessageRepository struct {
	messages []domain.Message
	observer MessageObserver
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (r *MessageRepository) AddMessage(message domain.Message) {
	r.messages = append(r.messages, message)
	if r.observer != nil {
		go r.observer(message)
	}
}

func (r *MessageRepository) GetMessages() []domain.Message {
	return r.messages
}

func (r *MessageRepository) AddObserver(o func(domain.Message)) {
	r.observer = MessageObserver(o)
}
