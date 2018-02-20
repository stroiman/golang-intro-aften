package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gohome/api"

	"github.com/gorilla/mux"
)

type Blog struct {
	Id string `json:"id"`
}

func handler(wr http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(wr)
	println("Request")
	wr.WriteHeader(200)
	blog := Blog{
		Id: "42",
	}
	encoder.Encode([]Blog{blog, blog})
}

func Pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

func CreateRootHandler() http.Handler {
	router := mux.NewRouter()
	dir := http.Dir("static")
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", api.Router))
	router.HandleFunc("/ping", http.HandlerFunc(Pong))
	router.PathPrefix("/static/").Handler(http.FileServer(dir))
	router.PathPrefix("/").Handler(http.FileServer(dir))
	return router
}

func main() {
	println("Starting server on port 9000")
	handler := CreateRootHandler()
	http.ListenAndServe("0.0.0.0:9000", handler)
}
