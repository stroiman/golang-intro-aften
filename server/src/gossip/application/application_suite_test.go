//go:generate mockgen -source=application.go -destination=mock_application/application.go
//go:generate mockgen -source=message_hub.go -destination=mock_application/message_hub.go
package application_test

// In order to update the mocks, you need the mockgen tool
// go get github.com/golang/mock/mockgen

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApplication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Application Suite")
}
