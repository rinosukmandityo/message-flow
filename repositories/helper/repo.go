package repohelper

import (
	"log"
	"os"

	repo "github.com/rinosukmandityo/message-flow/repositories"
	mr "github.com/rinosukmandityo/message-flow/repositories/memory"
)

func ChooseRepo() repo.MessageRepository {
	switch os.Getenv("driver") {
	default:
		repo, e := mr.NewMessageRepository()
		if e != nil {
			log.Fatal(e)
		}
		return repo
	}
	return nil
}
