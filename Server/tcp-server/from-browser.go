package main

import (
	"bufio"
	"fmt"
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

		go handle(connection)
	}

}

func handle(connection net.Conn) {
	scan := bufio.NewScanner(connection)
	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
	}
	defer connection.Close()

	fmt.Println("Code got here!")
}
