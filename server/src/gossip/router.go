package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

func getMessages(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write([]byte("[{},{}]"))
}

func createRootHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", pong).Methods("get")
	router.HandleFunc("/messages", getMessages).Methods("get")
	return router
}
