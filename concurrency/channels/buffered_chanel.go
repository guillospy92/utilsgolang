package main

import "fmt"

func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)
		for i := 0; i <= 10; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
