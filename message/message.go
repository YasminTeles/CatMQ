package message

import (
	"encoding/json"
	"log"
)

const (
	OperationPut      = "PUT"
	OperationGet      = "GET"
	OperationEmpty    = "EMP"
	OperationError    = "ERR"
	OperationOK       = "OK"
	OperationMessage  = "MSG"
	OperationConsumer = "CON"
	OperationProducer = "PRO"

	MessageEmpty = ""
	MessageError = "Operation failed!"
)

type Message struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
}

func NewMessage(data string) *Message {
	msg := &Message{}
	msg.ToUnpack(data)

	return msg
}

func NewResponseMessage(data string) *Message {
	return &Message{
		Operation: OperationMessage,
		Data:      data,
	}
}

func NewOKMessage() *Message {
	return &Message{
		Operation: OperationOK,
		Data:      MessageEmpty,
	}
}

func NewErrorMessage() *Message {
	return &Message{
		Operation: OperationError,
		Data:      MessageError,
	}
}

func NewEmptyMessage() *Message {
	return &Message{
		Operation: OperationEmpty,
		Data:      MessageEmpty,
	}
}

func NewPutMessage(data string) *Message {
	return &Message{
		Operation: OperationPut,
		Data:      data,
	}
}

func NewGetMessage() *Message {
	return &Message{
		Operation: OperationGet,
		Data:      MessageEmpty,
	}
}

func NewConsumerMessage() *Message {
	return &Message{
		Operation: OperationConsumer,
		Data:      MessageEmpty,
	}
}

func NewProducerMessage() *Message {
	return &Message{
		Operation: OperationProducer,
		Data:      MessageEmpty,
	}
}

func (message *Message) ToUnpack(data string) {
	if err := json.Unmarshal([]byte(data), message); err != nil {
		log.Panicf("Some unpack error: %s.\n", err)
	}
}

func (message *Message) ToPack() string {
	data, err := json.Marshal(message)
	if err != nil {
		log.Panicf("Some pack error: %s.\n", err)
	}

	return string(data)
}

func (message *Message) IsEmptyData() bool {
	return !(len(message.Data) > 0)
}
