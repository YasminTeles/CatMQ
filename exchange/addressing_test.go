package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	t.Parallel()

	address := NewAddress()

	newAddress := &Address{
		kind: AddressProducer,
	}
	assert.Exactly(t, newAddress, address)
}

func TestSetConsumer(t *testing.T) {
	t.Parallel()

	address := NewAddress()

	address.SetConsumer()

	assert.Equal(t, AddressConsumer, address.kind)
}

func TestSetProducer(t *testing.T) {
	t.Parallel()

	address := NewAddress()

	address.SetProducer()

	assert.Equal(t, AddressProducer, address.kind)
}

func TestIsConsumer(t *testing.T) {
	t.Parallel()

	address := NewAddress()
	address.SetConsumer()

	isConsumer := address.IsConsumer()

	assert.True(t, isConsumer)
}

func TestIsProducer(t *testing.T) {
	t.Parallel()

	address := NewAddress()
	address.SetProducer()

	isProducer := address.IsProducer()

	assert.True(t, isProducer)
}
