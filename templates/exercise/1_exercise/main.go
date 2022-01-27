package main

import (
	"log"
	"os"
	"text/template"
)

type Course struct {
	Number string
	Name   string
	Units  string
}

type Semester struct {
	Term    string
	Courses []Course
}

type Year struct {
	AcaYear string
	Fall    Semester
	Spring  Semester
	Summer  Semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	Years := []Year{
		Year{
			AcaYear: "2020-2021",
			Fall: Semester{
				Term: "Fall",
				Courses: []Course{
					Course{"CSCI-40", "Introduction to Programming in Go", "4"},
					Course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					Course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: Semester{
				Term: "Spring",
				Courses: []Course{
					Course{"CSCI-50", "Advanced Go", "5"},
					Course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					Course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
		Year{
			AcaYear: "2021-2022",
			Fall: Semester{
				Term: "Fall",
				Courses: []Course{
					Course{"CSCI-40", "Introduction to Programming in Go", "4"},
					Course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
					Course{"CSCI-140", "Mobile Apps Using Go", "4"},
				},
			},
			Spring: Semester{
				Term: "Spring",
				Courses: []Course{
					Course{"CSCI-50", "Advanced Go", "5"},
					Course{"CSCI-190", "Advanced Web Programming with Go", "5"},
					Course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, Years)
	if err != nil {
		log.Fatalln(err)
	}
}
