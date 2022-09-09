package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = new(sync.WaitGroup)
	a := make(chan int)
	b := make(chan int)
	isError := make(chan bool)

	wg.Add(4)

	// send channels
	go writeChannelA(a, wg, isError)
	go writeChannelB(b, wg)

	// receptors chanel
	go readChannelB(b, wg)
	go readChanelA(a, wg)

	fmt.Println(<-isError)

	go func() {
		wg.Wait()
	}()
	
	// block program if there is error
}

func writeChannelA(a chan<- int, wg *sync.WaitGroup, isError chan<- bool) {
	defer close(a)
	for i := 0; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		if i == 4 {
			isError <- true
			break
		}
		a <- i
	}
	defer wg.Done()
}

func readChanelA(a <-chan int, wg *sync.WaitGroup) {
	for value := range a {
		fmt.Println("receptor channel a", value)
	}
	defer wg.Done()
}

func writeChannelB(b chan<- int, wg *sync.WaitGroup) {
	defer close(b)
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		b <- i
	}
	defer wg.Done()
}

func readChannelB(b <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range b {
		fmt.Println("receptor channel b", value)
	}
}
