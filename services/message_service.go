package services

import (
	m "github.com/rinosukmandityo/message-flow/models"
)

type MessageService interface {
	GetAll() ([]m.Message, error)
	Store(data *m.Message) error
}
