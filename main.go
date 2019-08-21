package main

import (
	"./go_routines/race_condition/mutual_exclusion"
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
	mutual_exclusion.ExampleOne()
	fmt.Println("//Main")
}
