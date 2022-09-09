package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// terminate cancel all routine
// http://www.inanzzz.com/index.php/post/ry5e/cancelling-all-running-goroutines-with-context-withcancel-if-one-returns-an-error-within-golang

func main() {
	// The WaitGroup lets the main goroutine wait for all other goroutines
	// to terminate. However, this is no implicit in Go. The WaitGroup must
	// be explicitely incremented prior to the execution of any goroutine
	// (i.e. before the `go` keyword) and it must be decremented by calling
	// wg.Done() at the end of every goroutine (typically via the `defer` keyword).
	wg := sync.WaitGroup{}

	// The stop channel is an unbuffered channel that is closed when the main
	// thread wants all other goroutines to terminate (there is no way to
	// interrupt another goroutine in Go). Each goroutine must multiplex its
	// work with the stop channel to guarantee liveness.
	stopCh := make(chan struct{})

	for i := 0; i < 10; i++ {

		// It is important that the WaitGroup is incremented before we start
		// the goroutine (and not within the goroutine) because the scheduler
		// makes no guarantee that the goroutine starts execution prior to
		// the main goroutine calling wg.Wait().
		wg.Add(1)
		go func(i int, stopCh chan struct{}) {
			time.Sleep(5 * time.Second)
			//fmt.Println("print go routine ")
			// defer keyword guarantees that the WaitGroup count is
			// decremented when the goroutine exits.
			defer wg.Done()

			log.Printf("started goroutine %d", i)

			select {
			// Since we never send empty structs on this channel we can
			// take the return of a receive on the channel to mean that the
			// channel has been closed (recall that receive never blocks on
			// closed channels).
			case <-stopCh:
				log.Printf("stopped goroutine %d", i)
			}
		}(i, stopCh)
	}

	fmt.Println("block code")
	close(stopCh)
	log.Printf("stopping goroutines")
	wg.Wait()
	log.Printf("all goroutines stopped")
}
