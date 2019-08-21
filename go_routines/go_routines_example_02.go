package go_routines

import (
	"fmt"
	"runtime"
	"sync"
)

func ExampleTwo() {
	//Allocate 1 logical proc for the scheduler to use
	runtime.GOMAXPROCS(1)

	//wh is used to wait for the program to finish
	// Add a count two, one for each goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	//Declare an anonymous function and create a goroutine
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
		fmt.Println()
	}()

	// Wait for the goroutines to finish
	wg.Wait()

	fmt.Println("\n Terminating program")
}
