package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST     = ""
	PORT     = "23023"
	PROTOCOL = "tcp"
)

func ListenAndServe() {
	log.Println("Starting server...")

	address := getAddress()

	listener, _ := net.Listen(PROTOCOL, address)
	defer listener.Close()

	log.Printf("Listening on %s.\n", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Panicf("Some connection error: %s.\n", err)
		}

		go handleConnection(connection)
	}
}

func getAddress() string {
	return fmt.Sprintf("%s:%s", HOST, PORT)
}

func handleConnection(connection net.Conn) {
	log.Printf("Client connected from %s.\n", connection.RemoteAddr().String())

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		received := scanner.Text()
		log.Printf("Message received: %s\n", received)

		send := "Message received."

		log.Printf("Message send: %s\n", send)
		fmt.Fprintln(connection, send)
	}
}
