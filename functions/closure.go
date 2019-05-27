package main

import (
	"fmt"
)

func count() (sum func() int, counter func() int) {
	_count := 0
	total := 0
	sum = func() int {
		_count++
		total = total + _count
		return total
	}

	counter = func() int {
		return _count
	}

	return sum, counter
}

func main() {
	sum, counter := count()

	for counter() <= 50 { //1
		fmt.Println("total = ", sum()) //2
		fmt.Println("Count = ", counter())
	}
}
