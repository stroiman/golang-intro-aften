package application

import (
	"gossip/domain"
	"time"

	"github.com/google/uuid"
)

type Queueing interface {
	PublishMessage(domain.Message) error
}

type DataAccess interface {
	GetMessages() ([]domain.Message, error)
	InsertMessage(domain.Message) error
	UpdateMessage(domain.Message) error
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
	msg.Id = uuid.New().String()
	msg.CreatedAt = time.Now()
	err := app.DataAccess.InsertMessage(msg)
	if err == nil {
		err = app.Queueing.PublishMessage(msg)
	}
	return err
}

func (app Application) UpdateMessage(msg domain.Message) error {
	msg.EditedAt = new(time.Time)
	*msg.EditedAt = time.Now()
	err := app.DataAccess.UpdateMessage(msg)
	if err != nil {
		return err
	}
	return app.Queueing.PublishMessage(msg)
}
