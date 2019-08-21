package go_routines

import (
	"fmt"
	"runtime"
	"sync"
)

//wg is used for wait for the program to finish
var wg sync.WaitGroup

func ExampleThree() {
	//Allocate one logical processors for scheduler to use

	runtime.GOMAXPROCS(1)

	//Add a count of two, one for each goroutine
	wg.Add(2)

	//Create two goroutines
	fmt.Println("Create goroutines")

	go printPrime("A")
	go printPrime("B")

	// Wait for goroutines to finish.
	fmt.Println("Waiting To Finish...")
	wg.Wait()
	fmt.Println("Terminating Program")
}

func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
