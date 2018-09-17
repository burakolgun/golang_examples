package main

import (
	"log"
	"os"
	"text/template"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template.gohtml"))
}

func main () {
	err := tpl.ExecuteTemplate(os.Stdout, "template.gohtml", `Release self-focus; embrace other-focus`)

	if err != nil {
		log.Fatal(err)
	}
}