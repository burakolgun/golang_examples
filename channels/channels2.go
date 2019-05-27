package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	fmt.Println("main run")
	go write(channel)
	fmt.Println("waiting")
	fmt.Println(<-channel)

}

func write(a chan string) {
	fmt.Println("in write")
	time.Sleep(3 * time.Second)

	a <- "completed"
}
