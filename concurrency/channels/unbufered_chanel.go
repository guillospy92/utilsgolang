package main

import (
	"sync"
	"time"
)

// unbufered channel esperan que un emetinte como un receptor esten listo

func main() {
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		c <- `foo`
		c <- `bar`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		println(`Message: ` + <-c)
		println(`Message: ` + <-c)
	}()

	wg.Wait()
}
