package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funMap).ParseFiles("review/date-tpl/tpl-date.gohtml"))
}

var funMap = template.FuncMap{
	"funDateMDY": funDate,
}

func funDate(t time.Time) string {
	return t.Format(time.RFC850)
	// return t.Format("Monday, 01-Feb-06 15:04:05 MST")
}

func main() {
	fmt.Println(time.Now())
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-date.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
