package main

import (
	"net/http"
)

func main() {
	handler := createRootHandler()
	http.ListenAndServe("0.0.0.0:9000", handler)
}
