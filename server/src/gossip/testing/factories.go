package testing

import (
	"github.com/google/uuid"
	"gossip/domain"
)

func NewMessage() domain.Message {
	return domain.Message{
		Id:      uuid.New().String(),
		Message: "Test message",
	}
}
