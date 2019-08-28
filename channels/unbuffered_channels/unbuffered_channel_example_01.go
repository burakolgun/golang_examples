package unbuffered_channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}
func ExampleOne() {
	// Create an unbuffered channel
	court := make(chan int)

	wg.Add(2)

	// Launch two player
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set.
	court <- 1

	// Wait for the game to finish
	wg.Wait()
}

// Player simulates a person playing the game of tennis
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for  {
		// Wait for the ball to be hit back to us

		ball, ok := <-court

		if !ok {
			// if the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n ", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}
