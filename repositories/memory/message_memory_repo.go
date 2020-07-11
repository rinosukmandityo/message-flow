package memory

import (
	"github.com/rinosukmandityo/message-flow/helper"
	m "github.com/rinosukmandityo/message-flow/models"
	repo "github.com/rinosukmandityo/message-flow/repositories"

	"github.com/pkg/errors"
)

type messageMemoryRepository struct {
	data []m.Message
}

func NewMessageRepository() (repo.MessageRepository, error) {
	repo := &messageMemoryRepository{
		data: []m.Message{},
	}
	return repo, nil
}

func (r *messageMemoryRepository) GetAll() ([]m.Message, error) {
	res := r.data
	if len(res) == 0 {
		return res, errors.Wrap(helper.ErrDataNotFound, "repository.Message.GetAll")
	}
	return res, nil
}

func (r *messageMemoryRepository) Store(data *m.Message) error {
	r.data = append(r.data, *data)

	return nil
}
