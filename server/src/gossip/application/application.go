package application

import (
	"gossip/domain"
)

type DataAccess interface {
	GetMessages() ([]domain.Message, error)
}

type Application struct {
	DataAccess
}

func NewApplication() Application {
	return Application{}
}

func (app Application) GetMessages() ([]domain.Message, error) {
	return app.DataAccess.GetMessages()
}
