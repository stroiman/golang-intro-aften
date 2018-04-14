package application

import (
	"gossip/domain"
)

type Queueing interface {
	PublishMessage(domain.Message) error
}

type DataAccess interface {
	GetMessages() ([]domain.Message, error)
	InsertMessage(domain.Message) error
}

type Application struct {
	DataAccess DataAccess
	Queueing   Queueing
}

func NewApplication() Application {
	return Application{}
}

func (app Application) GetMessages() ([]domain.Message, error) {
	return app.DataAccess.GetMessages()
}

func (app Application) InsertMessage(msg domain.Message) error {
	err := app.DataAccess.InsertMessage(msg)
	if err == nil {
		err = app.Queueing.PublishMessage(msg)
	}
	return err
}
