package go_routines

import (
	"fmt"
	"sync"
)

func main() {
	println("Hello, World!")

	ps := NewPubSub[int]()

	wg := sync.WaitGroup{}

	s1 := ps.Subscribe()
	s2 := ps.Subscribe()
	s3 := ps.Subscribe()

	subscribers := []<-chan int{s1, s2, s3}

	for i, s := range subscribers {
		wg.Add(1)
		i := i
		s := s
		go func() {
			for {
				select {
				case val, ok := <-s:
					if !ok {
						fmt.Printf("s%d closed\n", i)
						wg.Done()
						return
					}

					fmt.Printf("s%d: %d\n", i, val)
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		ps.Publish(i)
	}

	ps.Close()
	wg.Wait()

	fmt.Println("Done")
}

type PubSub[T interface{}] struct {
	subscribers []chan T
	mu          sync.RWMutex
	closed      bool
}

func NewPubSub[T interface{}]() *PubSub[T] {
	return &PubSub[T]{
		mu: sync.RWMutex{},
	}
}

func (ps *PubSub[T]) Subscribe() <-chan T {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan T)

	ps.subscribers = append(ps.subscribers, ch)

	return ch
}

func (ps *PubSub[T]) Publish(msg T) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subscribers {
		ch <- msg
	}
}

func (ps *PubSub[T]) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subscribers {
		close(ch)
	}

	ps.closed = true
}
