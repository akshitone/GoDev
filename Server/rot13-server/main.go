package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	serverListen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer serverListen.Close()

	for {
		connection, err := serverListen.Accept()
		if err != nil {
			panic(err)
		}
		go handle(connection)
	}
}

func handle(connection net.Conn) {
	scan := bufio.NewScanner(connection)
	for scan.Scan() {
		line := strings.ToLower(scan.Text())
		byteString := []byte(line)
		rot := rotate13(byteString)

		fmt.Fprintf(connection, "%s - %s\n\n", line, rot)
	}
}

func rotate13(byteString []byte) []byte {
	var rotate13 = make([]byte, len(byteString))
	for index, value := range byteString {
		if value <= 109 {
			rotate13[index] = value + 13
		} else {
			rotate13[index] = value - 13
		}
	}
	return rotate13
}
