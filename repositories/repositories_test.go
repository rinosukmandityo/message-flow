// +build message_repo

package repositories_test

import (
	"testing"

	m "github.com/rinosukmandityo/message-flow/models"
	. "github.com/rinosukmandityo/message-flow/repositories"
	rh "github.com/rinosukmandityo/message-flow/repositories/helper"
)

var (
	repo MessageRepository
)

func MessageTestData() []m.Message {
	return []m.Message{
		{Message: "Message 01"},
		{Message: "Message 02"},
		{Message: "Message 03"},
	}
}

func init() {
	repo = rh.ChooseRepo()
}

func TestMessageService(t *testing.T) {
	t.Run("Insert Message", InsertMessage)
	t.Run("Get Message", GetMEssage)
}

func InsertMessage(t *testing.T) {
	testdata := MessageTestData()

	t.Run("Case 1: Save data", func(t *testing.T) {
		for _, _data := range testdata {
			if e := repo.Store(&_data); e != nil {
				t.Errorf("[ERROR] - Failed to save data %s ", e.Error())
			}
		}
	})
}

func GetMEssage(t *testing.T) {
	t.Run("Case 1: Get data", func(t *testing.T) {
		if _, e := repo.GetAll(); e != nil {
			t.Errorf("[ERROR] - Failed to get data %s ", e.Error())
		}
	})
}
