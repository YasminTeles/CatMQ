package protocol

import (
	"github.com/YasminTeles/CatMQ/message"
	"github.com/YasminTeles/CatMQ/queue"
)

func HandleMessage(msg string, queue *queue.Queue) string {
	request := message.NewMessage(msg)

	var response *message.Message

	switch request.Operation {
	case message.OperationPut:
		response = publish(request, queue)

	case message.OperationGet:
		response = get(queue)

	default:
		response = message.NewErrorMessage()
	}

	return response.ToPack()
}

func publish(msg *message.Message, queue *queue.Queue) *message.Message {
	if msg.IsEmptyData() {
		return message.NewErrorMessage()
	}

	queue.Push(msg.Data)

	return message.NewOKMessage()
}

func get(queue *queue.Queue) *message.Message {
	if queue.IsEmpty() {
		return message.NewEmptyMessage()
	}

	data, err := queue.Pop()
	if err != nil {
		return message.NewErrorMessage()
	}

	return message.NewResponseMessage(data.(string))
}
