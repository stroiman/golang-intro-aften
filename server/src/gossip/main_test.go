package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"testing"
)

func TestPoing(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/ping", nil)
	handler := HttpHandler{}
	handler.Init()
	handler.ServeHTTP(recorder, request)
	body := string(recorder.Body.Bytes())
	if body != "pong" {
		t.Errorf("Error, invalid response \"%s\"", body)
	}
}

var _ = Describe("Bootstrapper", func() {
	It("Succeeds", func() {
		_, err := CreateRootObj()
		Expect(err).ToNot(HaveOccurred())
	})
})
