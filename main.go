package main

import (
	"./go_routines/race_condition"
	"fmt"
)

func main() {
	//embeddedStatements.Embedded(true)
	//go_routines.ExampleOne()
	//go_routines.ExampleTwo()
	//go_routines.ExampleThree()
	//go_routines.ExampleFour()
	race_condition.ExampleOne()
	fmt.Println("//Main")
}
