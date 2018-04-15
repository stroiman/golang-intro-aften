package matchers

import (
	"fmt"

	. "github.com/onsi/gomega/types"
	"gossip/domain"
)

type MessageMatcher struct {
	message domain.Message
	matcher GomegaMatcher
}

func HaveMessage(matcher GomegaMatcher) *MessageMatcher { return &MessageMatcher{matcher: matcher} }

func (m *MessageMatcher) Match(actual interface{}) (success bool, err error) {
	if message, ok := actual.(domain.Message); ok {
		m.message = message
		return m.matcher.Match(message.Message)
	}
	return false, fmt.Errorf("MessageMatcher expects a message")
}

func (m *MessageMatcher) FailureMessage(actual interface{}) string {
	return m.matcher.FailureMessage(m.message.Message)
}

func (m *MessageMatcher) NegatedFailureMessage(actual interface{}) string {
	return m.matcher.NegatedFailureMessage(m.message.Message)
}
