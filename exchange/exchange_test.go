package exchange

import (
	"testing"

	"github.com/YasminTeles/CatMQ/queue"
	"github.com/stretchr/testify/assert"
)

func TestNewExchange(t *testing.T) {
	t.Parallel()

	exchange := NewExchange()

	newExchange := &Exchange{
		unprocessed: queue.NewQueue(),
		processed:   queue.NewQueue(),
	}

	assert.Exactly(t, newExchange, exchange)
}

func TestGetQueue(t *testing.T) {
	t.Parallel()

	consumer := NewAddress()
	consumer.SetConsumer()

	producer := NewAddress()
	producer.SetProducer()

	ex := NewExchange()
	unprocessed := queue.NewQueue()
	unprocessed.Push("Alejandro")
	ex.unprocessed = unprocessed

	processed := queue.NewQueue()
	processed.Push("Fernando")
	ex.processed = processed

	cases := []struct {
		kind   *Address
		action string
		want   *queue.Queue
	}{
		{consumer, RouteGet, unprocessed},
		{consumer, RoutePut, processed},
		{producer, RouteGet, processed},
		{producer, RoutePut, unprocessed},
	}

	for _, test := range cases {
		queue := ex.GetQueue(test.kind, test.action)

		assert.Exactly(t, test.want, queue)
	}
}

func TestIsUnprocessedQueue(t *testing.T) {
	t.Parallel()

	consumer := NewAddress()
	consumer.SetConsumer()

	producer := NewAddress()
	producer.SetProducer()

	cases := []struct {
		kind   *Address
		action string
		want   bool
	}{
		{consumer, RouteGet, true},
		{consumer, RoutePut, false},
		{producer, RouteGet, false},
		{producer, RoutePut, true},
	}

	for _, test := range cases {
		isUnprocessedQueue := isUnprocessedQueue(test.kind, test.action)

		assert.Equal(t, test.want, isUnprocessedQueue)
	}
}

func TestIsProcessedQueue(t *testing.T) {
	t.Parallel()

	consumer := NewAddress()
	consumer.SetConsumer()

	producer := NewAddress()
	producer.SetProducer()

	cases := []struct {
		kind   *Address
		action string
		want   bool
	}{
		{consumer, RouteGet, false},
		{consumer, RoutePut, true},
		{producer, RouteGet, true},
		{producer, RoutePut, false},
	}

	for _, test := range cases {
		isProcessedQueue := isProcessedQueue(test.kind, test.action)

		assert.Equal(t, test.want, isProcessedQueue)
	}
}
