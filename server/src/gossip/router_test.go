package main

import (
	"encoding/json"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	Describe("/api/messages", func() {
		var (
			recorder *httptest.ResponseRecorder
		)

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/api/messages", nil)
			handler := createRootHandler()
			handler.ServeHTTP(recorder, request)
		})

		It("Returns HTTP 200", func() {
			Expect(recorder.Code).To(Equal(200))
		})

		It("Returns Content-Type=application/json", func() {
			actual := recorder.Header().Get("Content-Type")
			Expect(actual).To(Equal("application/json"))
		})

		It("Has two objects", func() {
			var result []interface{}
			err := json.NewDecoder(recorder.Body).Decode(&result)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(HaveLen(2))
		})
	})
})
