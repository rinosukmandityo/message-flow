package api

import (
	slz "github.com/rinosukmandityo/message-flow/api/serializer"
	js "github.com/rinosukmandityo/message-flow/api/serializer/json"
	ms "github.com/rinosukmandityo/message-flow/api/serializer/msgpack"
)

var (
	ContentTypeJson    = "application/json"
	ContentTypeMsgPack = "application/x-msgpack"
)

func GetSerializer(contentType string) slz.MessageSerializer {
	if contentType == ContentTypeMsgPack {
		return &ms.Message{}
	}
	return &js.Message{}
}
