package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	var completed bool

	go func() {
		for i := 0; i < 3; i++ {
			ch <- "message"
		}
		defer close(ch)
	}()

	for !completed {
		select {
		case m, ok := <-ch:
			if !ok {
				completed = true
			}
			fmt.Println(ok)
			fmt.Println(m)
		}
	}

	fmt.Println("bye")

	/*	for i := 0; i < 3; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no data")
		}

		fmt.Println("processing")
	}*/
}
