package application

import (
	"gossip/domain"
)

type DataAccess interface {
	GetMessages() ([]domain.Message, error)
	InsertMessage(domain.Message) error
}

type Application struct {
	DataAccess DataAccess
}

func NewApplication() Application {
	return Application{}
}

func (app Application) GetMessages() ([]domain.Message, error) {
	return app.DataAccess.GetMessages()
}

func (app Application) InsertMessage(msg domain.Message) error {
	return app.DataAccess.InsertMessage(msg)
}
