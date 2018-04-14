package dataaccess

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gossip/domain"
	"gossip/testing"
)

var url = "postgres://gossip:gossip@127.0.0.1/gossip?sslmode=disable"

func Must(err error) {
	Expect(err).ToNot(HaveOccurred())
}

func mustParseTime(value string) time.Time {
	result, err := time.Parse("2006-01-02T15:04:05", value)
	Expect(err).ToNot(HaveOccurred())
	return result
}

var _ = Describe("Dataaccess", func() {
	var conn Connection

	BeforeSuite(func() {
		var err error
		conn, err = NewConnection(url)
		Expect(err).ToNot(HaveOccurred())
		Expect(conn.Migrate()).To(Succeed())
	})

	Describe("Insert", func() {
		It("Creates a readable record", func() {
			message := testing.NewMessage()
			Expect(conn.InsertMessage(message)).To(Succeed())
			msg, err := conn.GetMessage(message.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(msg).To(Equal(message))
		})
	})

	Describe("Update", func() {
		It("Updates the record", func() {
			message := testing.NewMessage()
			message.Message = "Old message"
			Expect(conn.InsertMessage(message)).To(Succeed())
			message.Message = "New message"
			Expect(conn.UpdateMessage(message)).To(Succeed())

			msg, err := conn.GetMessage(message.Id)
			Expect(err).ToNot(HaveOccurred())
			Expect(msg).To(Equal(message))
		})
	})

	Describe("GetMessages", func() {
		BeforeEach(func() {
			_, err := conn.db.Exec("delete from messages")
			Expect(err).ToNot(HaveOccurred())
		})

		It("Retrieves all messages in create order", func() {
			// Setup
			m1 := testing.NewMessage()
			m2 := testing.NewMessage()
			m3 := testing.NewMessage()
			m1.CreatedAt = mustParseTime("2018-01-01T12:00:00")
			m2.CreatedAt = mustParseTime("2018-01-01T12:01:00")
			m3.CreatedAt = mustParseTime("2018-01-01T12:02:00")
			Must(conn.InsertMessage(m1))
			Must(conn.InsertMessage(m3))
			Must(conn.InsertMessage(m2))

			// Exercise
			result, err := conn.GetMessages()

			// Verify
			Expect(err).ToNot(HaveOccurred())
			expected := []domain.Message{m1, m2, m3}
			Expect(result).To(Equal(expected))
		})
	})
})
