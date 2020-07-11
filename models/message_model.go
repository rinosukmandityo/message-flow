package models

type Message struct {
	Message string `json:"Message" bson:"Message" msgpack:"Message" db:"Message"`
}

func NewMessage() *Message {
	m := new(Message)
	return m
}

func (m *Message) TableName() string {
	return "message"
}
