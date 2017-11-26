package main

import ("fmt")

func factorial (num int) int {
	if num == 0 {
		return 1
	}
	fmt.Println("num", num)
	return num * factorial(num -1)
}

func main() {
	fmt.Println("5 factarial = ", factorial(5))
}