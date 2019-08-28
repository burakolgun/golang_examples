package main

import (
	"./channels/unbuffered_channels"
	"./channels/buffered_channels"
	"fmt"
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
	unbuffered_channels.ExampleOne()
	buffered_channels.ExampleOne()
	fmt.Println("//Main")
}
