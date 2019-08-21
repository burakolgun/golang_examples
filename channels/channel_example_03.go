package channels

import (
	"fmt"
	"time"
)

func ExampleThree() {
	a := make(chan string)
	b := make(chan string)

	fmt.Println("start")

	go getA(a)
	go getBAfterA(a, b)
	<-b
	fmt.Println("b , a ok")
}

func getA(a chan string) {
	fmt.Println("in a")
	time.Sleep(3 * time.Second)
	a <- "A ok"
}

func getBAfterA(a, b chan string) {
	fmt.Println("in b")
	<-a
	time.Sleep(3 * time.Second)
	b <- "B ok"
}
