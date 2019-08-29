package logging

import (
	"log"
)

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func ExampleOne() {
	log.Println("println message")

	log.Fatalln("Fatalln message")

	log.Panicln("Panic message")
}
