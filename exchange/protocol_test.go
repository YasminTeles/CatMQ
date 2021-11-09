package exchange

import (
	"testing"

	"github.com/YasminTeles/CatMQ/message"
	"github.com/YasminTeles/CatMQ/queue"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	data := "Que dia lindo"
	msg := message.NewPutMessage(data)

	response := Publish(msg, queue)

	okMessage := message.NewOKMessage()
	assert.Exactly(t, okMessage, response)
}

func TestErrorPublish(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	msg := &message.Message{}

	response := Publish(msg, queue)

	errorMessage := message.NewErrorMessage()
	assert.Exactly(t, errorMessage, response)
}

func TestEmptyQueueGet(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()

	response := Get(queue)

	emptyQueueMesage := message.NewEmptyMessage()
	assert.Exactly(t, emptyQueueMesage, response)
}

func TestGet(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	queue.Push("Alejandro")
	queue.Push("Diego")

	response := Get(queue)

	dataMessage := message.NewResponseMessage("Alejandro")
	assert.Exactly(t, dataMessage, response)
}

func TestSetAddressing(t *testing.T) {
	t.Parallel()

	msg := message.NewConsumerMessage()
	address := NewAddress()

	responseMessage := SetAddressing(msg, address)

	okMessage := message.NewOKMessage()

	assert.Exactly(t, okMessage, responseMessage)
}
