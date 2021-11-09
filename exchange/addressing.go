package exchange

const (
	AddressConsumer = "consumer"
	AddressProducer = "producer"
)

type Address struct {
	kind string
}

func NewAddress() *Address {
	return &Address{
		kind: AddressProducer,
	}
}

func (addr *Address) SetConsumer() {
	addr.kind = AddressConsumer
}

func (addr *Address) SetProducer() {
	addr.kind = AddressProducer
}

func (addr *Address) IsConsumer() bool {
	return addr.kind == AddressConsumer
}

func (addr *Address) IsProducer() bool {
	return addr.kind == AddressProducer
}
