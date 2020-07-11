package serializer

import (
	m "github.com/rinosukmandityo/message-flow/models"
)

type MessageSerializer interface {
	Decode(input []byte) (*m.Message, error)
	Encode(input []m.Message) ([]byte, error)
}
