package exchange

import (
	"testing"

	"github.com/YasminTeles/CatMQ/message"
	"github.com/stretchr/testify/assert"
)

func TestRoutePut(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"PUT","data":"bad word"}`
	exchange := NewExchange()
	addr := NewAddress()

	response := Route(msg, exchange, addr)

	okMessage := `{"operation":"OK","data":""}`
	assert.Equal(t, okMessage, response)
}

func TestRouteGet(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"GET","data":""}`
	exchange := NewExchange()
	exchange.unprocessed.Push("cool word")

	addr := NewAddress()
	addr.SetConsumer()

	response := Route(msg, exchange, addr)

	responseMessage := `{"operation":"MSG","data":"cool word"}`
	assert.Equal(t, responseMessage, response)
}

func TestRouteError(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"POST","data":"bad word"}`
	exchange := NewExchange()
	addr := NewAddress()

	response := Route(msg, exchange, addr)

	errorMessage := `{"operation":"ERR","data":"Operation failed!"}`
	assert.Equal(t, errorMessage, response)
}

func TestRouteSetConsumerAddress(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"CON","data":""}`
	exchange := NewExchange()
	addr := NewAddress()

	response := Route(msg, exchange, addr)

	okMessage := `{"operation":"OK","data":""}`
	assert.Equal(t, okMessage, response)
}

func TestRouteSetProducerAddress(t *testing.T) {
	t.Parallel()

	msg := `{"operation":"PRO","data":""}`
	exchange := NewExchange()
	addr := NewAddress()

	response := Route(msg, exchange, addr)

	okMessage := `{"operation":"OK","data":""}`
	assert.Equal(t, okMessage, response)
}

func TestHandlePut(t *testing.T) {
	t.Parallel()

	request := message.NewPutMessage("<some data>")
	exchange := NewExchange()
	addr := NewAddress()

	response := put(request, exchange, addr)

	okMessage := message.NewOKMessage()
	assert.Exactly(t, okMessage, response)
}

func TestHandleGet(t *testing.T) {
	t.Parallel()

	request := message.NewGetMessage()
	exchange := NewExchange()
	addr := NewAddress()

	response := put(request, exchange, addr)

	errorMessage := message.NewErrorMessage()
	assert.Exactly(t, errorMessage, response)
}
