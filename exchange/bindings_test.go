package exchange

import (
	"testing"
)

func TestRoute(t *testing.T) {
	t.Parallel()
}

// func TestHandlePutMessage(t *testing.T) {
// 	t.Parallel()

// 	msg := `{"operation":"PUT","data":"bad word"}`
// 	exchange := NewExchange()

// 	response := HandleMessage(msg, exchange)

// 	okMessage := `{"operation":"OK","data":""}`
// 	assert.Equal(t, okMessage, response)
// }

// func TestHandleGetMessage(t *testing.T) {
// 	t.Parallel()

// 	msg := `{"operation":"GET","data":""}`
// 	exchange := NewExchange()
// 	exchange.Push("cool word")

// 	response := HandleMessage(msg, exchange)

// 	responseMessage := `{"operation":"MSG","data":"cool word"}`
// 	assert.Equal(t, responseMessage, response)
// }

// func TestHandleErrorMessage(t *testing.T) {
// 	t.Parallel()

// 	msg := `{"operation":"POST","data":"bad word"}`
// 	exchange := NewExchange()

// 	response := HandleMessage(msg, exchange)

// 	errorMessage := `{"operation":"ERR","data":"Operation failed!"}`
// 	assert.Equal(t, errorMessage, response)
// }

// func TestHandleSetConsumerAddress(t *testing.T) {
// 	t.Parallel()

// 	msg := `{"operation":"CON","data":""}`
// 	exchange := NewExchange()

// 	response := HandleMessage(msg, exchange)

// 	okMessage := `{"operation":"OK","data":""}`
// 	assert.Equal(t, okMessage, response)
// }

// func TestHandleSetProducerAddress(t *testing.T) {
// 	t.Parallel()

// 	msg := `{"operation":"PRO","data":""}`
// 	exchange := NewExchange()

// 	response := HandleMessage(msg, exchange)

// 	okMessage := `{"operation":"OK","data":""}`
// 	assert.Equal(t, okMessage, response)
// }
