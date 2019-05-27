package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)

	go func() {
		fmt.Println("on channel")
		time.Sleep(3 * time.Second)
		channel <- "channel data"
	}()

	fmt.Println("here1")

	b := <-channel

	fmt.Println(b)
}
