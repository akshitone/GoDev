package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
		go handle(connection)
	}
}
func handle(connection net.Conn) {
	// 10 second timer start
	err := connection.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection timeout!")
	}
	// 10 second timer end

	scan := bufio.NewScanner(connection)
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
		fmt.Fprintf(connection, "Akshit Mithaiwala: %s\n", line)
	}

	defer connection.Close()

	fmt.Println("Code got here!")
}
