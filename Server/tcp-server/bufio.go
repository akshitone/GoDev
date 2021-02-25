package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	song := "It's been a long day \nWithout you my friend!"

	scan := bufio.NewScanner(strings.NewReader(song))
	// scan.Split(bufio.ScanRunes)

	for scan.Scan() {
		line := scan.Text()
		fmt.Println(line)
	}
}
