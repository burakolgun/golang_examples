package functions

import (
	"fmt"
)

func multipleOperation(number1, number2 int) (int, int, int, int) {
	sum := number1 + number2
	sub := number1 - number2
	div := number1 / number2
	mul := number1 * number2

	return sum, sub, mul, div
}

func ExampleMultipleReturn() {
	sum, sub, div, mul := multipleOperation(30, 12)
	fmt.Println(sum, sub, div, mul)

}
