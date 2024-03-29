package main

import (
	"log"
	"os"
	"text/template"
)

type car struct {
	Name  string
	Model int
}

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

	audi := car{
		Name:  "Audi",
		Model: 4,
	}

	renault := car{
		Name:  "Renault",
		Model: 3,
	}

	sages := []sage{buddha, gandhi, mlk}
	cars := []car{audi, renault}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatal(err)
	}
}
