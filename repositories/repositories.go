package repositories

import (
	m "github.com/rinosukmandityo/message-flow/models"
)

type MessageRepository interface {
	GetAll() ([]m.Message, error)
	Store(data *m.Message) error
}
