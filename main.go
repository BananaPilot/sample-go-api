package main

import (
	"github.com/BananaPilot/sample-api/api"
	"github.com/BananaPilot/sample-api/storage"
)

func main() {
	conn := api.NewDbConn(storage.Connection(storage.CreateUri("root", "1234", "localhost", "3306", "mystuff")))
	server := api.NewApiServer(":8080", conn)
	server.Run()
}
