package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funMap).ParseFiles("review/func-data-tpl/tpl-func-data.gohtml"))
}

var funMap = template.FuncMap{
	"upperCase":      strings.ToUpper,
	"firstThreeChar": firstThree,
}

func firstThree(str string) string {
	str = strings.TrimSpace(str)
	str = str[:3]
	return str
}

type person struct {
	Name string
	Age  int
}

func main() {
	akshit := person{"akshit", 21}
	viral := person{"viral", 22}
	users := []person{akshit, viral}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl-func-data.gohtml", users)
	if err != nil {
		log.Fatalln(err)
	}
}
