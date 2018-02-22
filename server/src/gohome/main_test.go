package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http/httptest"
)

func getPath(path string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/ping", nil)
	handler := CreateRootHandler()
	handler.ServeHTTP(recorder, request)
	return recorder
}

var _ = Describe("Main", func() {
	It("responds to /ping", func() {
		recorder := getPath("/ping")
		Expect(recorder.Body.String()).To(Equal("pong"))
	})

	It("Serves /", func() {
		recorder := getPath("/")
		Expect(recorder.Code).To(Equal(200))
	})
	It("Serves /foobar", func() {
		recorder := getPath("/")
		Expect(recorder.Code).To(Equal(200))
	})
})
