package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	t.Parallel()

	data := `{"operation":"PUT","data":"<some data>"}`
	message := NewMessage(data)

	expected := &Message{
		Operation: OperationPut,
		Data:      "<some data>",
	}
	assert.Exactly(t, expected, message)
}

func TestResponseNewMessage(t *testing.T) {
	t.Parallel()

	data := "Message received."
	message := NewResponseMessage(data)

	responseMessage := &Message{
		Operation: OperationMessage,
		Data:      data,
	}
	assert.Exactly(t, responseMessage, message)
}

func TestNewOKMessage(t *testing.T) {
	t.Parallel()

	message := NewOKMessage()

	okMessage := &Message{
		Operation: OperationOK,
		Data:      MessageEmpty,
	}

	assert.Exactly(t, okMessage, message)
}

func TestErrorMessage(t *testing.T) {
	t.Parallel()

	message := NewErrorMessage()

	errorMessage := &Message{
		Operation: OperationError,
		Data:      MessageError,
	}

	assert.Exactly(t, errorMessage, message)
}

func TestEmptyMessage(t *testing.T) {
	t.Parallel()

	message := NewEmptyMessage()

	emptyMessage := &Message{
		Operation: OperationEmpty,
		Data:      MessageEmpty,
	}

	assert.Exactly(t, emptyMessage, message)
}

func TestPutMessage(t *testing.T) {
	t.Parallel()

	data := "Message sending"
	message := NewPutMessage(data)

	putMessage := &Message{
		Operation: OperationPut,
		Data:      data,
	}

	assert.Exactly(t, putMessage, message)
}

func TestGetMessage(t *testing.T) {
	t.Parallel()

	message := NewGetMessage()

	getMessage := &Message{
		Operation: OperationGet,
		Data:      MessageEmpty,
	}

	assert.Exactly(t, getMessage, message)
}

func TestConsumerMessage(t *testing.T) {
	t.Parallel()

	message := NewConsumerMessage()

	getMessage := &Message{
		Operation: OperationConsumer,
		Data:      MessageEmpty,
	}

	assert.Exactly(t, getMessage, message)
}

func TestProducerMessage(t *testing.T) {
	t.Parallel()

	message := NewProducerMessage()

	getMessage := &Message{
		Operation: OperationProducer,
		Data:      MessageEmpty,
	}

	assert.Exactly(t, getMessage, message)
}

func TestToUnpack(t *testing.T) {
	t.Parallel()

	message := &Message{}
	msg := `{"operation":"PUT","data":"Ja chegou o disco voador..."}`

	message.ToUnpack(msg)

	expected := &Message{
		"PUT",
		"Ja chegou o disco voador...",
	}
	assert.Exactly(t, expected, message)
}

func TestToPack(t *testing.T) {
	t.Parallel()

	message := &Message{
		"GET",
		"Ja se foi o disco voador...",
	}

	msg := message.ToPack()

	expected := `{"operation":"GET","data":"Ja se foi o disco voador..."}`
	assert.Exactly(t, expected, msg)
}

func TestIsEmptyData(t *testing.T) {
	t.Parallel()

	message := &Message{
		Operation: "GET",
		Data:      "",
	}

	isEmptyData := message.IsEmptyData()

	assert.True(t, isEmptyData)
}
