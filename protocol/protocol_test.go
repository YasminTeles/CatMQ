package protocol

import (
	"testing"

	"github.com/YasminTeles/CatMQ/message"
	"github.com/YasminTeles/CatMQ/queue"
	"github.com/stretchr/testify/assert"
)

func TestHandlePutMessage(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"PUT","data":"bad word"}`
	queue := queue.NewQueue()

	response := HandleMessage(msg, queue)

	okMessage := `{"operation":"OK","data":""}`
	assert.Equal(t, okMessage, response)
}

func TestHandleGetMessage(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"GET","data":""}`
	queue := queue.NewQueue()
	queue.Push("cool word")

	response := HandleMessage(msg, queue)

	responseMessage := `{"operation":"MSG","data":"cool word"}`
	assert.Equal(t, responseMessage, response)
}

func TestHandleErrorMessage(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"POST","data":"bad word"}`
	queue := queue.NewQueue()

	response := HandleMessage(msg, queue)

	errorMessage := `{"operation":"ERR","data":"Operation failed!"}`
	assert.Equal(t, errorMessage, response)
}

func TestPublish(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	msg := &message.Message{
		Operation: message.OperationPut,
		Data:      "Que dia lindo",
	}

	response := publish(msg, queue)

	okMessage := message.NewOKMessage()
	assert.Exactly(t, okMessage, response)
}

func TestErrorPublish(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	msg := &message.Message{}

	response := publish(msg, queue)

	errorMessage := message.NewErrorMessage()
	assert.Exactly(t, errorMessage, response)
}

func TestEmptQueueGet(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()

	response := get(queue)

	emptyQueueMesage := message.NewEmptyMessage()
	assert.Exactly(t, emptyQueueMesage, response)
}

func TestGet(t *testing.T) {
	t.Parallel()

	queue := queue.NewQueue()
	queue.Push("Alejandro")
	queue.Push("Diego")

	response := get(queue)

	dataMessage := message.NewResponseMessage("Alejandro")
	assert.Exactly(t, dataMessage, response)
}
