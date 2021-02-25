package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("review/data-tpl/template-struct.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func main() {
	akshit := person{"akshit", 21}
	viral := person{"viral", 22}
	users := []person{akshit, viral}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
