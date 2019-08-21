package channels

import (
	"fmt"
	"time"
)

var chan1 = make(chan bool, 1)
var chan2 = make(chan bool, 1)

func ExampleFour() {
	fmt.Println("Process Starting...")
	go worker(5)
	go timer(6)

	select {
	case <-chan1:
		fmt.Println("Mission Completed")
	case <-chan2:
		fmt.Println("Time Out")
	}
}

func timer(second time.Duration) {
	time.Sleep(second * time.Second)
	chan2 <- true
}

func worker(second time.Duration) {
	fmt.Println("Process Started")
	time.Sleep(second * time.Second)
	chan1 <- true
}
