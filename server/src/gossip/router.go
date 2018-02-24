package main

import (
	"encoding/json"
	"net/http"

	. "gossip/domain"
	"gossip/repository"

	"github.com/gorilla/mux"
)

type MessageRepository interface {
	GetMessages() []Message
	AddMessage(Message)
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

func createRootHandler() http.Handler {
	repo := repository.NewMessageRepository()
	messageHandler := NewMessageHandler(repo)
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.Handle("/api/messages", messageHandler)
	return router
}
