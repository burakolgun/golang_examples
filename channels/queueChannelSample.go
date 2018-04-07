package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"
)

const (
	queueSize = 10
	numberOfReaders = 3
)

var queue = make(chan int, queueSize)
var done = make(chan bool)

func main() {
	fmt.Printf("Queue Size %d \n", queueSize)
	fmt.Printf("Number of Readers %d \n", numberOfReaders)
	defer fmt.Println("Channel Sample Finished")

	//creating  readers
	for i := 1; i <= numberOfReaders; i++ {
		go reader(i)
	}

	//creating  producer
	go producer()

	<-done
}

func producer() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please enter the queue message limit")
		fmt.Println("If you want to quit, just write the  \"quit\" ")

		//multiple return
		bytes, _, err := reader.ReadLine()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		quitCheck := string(bytes)

		if quitCheck == "quit" {
			done <- true
			return
		}

		//multiple return
		limit, err := strconv.Atoi(quitCheck)

		if err != nil {
			fmt.Println("Just write the number")
			continue
		}

		for i := 0; i < limit; i++ {
			// Messages sending to queue
			queue <- i
		}
	}
}

func reader(number int) {
	for {
		message := <-queue
		fmt.Printf("Im Reader #%d --- Message reading in queue --- %d \n", number, message)

		time.Sleep(1 * time.Second)
	}
}
