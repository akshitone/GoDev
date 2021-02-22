package main

import (
	"fmt"
)

func main() {
	name := ""
	fmt.Print("Enter the name you want to send the mail: ")
	fmt.Scan(&name)
	emailMe(name)
}

func emailMe(name string) {
	str := `Heyy my name is ` + name + `.`
	fmt.Println(str)
}
