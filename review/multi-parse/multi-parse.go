package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("review/multi-parse/one.gmao", "review/multi-parse/two.gmao")
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
