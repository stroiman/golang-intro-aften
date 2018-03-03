package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "gossip/domain"
	"gossip/repository"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type MessageRepository interface {
	GetMessages() []Message
	AddMessage(Message)
}

type MessageObservable interface {
	AddObserver(func(Message))
}

type SocketPublisher struct {
	observable MessageObservable
}

func NewSocketPublisher(o MessageObservable) *SocketPublisher {
	result := &SocketPublisher{o}
	return result
}

func startListener(conn *websocket.Conn, o MessageObservable) {
	o.AddObserver(func(m Message) {
		conn.WriteJSON(m)
	})
}

func (p *SocketPublisher) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(wr, req, nil)
	fmt.Println("Connection attempt", err)
	if err != nil {
		fmt.Println(err)
		return
	}
	startListener(conn, p.observable)
}

func GetMessages() []Message {
	return []Message{{
		Id:      "1",
		Message: "Foo",
	}, {
		Id:      "2",
		Message: "Bar",
	}}
}

func pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

type MessageHandler struct {
	repository MessageRepository
}

func NewMessageHandler(repo MessageRepository) *MessageHandler {
	return &MessageHandler{
		repository: repo,
	}
}

func (h *MessageHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var message Message
		err := json.NewDecoder(req.Body).Decode(&message)
		if err == nil {
			if message.IsValidInput() {
				h.repository.AddMessage(message)
				wr.Header().Set("Content-Type", "application/json")
				wr.Write([]byte(`{ "status": "ok" }`))
				return
			}
		}
		wr.WriteHeader(500)
	}
	if response, err := json.Marshal(h.repository.GetMessages()); err == nil {
		wr.Header().Set("Content-Type", "application/json")
		wr.Header().Set("Access-Control-Allow-Origin", "*")
		wr.WriteHeader(http.StatusOK)
		wr.Write(response)
	} else {
		wr.WriteHeader(500)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true },
}

func handleSocket(wr http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(wr, req, nil)
	fmt.Println("Connection attempt", err)
	if err != nil {
		fmt.Println(err)
		return
	}
	message := Message{
		Id:      "42",
		Message: "From websocket",
	}
	err = conn.WriteJSON(message)
	fmt.Println("Write result", err)
}

func createRootHandler() http.Handler {
	repo := repository.NewMessageRepository()
	socketPublisher := NewSocketPublisher(repo)
	messageHandler := NewMessageHandler(repo)
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.Handle("/api/messages", messageHandler)
	router.Handle("/socket", socketPublisher)
	return router
}
