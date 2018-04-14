package application_test

import (
	"errors"
	. "gossip/application"
	. "gossip/application/mock_application"
	"gossip/domain"
	"gossip/testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application", func() {
	var ctrl *gomock.Controller
	var dataAccessMock *MockDataAccess
	var queueMock *MockQueueing
	var app Application

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dataAccessMock = NewMockDataAccess(ctrl)
		queueMock = NewMockQueueing(ctrl)
		app = Application{
			dataAccessMock,
			queueMock,
		}
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("GetMessages", func() {
		It("Returns the messages from the data access layer", func() {
			messages := []domain.Message{
				testing.NewMessage(),
				testing.NewMessage(),
			}
			dataAccessMock.EXPECT().GetMessages().Return(messages, nil)
			result, err := app.GetMessages()
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(messages))
		})
	})

	Describe("CreateMessage", func() {
		var (
			insertMessageCall  *gomock.Call
			publishMessageCall *gomock.Call
			insertedMessage    domain.Message
			publishedMessage   domain.Message
		)

		BeforeEach(func() {
			insertMessageCall = dataAccessMock.EXPECT().
				InsertMessage(gomock.Any()).
				Return(nil).AnyTimes().
				Do(func(m domain.Message) { insertedMessage = m })
			publishMessageCall = queueMock.EXPECT().
				PublishMessage(gomock.Any()).
				Return(nil).AnyTimes().
				Do(func(m domain.Message) { publishedMessage = m })
		})

		It("Saves a message in the database", func() {
			message := testing.NewMessage()
			message.Message = "Mock message"
			insertMessageCall.Times(1)
			err := app.InsertMessage(message)
			Expect(err).ToNot(HaveOccurred())
			Expect(insertedMessage.Message).To(Equal("Mock message"))
		})

		It("Assigns an Id to the message", func() {
			message := testing.NewMessage()
			message.Id = ""
			insertMessageCall.Times(1)
			err := app.InsertMessage(message)
			Expect(err).ToNot(HaveOccurred())
			Expect(insertedMessage.Id).ToNot(BeEmpty())
		})

		It("Publishes a message on a queue", func() {
			publishMessageCall.Times(1)
			message := testing.NewMessage()
			message.Message = "Mocked message"
			err := app.InsertMessage(message)
			Expect(err).ToNot(HaveOccurred())
			Expect(publishedMessage.Message).To(Equal("Mocked message"))
		})

		Describe("Publishing fails", func() {
			BeforeEach(func() {
				publishMessageCall.Times(1).Return(errors.New("Mock queue error"))
			})
			It("returns the error", func() {
				err := app.InsertMessage(testing.NewMessage())
				Expect(err).To(MatchError("Mock queue error"))
			})
		})

		Describe("Inserting in database fails", func() {
			BeforeEach(func() {
				insertMessageCall.Return(errors.New("Mock DB error"))
			})

			It("Doesn't publish the message", func() {
				publishMessageCall.Times(0)
				app.InsertMessage(testing.NewMessage())
			})

			It("Returns the error", func() {
				err := app.InsertMessage(testing.NewMessage())
				Expect(err).To(MatchError("Mock DB error"))
			})
		})
	})
})
