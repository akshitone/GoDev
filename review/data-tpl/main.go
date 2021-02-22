package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("review/data-tpl/template-slice.gohtml"))
}

func main() {
	names := []string{"Akshit", "Viral", "Rajan", "Tushar"}

	err := tpl.ExecuteTemplate(os.Stdout, "template-slice.gohtml", names)
	if err != nil {
		log.Fatalln(err)
	}
}
