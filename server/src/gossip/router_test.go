package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gossip/domain"
)

type MessageList struct {
	initialMessages []domain.Message
	newMessages     []domain.Message
}

func (m *MessageList) GetMessages() []domain.Message {
	return m.initialMessages
}

func (m *MessageList) AddMessage(message domain.Message) {
	m.newMessages = append(m.newMessages, message)
}

func NewMessage() domain.Message {
	return domain.Message{
		Id: uuid.New().String(),
	}
}

var _ = Describe("Router", func() {
	Describe("/api/messages", func() {
		var (
			recorder *httptest.ResponseRecorder
			messages *MessageList
			request  *http.Request
		)

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			messages = &MessageList{}
			request = httptest.NewRequest("GET", "/api/messages", nil)
		})

		JustBeforeEach(func() {
			handler := NewMessageHandler(messages)
			handler.ServeHTTP(recorder, request)
		})

		It("Returns HTTP 200", func() {
			Expect(recorder.Code).To(Equal(200))
		})

		It("Returns Content-Type=application/json", func() {
			actual := recorder.Header().Get("Content-Type")
			Expect(actual).To(Equal("application/json"))
		})

		Context("Repository has two messages", func() {
			BeforeEach(func() {
				messages = &MessageList{
					initialMessages: []domain.Message{NewMessage(), NewMessage()},
				}
			})
			It("Has two objects", func() {
				var result []interface{}
				err := json.NewDecoder(recorder.Body).Decode(&result)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(HaveLen(2))
			})
		})

		Describe("POST a valid message", func() {
			BeforeEach(func() {
				input := `{
					"id": "42",
					"message": "foobar"
				}`
				body := strings.NewReader(input)
				request = httptest.NewRequest("POST", "/api/messages", body)
			})

			Context("a valid message", func() {
				It("Adds a message, if valid message posted", func() {
					Expect(messages.newMessages).To(HaveLen(1))
				})
			})
		})
	})
})
