package main

import (
	"APItry/api"
	"APItry/storage"
)

func main() {
	conn := api.NewDbConn(storage.Connection(storage.CreateUri("root", "1234", "localhost", "3306", "mystuff")))
	server := api.NewApiServer(":8080", conn)
	server.Run()
}
