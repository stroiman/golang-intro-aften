package application

import "gossip/domain"

var idChan <-chan ObserverHandle

func init() {
	idChan = createIdRange()
}

type MessageExchange interface {
	Listen() <-chan domain.Message
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

// MessageHub handles listening for incoming messages and publishing it to multiple
// consumers. Each publish is carried out in a separate gorouting in order to not
// block the overall processing of messages because of one bad consumer.
type MessageHub struct {
	observer observerMap
}

type MessageObserver func(domain.Message)
type observerMap map[ObserverHandle]MessageObserver

type ObserverHandle struct {
	handle int
}

func (hub *MessageHub) AddObserver(o func(domain.Message)) ObserverHandle {
	hub.ensureObserver()
	handle := <-idChan
	hub.observer[handle] = MessageObserver(o)
	return handle
}

func (hub *MessageHub) RemoveObserver(handle ObserverHandle) {
	delete(hub.observer, handle)
}

func (hub *MessageHub) ensureObserver() {
	if hub.observer == nil {
		hub.observer = make(observerMap)
	}
}

func (hub *MessageHub) Notify(msg domain.Message) {
	for _, o := range hub.observer {
		go o(msg)
	}
}
