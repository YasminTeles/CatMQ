// nolint: wsl
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

type Server struct {
	port       string
	connection net.Conn
	listener   net.Listener
	exch       *exchange.Exchange
}

func NewServer(port string) *Server {
	return &Server{
		port:       port,
		connection: nil,
		listener:   nil,
		exch:       exchange.NewExchange(),
	}
}

func NewServerDefault() *Server {
	return NewServer(ServerPort)
}

func (server *Server) ListenAndServe() {
	log.Println("Starting server...")

	address := server.getAddress()

	var err error

	server.listener, err = net.Listen(ServerProtocol, address)
	if err != nil {
		log.Panicf("Some listener error: %s.\n", err)
	}
	defer server.listener.Close()

	log.Printf("Listening on %s.\n", address)

	var connection net.Conn
	for {
		connection, err = server.listener.Accept()
		if err != nil {
			log.Panicf("Some listener connection error: %s.\n", err)
		}

		go server.handleConnection(connection)
	}
}

func (server *Server) getAddress() string {
	return net.JoinHostPort(ServerHost, server.port)
}

func (server *Server) handleConnection(connection net.Conn) {
	log.Printf("Client connected from %s.\n", connection.RemoteAddr().String())

	addr := exchange.NewAddress()

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		request := scanner.Text()
		if request == "" {
			continue
		}

		log.Printf("Message received: %s\n", request)

		response := exchange.Route(request, server.exch, addr)

		log.Printf("Message send: %s\n", response)
		fmt.Fprintln(connection, response)
	}
}

func (server *Server) Close() {
	log.Println("Shutting down server...")

	// if err := server.connection.Close(); err != nil {
	// 	log.Panicf("Some disconnection error: %s.\n", err)
	// }

	if err := server.listener.Close(); err != nil {
		log.Panicf("Some listener error: %s.\n", err)
	}
}
