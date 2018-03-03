package domain

import "time"

type Message struct {
	Id        string    `json:"id"`
	Message   string    `json:"message"`
	UserName  string    `json:"userName"`
	CreatedAt time.Time `json:"createdAt"`
}

func (m Message) IsValidInput() bool {
	return m.Message != ""
}
