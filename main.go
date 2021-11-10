package main

import "github.com/YasminTeles/CatMQ/server"

func main() {
	server := server.NewServerDefault()
	server.ListenAndServe()
}
