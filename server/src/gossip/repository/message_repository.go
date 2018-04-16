package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gossip/domain"
	"time"
)

var idChan <-chan ObserverHandle

type ObserverHandle struct {
	handle int
}

func NewObserverHandle(i int) ObserverHandle {
	return ObserverHandle{i}
}

func init() {
	idChan = createIdRange()
}

func createIdRange() <-chan ObserverHandle {
	ch := make(chan ObserverHandle)
	go func() {
		i := 0
		for {
			i++
			ch <- ObserverHandle{i}
		}
	}()
	return ch
}

type MessageObserver func(domain.Message)
type observerMap map[ObserverHandle]MessageObserver

type MessageRepository struct {
	messages []domain.Message
	observer observerMap
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		messages: []domain.Message{},
		observer: make(observerMap),
	}
}

func (r *MessageRepository) AddMessage(message domain.Message) (domain.Message, error) {
	message.Id = uuid.New().String()
	message.CreatedAt = time.Now()
	r.messages = append(r.messages, message)
	for _, o := range r.observer {
		go o(message)
	}
	return message, nil
}

func (r *MessageRepository) UpdateMessage(m domain.Message) error {
	for i := range r.messages {
		if r.messages[i].Id == m.Id {
			r.messages[i].Message = m.Message
			now := time.Now()
			r.messages[i].EditedAt = &now
			for _, o := range r.observer {
				go o(r.messages[i])
			}
			return nil
		}
	}
	return fmt.Errorf("No message found with id: %s", m.Id)
}

func (r *MessageRepository) GetMessages() ([]domain.Message, error) {
	return r.messages, nil
}

func (r *MessageRepository) AddObserver(o func(domain.Message)) ObserverHandle {
	handle := <-idChan
	r.observer[handle] = MessageObserver(o)
	return handle
}

func (r *MessageRepository) RemoveObserver(handle ObserverHandle) {
	delete(r.observer, handle)
}
