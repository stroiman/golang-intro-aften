package application_test

import (
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
	var mock *MockDataAccess
	var app Application

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mock = NewMockDataAccess(ctrl)
		app = Application{
			mock,
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
			mock.EXPECT().GetMessages().Return(messages, nil)
			result, err := app.GetMessages()
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(Equal(messages))
		})
	})
})
