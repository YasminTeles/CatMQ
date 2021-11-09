package exchange

import (
	"github.com/YasminTeles/CatMQ/message"
)

const (
	RouteGet = "GET"
	RoutePut = "PUT"
)

func Route(msg string, exchange *Exchange, address *Address) string {
	request := message.NewMessage(msg)

	var response *message.Message

	switch request.Operation {
	case message.OperationConsumer:
		response = SetAddressing(request, address)

	case message.OperationProducer:
		response = SetAddressing(request, address)

	case message.OperationPut:
		response = put(request, exchange, address)

	case message.OperationGet:
		response = get(exchange, address)

	default:
		response = message.NewErrorMessage()
	}

	return response.ToPack()
}

func put(request *message.Message, exchange *Exchange, address *Address) *message.Message {
	queue := exchange.GetQueue(address, RoutePut)

	return Publish(request, queue)
}

func get(exchange *Exchange, address *Address) *message.Message {
	queue := exchange.GetQueue(address, RouteGet)

	return Get(queue)
}
