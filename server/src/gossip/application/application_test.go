package application_test

import (
	"errors"
	. "gossip/application"
	. "gossip/application/mock_application"
	"gossip/domain"
	"gossip/testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func parseDate(input string) time.Time {
	r, err := time.Parse("2006-01-02", input)
	if err != nil {
		GinkgoT().Error("Invalid test date")
	}
	return r
}

var _ = Describe("Application", func() {
	var (
		ctrl               *gomock.Controller
		dataAccessMock     *MockDataAccess
		queueMock          *MockQueueing
		app                Application
		publishMessageCall *gomock.Call
		publishedMessage   domain.Message
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dataAccessMock = NewMockDataAccess(ctrl)
		queueMock = NewMockQueueing(ctrl)
		app = Application{
			dataAccessMock,
			queueMock,
		}
		publishedMessage = domain.Message{}
		publishMessageCall = queueMock.EXPECT().
			PublishMessage(gomock.Any()).
			Return(nil).AnyTimes().
			Do(func(m domain.Message) { publishedMessage = m })
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
			addMessageCall *gomock.Call
			addedMessage   domain.Message
		)

		BeforeEach(func() {
			addedMessage = domain.Message{}
			addMessageCall = dataAccessMock.EXPECT().
				InsertMessage(gomock.Any()).
				Return(nil).AnyTimes().
				Do(func(m domain.Message) { addedMessage = m })
		})

		Context("", func() {
			var (
				insertErr   error
				input       domain.Message
				returnedMsg domain.Message
			)

			BeforeEach(func() {
				returnedMsg = domain.Message{}
				input = testing.NewMessage()
				input.Id = ""
				input.Message = "Mock message"
			})

			JustBeforeEach(func() {
				returnedMsg, insertErr = app.AddMessage(input)
			})

			Context("Input already has an ID assigned", func() {
				BeforeEach(func() {
					input.Id = "Some dummy ID"
				})

				It("Returns an error", func() {
					Expect(insertErr).To(HaveOccurred())
				})
			})

			It("Retuns a message with an updated ID", func() {
				Expect(returnedMsg.Id).ToNot(BeEmpty())
			})

			Describe("No of calls", func() {
				BeforeEach(func() {
					addMessageCall.Times(1)
				})

				It("Inserts the right message exactly once", func() {
					Expect(addedMessage.Message).To(Equal("Mock message"))
				})
			})

			It("Succeeds", func() {
				Expect(insertErr).ToNot(HaveOccurred())
			})

			It("Assigns an Id to the message", func() {
				Expect(addedMessage.Id).ToNot(BeEmpty())
			})

			It("Sets CreatedAt", func() {
				Expect(addedMessage.CreatedAt).To(BeTemporally("~", time.Now(), time.Second))
			})

			Describe("Publishing", func() {
				BeforeEach(func() {
					publishMessageCall.Times(1)
					input.Message = "Mocked message"
				})

				It("Publishes a message on a queue", func() {
					Expect(publishedMessage.Message).To(Equal("Mocked message"))
				})

				Context("It fails", func() {
					BeforeEach(func() {
						publishMessageCall.Return(errors.New("Mock queue error"))
					})

					It("returns the error", func() {
						Expect(insertErr).To(MatchError("Mock queue error"))
					})
				})
			})

			Describe("Inserting in database fails", func() {
				BeforeEach(func() {
					addMessageCall.Return(errors.New("Mock DB error"))
				})

				It("Doesn't publish the message", func() {
					Expect(publishedMessage).To(Equal(domain.Message{}))
				})

				It("Returns the error", func() {
					Expect(insertErr).To(MatchError("Mock DB error"))
				})
			})
		})
	})

	Describe("UpdateMessage", func() {
		var (
			input             domain.Message
			updatedMessage    domain.Message
			updateMessageCall *gomock.Call
			updateErr         error
		)

		BeforeEach(func() {
			updatedMessage = domain.Message{}
			updateMessageCall =
				dataAccessMock.EXPECT().
					UpdateMessage(gomock.Any()).AnyTimes().
					Do(func(m domain.Message) { updatedMessage = m })
			input = testing.NewMessage()
			input.Id = "Message ID"
			input.Message = "Updated message"
			input.CreatedAt = parseDate("2018-01-01")
			input.EditedAt = nil
		})

		JustBeforeEach(func() {
			updateErr = app.UpdateMessage(input)
		})

		It("Updates the message in the database", func() {
			updateMessageCall.Times(1)
		})

		It("Sets the new message", func() {
			Expect(updatedMessage.Message).To(Equal("Updated message"))
		})

		It("Does not change 'CreatedAt'", func() {
			expected := parseDate("2018-01-01")
			Expect(updatedMessage.CreatedAt).To(BeTemporally("==", expected))
		})

		It("Sets 'EditedAt'", func() {
			minimumDate := time.Now()
			Expect(updatedMessage.EditedAt).ToNot(BeNil())
			Expect(*updatedMessage.EditedAt).To(BeTemporally("~", minimumDate, time.Second))
		})

		It("Publishes to the queue", func() {
			Expect(publishedMessage).To(Equal(updatedMessage))
		})

		It("Does not alter the ID", func() {
			Expect(updatedMessage.Id).To(Equal("Message ID"))
		})

		Describe("updating message fails in the data store", func() {
			BeforeEach(func() {
				updateMessageCall.Return(errors.New("Mocked error"))
			})

			It("does not publish a message", func() {
				Expect(publishedMessage).To(Equal(domain.Message{}))
			})

			It("returns the error", func() {
				Expect(updateErr).To(MatchError("Mocked error"))
			})
		})

		Describe("Publishing fails", func() {
			BeforeEach(func() {
				publishMessageCall.Return(errors.New("Mocked error"))
			})

			It("Returns the error", func() {
				err := app.UpdateMessage(input)
				Expect(err).To(MatchError("Mocked error"))
			})
		})
	})
})
