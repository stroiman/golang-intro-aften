package main

import (
  "net/http"
)

func pong(wr http.ResponseWriter, req *http.Request) {
  wr.Write([]byte("pong"));
}

func createHttpHandler() http.Handler {
  return http.HandlerFunc(pong);
}

func main() {
  handler := createHttpHandler();
  http.ListenAndServe("0.0.0.0:9000", handler);
}
