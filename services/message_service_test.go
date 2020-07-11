// +build message_service

package services_test

import (
	"testing"

	m "github.com/rinosukmandityo/message-flow/models"
	rh "github.com/rinosukmandityo/message-flow/repositories/helper"
	. "github.com/rinosukmandityo/message-flow/services"
	"github.com/rinosukmandityo/message-flow/services/logic"
)

var (
	messageService MessageService
)

func MessageTestData() []m.Message {
	return []m.Message{
		{Message: "Message 01"},
		{Message: "Message 02"},
		{Message: "Message 03"},
	}
}

func init() {
	repo := rh.ChooseRepo()
	messageService = logic.NewMessageService(repo)
}

func TestMessageService(t *testing.T) {
	t.Run("Insert Message", InsertMessage)
	t.Run("Get Message", GetMEssage)
}

func InsertMessage(t *testing.T) {
	testdata := MessageTestData()

	t.Run("Case 1: Save data", func(t *testing.T) {
		for _, _data := range testdata {
			if e := messageService.Store(&_data); e != nil {
				t.Errorf("[ERROR] - Failed to save data %s ", e.Error())
			}
		}
	})
}

func GetMEssage(t *testing.T) {
	t.Run("Case 1: Get data", func(t *testing.T) {
		if _, e := messageService.GetAll(); e != nil {
			t.Errorf("[ERROR] - Failed to get data %s ", e.Error())
		}
	})
}
