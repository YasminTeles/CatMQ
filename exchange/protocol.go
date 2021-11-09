package exchange

import (
	"github.com/YasminTeles/CatMQ/message"
	"github.com/YasminTeles/CatMQ/queue"
)

func Publish(msg *message.Message, queue *queue.Queue) *message.Message {
	if msg.IsEmptyData() {
		return message.NewErrorMessage()
	}

	if queue == nil {
		return message.NewErrorMessage()
	}

	queue.Push(msg.Data)

	return message.NewOKMessage()
}

func Get(q *queue.Queue) *message.Message {
	if q == nil {
		return message.NewErrorMessage()
	}

	data, err := q.Pop()
	if err != nil {
		if err == queue.ErrEmptyQueue {
			return message.NewEmptyMessage()
		}
		return message.NewErrorMessage()
	}

	return message.NewResponseMessage(data.(string))
}

func SetAddressing(msg *message.Message, address *Address) *message.Message {
	switch msg.Operation {
	case message.OperationConsumer:
		address.SetConsumer()

	case message.OperationProducer:
		address.SetProducer()
	}

	return message.NewOKMessage()
}
