package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := []string{"Ghandi", "MLK", "buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)

	if err != nil {
		log.Fatal(err)
	}
}
