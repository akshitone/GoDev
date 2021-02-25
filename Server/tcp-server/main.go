package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	serverListen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer serverListen.Close()

	for {
		connection, err := serverListen.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(connection, "Hello from TCP server, I'm GO Program\n")
		fmt.Fprintln(connection, "How is your day?")
		fmt.Fprintf(connection, "%v", "Well, i hope!")

		connection.Close()
	}
}
