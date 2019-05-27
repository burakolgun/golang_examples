package main

import (
	"fmt"
)

func main() {
	var variableTypesForInt string = "int, int8, int16, int32, int64"
	var variableTypesForUint string = "uint, uint8, uint16, uint32, uint64"
	var boolTrue = true
	var boolFalse = false
	var float float64 = 12.442131

	fmt.Println(variableTypesForInt)
	fmt.Println(variableTypesForUint)
	fmt.Println("boolean", boolFalse, boolTrue)
	fmt.Println("float", float)

}
