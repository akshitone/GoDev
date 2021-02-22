package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("review/data-tpl/template-map.gohtml"))
}

func main() {
	users := map[string]int{"akshit": 21, "viral": 22, "rajan": 23, "tushar": 22}

	err := tpl.ExecuteTemplate(os.Stdout, "template-map.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
