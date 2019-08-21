package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no belief",
	}

	gandhi := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	sages := []sage{buddha, gandhi, mlk}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatal(err)
	}
}
