package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	queueSize       int = 100
	numberOfReaders int = 3
)

var queue = make(chan string, queueSize)
var done = make(chan bool)
var delay time.Duration = 15

func main() {
	defer fmt.Println("Channel Sample Finished")

	fmt.Printf("Queue Size %d \n", queueSize)

	fmt.Printf("Number of Readers %d \n", numberOfReaders)

	//creating  queue reader(s)
	for i := 1; i <= numberOfReaders; i++ {
		delay = 10
		fmt.Printf("Please gime me a special delay time for reader #%d initial delay time -> %d \n", i, delay)

		_, err := fmt.Scanf("%d \n", &delay)

		if err != nil {
			fmt.Println(err)
			delay = 10
		}

		go queueReader(i, delay)
	}

	//creating a producer
	go producer()

	<-done
}

func producer() {
	for {
		fmt.Println("Please enter the queue message limit")
		fmt.Println("If you want to quit, just write the  \"quit\" ")

		//multiple return
		bytes, _, err := bufio.NewReader(os.Stdin).ReadLine()

		if err != nil {
			fmt.Println(err.Error())

			continue
		}

		quitCheck := string(bytes)

		if quitCheck == "quit" {
			done <- true
		}

		//multiple return
		limit, err := strconv.Atoi(quitCheck)

		if err != nil {
			fmt.Println("Just write the number")

			continue
		}

		for i := 0; i < limit; i++ {
			fmt.Println("Please give the message for queue")

			queueMessageByte, _, err := bufio.NewReader(os.Stdin).ReadLine()

			if err != nil {
				fmt.Println(err)

				continue
			}

			queueMessage := string(queueMessageByte)
			// Messages sending to queue
			queue <- queueMessage
		}
	}
}

func queueReader(queueName int, delayTime time.Duration) {
	for {
		message := <-queue
		fmt.Printf("Im Reader #%d --- Message reading in queue --- %s", queueName, message)
		fmt.Printf("delay time -> %d", delayTime)

		time.Sleep(delayTime * time.Second)
	}
}
