package msgpack

import (
	m "github.com/rinosukmandityo/message-flow/models"

	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Message struct{}

func (u *Message) Decode(input []byte) (*m.Message, error) {
	msg := new(m.Message)
	if e := msgpack.Unmarshal(input, msg); e != nil {
		return nil, errors.Wrap(e, "serializer.Message.Decode")
	}
	return msg, nil
}

func (u *Message) Encode(input []m.Message) ([]byte, error) {
	rawMsg, e := msgpack.Marshal(input)
	if e != nil {
		return nil, errors.Wrap(e, "serializer.Message.Encode")
	}
	return rawMsg, nil
}
