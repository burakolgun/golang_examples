package functions

import "fmt"

func variadicFunc(variables ...int) int {
	total := 0

	for _, i := range variables {
		total += i
	}

	return total
}

func ExampleVariadic() {
	fmt.Println(variadicFunc(10, 20, 30, 40))

}
