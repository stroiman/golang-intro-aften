package repository_test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//. "gossip/repository"

	"gossip/domain"
)

var _ = Describe("MessageRepository", func() {
	It("Posts to AWS", func() {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("eu-west-1"),
			Credentials: credentials.NewSharedCredentials("/Users/peter/.aws/credentials", "ps"),
		})
		Expect(err).ToNot(HaveOccurred())
		svc := dynamodb.New(sess)

		tabinput := &dynamodb.ListTablesInput{}
		result, err := svc.ListTables(tabinput)
		Expect(err).ToNot(HaveOccurred())
		fmt.Println("TABLES", result)

		msg := domain.Message{}
		info, err := dynamodbattribute.MarshalMap(msg)
		Expect(err).ToNot(HaveOccurred())
		input := &dynamodb.PutItemInput{
			Item:      info,
			TableName: aws.String("Gossip_Messages"),
		}
		_, err = svc.PutItem(input)
		Expect(err).ToNot(HaveOccurred())
	})
})
