package application

import (
	"errors"
	"fmt"
	"time"

	"gossip/domain"

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
	DataAccess DataAccess `inject:""`
	Queueing   Queueing   `inject:""`
}

func NewApplication() Application {
	return Application{}
}

func (app Application) GetMessages() ([]domain.Message, error) {
	return app.DataAccess.GetMessages()
}

func (app Application) AddMessage(msg domain.Message) (domain.Message, error) {
	defer func() {
		r := recover()
		fmt.Println("Exiting", r)
	}()
	fmt.Println("Application.AddMessage", app.DataAccess)
	if msg.Id != "" {
		return msg, errors.New("Cannot create a message that already has an ID")
	}
	msg.Id = uuid.New().String()
	msg.CreatedAt = time.Now()
	err := app.DataAccess.InsertMessage(msg)
	if err == nil {
		err = app.Queueing.PublishMessage(msg)
	}
	return msg, err
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
