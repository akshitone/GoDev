package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl, err := template.ParseFiles("review/multi-parse/docs/one.gmao", "review/multi-parse/two.gmao")
	// tpl, err := template.ParseGlob("review/multi-parse/docs/*.gmao")
	tpl, err := template.ParseGlob("review/multi-parse/docs/*")
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
