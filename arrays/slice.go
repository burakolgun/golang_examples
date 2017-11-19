package main

import (
	"fmt"
)

func main() {
	var array = make([]int, 3)
	var copyArray = make([]int ,5)

	array[1] = 1
	array[0] = 2
	array[2] = 99
	fmt.Println(array)

	array = append(array,1,2,3,4,5,6)
	copy(copyArray,array)
	fmt.Println(array)
	fmt.Println(copyArray)

	copyArray = make([]int, len(array))
	copy(copyArray,array)
	fmt.Println(copyArray)
	


}