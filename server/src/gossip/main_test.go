package main

import "testing"
import "net/http/httptest"

func TestPoing(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/ping", nil)
	handler := createRootHandler()
	handler.ServeHTTP(recorder, request)
	body := string(recorder.Body.Bytes())
	if body != "pong" {
		t.Errorf("Error, invalid response \"%s\"", body)
	}
}
