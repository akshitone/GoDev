package main

import (
	"bufio"
	"fmt"
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

		scan := bufio.NewScanner(connection)
		for scan.Scan() {
			line := scan.Text()
			fmt.Println(line)
			if line == "" {
				fmt.Println("End of the http request")
				break
			}
		}

		defer connection.Close()

		fmt.Println("Code got here!")
		io.WriteString(connection, "I see you connected")
		connection.Close()
	}

}
