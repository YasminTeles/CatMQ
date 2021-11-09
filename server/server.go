// nolint: gochecknoglobals, wsl
package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/YasminTeles/CatMQ/exchange"
)

const (
	ServerHost     = "127.0.0.1"
	ServerPort     = "23023"
	ServerProtocol = "tcp"
)

var Connection net.Conn

func ListenAndServe() {
	log.Println("Starting server...")

	address := GetAddress()

	listener, _ := net.Listen(ServerProtocol, address)
	defer listener.Close()

	log.Printf("Listening on %s.\n", address)

	exch := exchange.NewExchange()
	var err error

	for {
		Connection, err = listener.Accept()
		if err != nil {
			log.Panicf("Some connection error: %s.\n", err)
		}

		go handleConnection(Connection, exch)
	}
}

func GetAddress() string {
	return fmt.Sprintf("%s:%s", ServerHost, ServerPort)
}

func handleConnection(connection net.Conn, exch *exchange.Exchange) {
	log.Printf("Client connected from %s.\n", connection.RemoteAddr().String())

	addr := exchange.NewAddress()

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		request := scanner.Text()
		if request == "" {
			continue
		}

		log.Printf("Message received: %s\n", request)

		response := exchange.Route(request, exch, addr)

		log.Printf("Message send: %s\n", response)
		fmt.Fprintln(connection, response)
	}
}

func Close() {
	if err := Connection.Close(); err != nil {
		log.Panicf("Some disconnection error: %s.\n", err)
	}
}
