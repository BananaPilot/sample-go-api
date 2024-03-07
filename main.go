package main

func main() {
	server := NewApiServer(":8080", newDbConn(Connection(CreateUri("root", "1234", "localhost", "3306", "mystuff"))))
	server.run()
}
