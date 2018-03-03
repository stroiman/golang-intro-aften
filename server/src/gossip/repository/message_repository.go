package repository

import (
	"gossip/domain"
)

var idChan <-chan ObserverHandle

type ObserverHandle struct {
	handle int
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
		observer: make(observerMap),
	}
}

func (r *MessageRepository) AddMessage(message domain.Message) {
	r.messages = append(r.messages, message)
	for _, o := range r.observer {
		go o(message)
	}
}

func (r *MessageRepository) GetMessages() []domain.Message {
	return r.messages
}

func (r *MessageRepository) AddObserver(o func(domain.Message)) ObserverHandle {
	handle := <-idChan
	r.observer[handle] = MessageObserver(o)
	return handle
}

func (r *MessageRepository) RemoveObserver(handle ObserverHandle) {
	delete(r.observer, handle)
}
