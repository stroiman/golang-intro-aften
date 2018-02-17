package main

import "net/http"

func handler(wr http.ResponseWriter, req *http.Request) {
	println("Request")
	w.WriteHeader(200)
	w.Write([]byte("hello"))
}

func main() {
	println("hello")
	http.ListenAndServe("0.0.0.0:9000", http.HandlerFunc(handler))
}
