package main

import (
	"fmt"
	"sync"
)

func generatorOutFan(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func squareOutFan(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {

		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch := generatorOutFan(2, 3, 4, 5, 6, 7)

	ch1 := squareOutFan(ch)
	ch2 := squareOutFan(ch)

	for l := range merge(ch1, ch2) {
		fmt.Println(l, "value total")
	}
}
