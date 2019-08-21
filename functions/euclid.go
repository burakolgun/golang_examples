package functions

import (
	"fmt"
)

func euclid(x, y int) int {
	if y == 0 {
		return x
	}

	return euclid(y, x%y)
}

func ExampleEuclid() {
	fmt.Println("euclid for 48 and 30 =", euclid(48, 30))
}
