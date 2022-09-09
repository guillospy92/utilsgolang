package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

func main() {
	generator := func(ctx context.Context, chError chan error) <-chan int {
		ch := make(chan int)

		n := 1

		go func() {
			defer close(ch)
			for {
				select {
				case ch <- n:
				case <-ctx.Done():
					fmt.Println("cancel generate 1")
					return
				}
				n++
			}
		}()

		return ch
	}

	generator2 := func(ctx context.Context, chError chan error) <-chan int {
		ch := make(chan int)

		n := 1

		go func() {
			defer close(ch)
			for {
				select {
				case ch <- n:
				case <-ctx.Done():
					fmt.Println("cancel generate 2")
					return
				}
				if n == 6 {
					chError <- errors.New("error propagated")
				}
				n++
			}
		}()

		return ch
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chError := make(chan error)

	go func() {
		for n := range generator(ctx, chError) {
			fmt.Println("generator 1", n)
		}
	}()

	go func() {
		for nn := range generator2(ctx, chError) {
			fmt.Println("generator 2", nn)
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()

	data, ok := <-chError

	if ok {
		fmt.Println("error")
		cancel()
		return
	}

	fmt.Println("data", data)

	time.Sleep(1 * time.Second)

}
