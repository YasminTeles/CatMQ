package exchange

import (
	"github.com/YasminTeles/CatMQ/queue"
)

type Exchange struct {
	unprocessed *queue.Queue
	processed   *queue.Queue
}

func NewExchange() *Exchange {
	return &Exchange{
		unprocessed: queue.NewQueue(),
		processed:   queue.NewQueue(),
	}
}

func (exch *Exchange) GetQueue(addr *Address, action string) *queue.Queue {
	if isUnprocessedQueue(addr, action) {
		return exch.unprocessed
	} else if isProcessedQueue(addr, action) {
		return exch.processed
	}

	return nil
}

func isUnprocessedQueue(addr *Address, action string) bool {
	return (addr.IsConsumer() && action == RouteGet) ||
		(addr.IsProducer() && action == RoutePut)
}

func isProcessedQueue(addr *Address, action string) bool {
	return (addr.IsConsumer() && action == RoutePut) ||
		(addr.IsProducer() && action == RouteGet)
}
