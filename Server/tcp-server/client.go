package main

import (
	"fmt"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Fprintln(connection, "I dialed you!")
}
