// nolint: gochecknoglobals,wrapcheck
package client

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/YasminTeles/CatMQ/message"
	"github.com/YasminTeles/CatMQ/server"
)

var Connection net.Conn

func Connect() bool {
	var err error

	address := server.GetAddress()
	log.Printf("Connecting to %s...\n", address)

	Connection, err = net.Dial(server.PROTOCOL, address)
	if err != nil {
		log.Printf("Some connection error: %s.\n", err)
	}

	return !(err != nil)
}

func Disconnect() bool {
	log.Println("Disconnecting to server...")

	if Connection != nil {
		if err := Connection.Close(); err != nil {
			log.Printf("Some disconnection error: %s.\n", err)

			return false
		}
	}

	return true
}

func Publish(data string) bool {
	request := message.NewPutMessage(data)

	pack, err := send(request.ToPack())
	if err != nil {
		return false
	}

	response := message.NewMessage(pack)

	return response.Operation == message.OperationOK
}

func send(pack string) (string, error) {
	fmt.Fprintln(Connection, pack)

	return bufio.NewReader(Connection).ReadString('\n')
}

func Get() string {
	request := message.NewGetMessage()

	pack, err := send(request.ToPack())
	if err != nil {
		return message.MessageError
	}

	response := message.NewMessage(pack)

	return response.Data
}
