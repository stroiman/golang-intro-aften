package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Id      string `json:"id"`
	Message string `json:"message"`
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

func getMessages(wr http.ResponseWriter, req *http.Request) {
	if response, err := json.Marshal(GetMessages()); err == nil {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusOK)
		wr.Write(response)
	} else {
		wr.WriteHeader(500)
	}
}

func createRootHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.HandleFunc("/api/messages", getMessages).Methods("get")
	return router
}
