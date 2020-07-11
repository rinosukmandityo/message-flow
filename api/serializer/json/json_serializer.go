package json

import (
	m "github.com/rinosukmandityo/message-flow/models"

	"encoding/json"
	"github.com/pkg/errors"
)

type Message struct{}

func (u *Message) Decode(input []byte) (*m.Message, error) {
	msg := new(m.Message)
	if e := json.Unmarshal(input, msg); e != nil {
		return nil, errors.Wrap(e, "serializer.Message.Decode")
	}
	return msg, nil
}

func (u *Message) Encode(input []m.Message) ([]byte, error) {
	rawMsg, e := json.Marshal(input)
	if e != nil {
		return nil, errors.Wrap(e, "serializer.Message.Encode")
	}
	return rawMsg, nil
}
