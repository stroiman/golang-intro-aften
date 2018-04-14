package dataaccess_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "gossip/dataaccess"
	"gossip/testing"
)

var url = "postgres://gossip:gossip@127.0.0.1/gossip?sslmode=disable"

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
})
