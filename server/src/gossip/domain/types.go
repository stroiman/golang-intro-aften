package domain

type Message struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func (m Message) IsValidInput() bool {
	return m.Message != ""
}
