package main

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gossip/domain"
)

type MessageList []domain.Message

func (m MessageList) GetMessages() []domain.Message { return m }

func NewMessage() domain.Message {
	return domain.Message{
		Id: uuid.New().String(),
	}
}

var _ = Describe("Router", func() {
	Describe("/api/messages", func() {
		var (
			recorder *httptest.ResponseRecorder
			messages []domain.Message
		)

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			messages = make([]domain.Message, 0)
		})

		JustBeforeEach(func() {
			request := httptest.NewRequest("GET", "/api/messages", nil)
			handler := NewMessageHandler(MessageList(messages))
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
				messages = append(messages, NewMessage(), NewMessage())
			})
			It("Has two objects", func() {
				var result []interface{}
				err := json.NewDecoder(recorder.Body).Decode(&result)
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(HaveLen(2))
			})
		})
	})
})
