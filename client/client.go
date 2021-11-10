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

type Client struct {
	connection net.Conn
}

func NewClient() *Client {
	return &Client{
		connection: nil,
	}
}

func (cli *Client) Connect() {
	address := server.GetAddress()
	log.Printf("Connecting to server %s...\n", address)

	connection, err := net.Dial(server.ServerProtocol, address)
	if err != nil {
		log.Printf("Some connection error: %s.\n", err)
	}

	cli.connection = connection
}

func (cli *Client) Disconnect() error {
	log.Println("Disconnecting to server...")

	if cli.connection != nil {
		if err := cli.connection.Close(); err != nil {
			log.Printf("Some disconnection error: %s.\n", err)
			return err
		}
	}

	return nil
}

func (cli *Client) Publish(data string) bool {
	request := message.NewPutMessage(data)

	pack, err := cli.send(request.ToPack())
	if err != nil {
		return false
	}

	response := message.NewMessage(pack)

	return response.Operation == message.OperationOK
}

func (cli *Client) send(pack string) (string, error) {
	fmt.Fprintln(cli.connection, pack)

	return bufio.NewReader(cli.connection).ReadString('\n')
}

func (cli *Client) Get() string {
	request := message.NewGetMessage()

	pack, err := cli.send(request.ToPack())
	if err != nil {
		return message.MessageError
	}

	response := message.NewMessage(pack)

	return response.Data
}

func (cli *Client) Consumer() bool {
	request := message.NewConsumerMessage()

	pack, err := cli.send(request.ToPack())
	if err != nil {
		return false
	}

	response := message.NewMessage(pack)

	return response.Operation == message.OperationOK
}

func (cli *Client) Producer() bool {
	request := message.NewProducerMessage()

	pack, err := cli.send(request.ToPack())
	if err != nil {
		return false
	}

	response := message.NewMessage(pack)

	return response.Operation == message.OperationOK
}
