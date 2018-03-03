package repository

import (
	"gossip/domain"
)

var idChan <-chan int
var foo = make(chan<- int)

func init() {
	idChan = createIdRange()
}

func createIdRange() <-chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			i++
			ch <- i
		}
	}()
	return ch
}

type MessageObserver func(domain.Message)
type ObserverHandle struct {
	id int
}

type MessageRepository struct {
	messages []domain.Message
	observer map[int]MessageObserver
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		observer: make(map[int]MessageObserver),
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

func (r *MessageRepository) AddObserver(o func(domain.Message)) int {
	token := <-idChan
	r.observer[token] = MessageObserver(o)
	return token
}

func (r *MessageRepository) RemoveObserver(token int) {
	delete(r.observer, token)
}
