package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("composition-tpl/tpl-composite.gohtml"))
}

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{"MCA-101", "Computer Science", "4"},
				{"MCA-102", "Computer Management", "5"},
				{"MCA-103", "Engineering", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{"MCA-201", "Computer Science", "4"},
				{"MCA-202", "Computer Management", "5"},
				{"MCA-203", "Engineering", "3"},
			},
		},
		Summer: semester{},
	}
	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}
