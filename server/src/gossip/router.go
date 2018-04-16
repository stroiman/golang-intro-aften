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
	AddMessage(Message) (Message, error)
	UpdateMessage(Message) error
}

type MessageObservable interface {
	AddObserver(func(Message)) repository.ObserverHandle
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

func pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

type MessageHandler struct {
	repository MessageRepository
}

func NewMessageHandler(repo MessageRepository) http.Handler {
	handler := &MessageHandler{
		repository: repo,
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PutMessage).Methods("PUT")
	return router
}

func (h *MessageHandler) GetMessages(wr http.ResponseWriter, req *http.Request) {
	if response, err := json.Marshal(h.repository.GetMessages()); err == nil {
		wr.Header().Set("Content-Type", "application/json")
		wr.Header().Set("Access-Control-Allow-Origin", "*")
		wr.WriteHeader(http.StatusOK)
		wr.Write(response)
	} else {
		wr.WriteHeader(500)
	}
}

func (h *MessageHandler) PutMessage(wr http.ResponseWriter, req *http.Request) {
	var message Message
	vars := mux.Vars(req)
	id := vars["id"]
	if err := json.NewDecoder(req.Body).Decode(&message); err == nil {
		if message.IsValidInput() {
			message.Id = id
			fmt.Println("Update message", message)
			h.repository.UpdateMessage(message)
			wr.Header().Set("Content-Type", "application/json")
			wr.Write([]byte(`{ "status": "ok" }`))
			return
		}
	}
	wr.WriteHeader(500)
}

func (h *MessageHandler) PostMessage(wr http.ResponseWriter, req *http.Request) {
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

type HttpHandler struct {
	Repository MessageRepository
	http.Handler
}

func (h *HttpHandler) Init() error {
	repo := repository.NewMessageRepository()
	socketPublisher := NewSocketPublisher(repo)
	messageHandler := NewMessageHandler(repo)
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.PathPrefix("/api/messages").Handler(messageHandler)
	router.Handle("/socket", socketPublisher)
	h.Handler = router
	return nil
}

func createRootHandler() http.Handler {
	repo := repository.NewMessageRepository()
	socketPublisher := NewSocketPublisher(repo)
	messageHandler := NewMessageHandler(repo)
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.PathPrefix("/api/messages").Handler(messageHandler)
	router.Handle("/socket", socketPublisher)
	return router
}
