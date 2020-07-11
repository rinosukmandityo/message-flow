package logic

import (
	m "github.com/rinosukmandityo/message-flow/models"
	repo "github.com/rinosukmandityo/message-flow/repositories"
	svc "github.com/rinosukmandityo/message-flow/services"
)

type messageService struct {
	messageRepo repo.MessageRepository
}

func NewMessageService(messageRepo repo.MessageRepository) svc.MessageService {
	return &messageService{
		messageRepo,
	}
}

func (u *messageService) GetAll() ([]m.Message, error) {
	res, e := u.messageRepo.GetAll()
	if e != nil {
		return res, e
	}
	return res, nil
}

func (u *messageService) Store(data *m.Message) error {
	return u.messageRepo.Store(data)

}
