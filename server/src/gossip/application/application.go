package application

import (
	"gossip/domain"
)

type Application struct {
}

func NewApplication() Application {
	return Application{}
}

func (app Application) GetMessages() ([]domain.Message, error) {
	return nil, nil
}
