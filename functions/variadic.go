package main

import "fmt"

func variadicFunc(variables ...int) int {
	total := 0

	for _, i := range variables {
		total += i
	}

	return total
}

func main() {
	fmt.Println(variadicFunc(10,20,30,40))

}