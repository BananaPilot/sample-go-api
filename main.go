package main

func main() {
	conn := newDbConn(Connection(CreateUri("root", "1234", "localhost", "3306", "mystuff")))
	server := NewApiServer(":8080", conn)
	server.run()
}
