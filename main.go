package main

import (
	"./channels/unbuffered_channels"
	"./channels/buffered_channels"
	"fmt"
	//"./logging" // because log.fatal
	"./logging/customized_loggers"
	"./decoding_encoding/json"
)

func main() {
	//embeddedStatements.Embedded(true)
	//go_routines.ExampleOne()
	//go_routines.ExampleTwo()
	//go_routines.ExampleThree()
	//go_routines.ExampleFour()
	//race_condition.ExampleOne()
	//atomic_functions.ExampleOne()
	//atomic_functions.ExampleTwo()
	//mutual_exclusion.ExampleOne()
	//channels.ExampleOne()
	//channels.ExampleTwo()
	//channels.ExampleThree()
	//channels.ExampleFour()
	//conditions.IfElseExample()
	//_defer.Example()
	customized_loggers.Info.Println("unbuffered_channels.ExampleOne()")
	unbuffered_channels.ExampleOne()
	customized_loggers.Info.Println("buffered_channels.ExampleOne()")
	buffered_channels.ExampleOne()
	//customized_loggers.Info.Println("logging.ExampleOne()")
	//logging.ExampleOne()
	customized_loggers.Info.Println("customized_loggers.ExampleOne()")
	customized_loggers.ExampleOne()
	customized_loggers.Info.Println("json.ExampleTwo()")
	json.ExampleTwo()
	customized_loggers.Info.Println("json.ExampleThree()")
	json.ExampleThree()
	fmt.Println("//Main")
}
