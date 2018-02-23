package main

import "net/http"

func pong(wr http.ResponseWriter, req *http.Request) {
	wr.Write([]byte("pong"))
}

func createRootHandler() http.Handler {
	return http.HandlerFunc(pong)
}
