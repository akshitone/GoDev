package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Listener error:", err)
	}

	defer listen.Close()

	for {
		connection, err := listen.Accept()
		if err != nil {
			log.Fatalln("Connection error:", err)
			continue
		}
		io.WriteString(connection, "I see you connected")
		connection.Close()
	}

}
