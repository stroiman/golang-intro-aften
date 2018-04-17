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
	GetMessages() ([]Message, error)
	AddMessage(Message) (Message, error)
	UpdateMessage(Message) error
}

type MessageObservable interface {
	AddObserver(func(Message)) repository.ObserverHandle
}

type SocketPublisher struct {
	Observable MessageObservable `inject:""`
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
	startListener(conn, p.Observable)
}

func pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

type MessageHandler struct {
	Repository MessageRepository
}

func NewMessageHandler(repo MessageRepository) http.Handler {
	handler := &MessageHandler{
		Repository: repo,
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/messages", handler.GetMessages).Methods("GET")
	router.HandleFunc("/api/messages", handler.PostMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", handler.PutMessage).Methods("PUT")
	return router
}

func (h *MessageHandler) GetMessages(wr http.ResponseWriter, req *http.Request) {
	if messages, err := h.Repository.GetMessages(); err == nil {
		if response, err := json.Marshal(messages); err == nil {
			wr.Header().Set("Content-Type", "application/json")
			wr.Header().Set("Access-Control-Allow-Origin", "*")
			wr.WriteHeader(http.StatusOK)
			wr.Write(response)
			return
		}
	}
	wr.WriteHeader(500)
}

func (h *MessageHandler) PutMessage(wr http.ResponseWriter, req *http.Request) {
	var message Message
	vars := mux.Vars(req)
	id := vars["id"]
	if err := json.NewDecoder(req.Body).Decode(&message); err == nil {
		if message.IsValidInput() {
			message.Id = id
			fmt.Println("Update message", message)
			h.Repository.UpdateMessage(message)
			wr.Header().Set("Content-Type", "application/json")
			wr.Write([]byte(`{ "status": "ok" }`))
			return
		}
	}
	wr.WriteHeader(500)
}

func (h *MessageHandler) PostMessage(wr http.ResponseWriter, req *http.Request) {
	var message Message
	fmt.Println("POST HANDLER")
	err := json.NewDecoder(req.Body).Decode(&message)
	if err == nil {
		if !message.IsValidInput() {
			wr.WriteHeader(400)
			return
		}
		if _, err = h.Repository.AddMessage(message); err == nil {
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
	Repository MessageRepository `inject:""`
	// MessageHandler *MessageHandler `inject:""`
	SocketPublisher *SocketPublisher `inject:""`
	http.Handler
}

func (h *HttpHandler) Init() error {
	// socketPublisher := NewSocketPublisher(h.Observable)
	messageHandler := NewMessageHandler(h.Repository)
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.PathPrefix("/api/messages").Handler(messageHandler)
	router.Handle("/socket", h.SocketPublisher)
	h.Handler = router
	return nil
}
