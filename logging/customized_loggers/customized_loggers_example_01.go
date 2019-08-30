package customized_loggers

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical Problem
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 777)

	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	// Discard -> DevNull
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)

	Info = log.New(os.Stdout, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)

	Warning = log.New(os.Stdout, "Warning: ", log.Ldate | log.Ltime | log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate | log.Ltime | log.Lshortfile)

}

func ExampleOne()  {
	Trace.Println("I have something standard to say")
	Info.Println("Special information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed.")
}
