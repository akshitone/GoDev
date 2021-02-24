package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("nesting-tpl/templates/*.gohtml"))
}

type person struct {
	Name    string
	Age     int
	Married bool
}

func main() {
	akshit := person{"akshit", 21, true}
	viral := person{"", 22, true}
	rajan := person{"rajan", 22, false}
	users := []person{akshit, viral, rajan}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
