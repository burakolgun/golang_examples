package main

import (
	"html/template"
	"log"
	"os"
)

var tmp *template.Template

func init() {
	tmp = template.Must(template.ParseGlob("*"))
}
func main() {
	err := tmp.Execute(os.Stdout, "title here")
	if err != nil {
		log.Fatal(err)
	}
}
