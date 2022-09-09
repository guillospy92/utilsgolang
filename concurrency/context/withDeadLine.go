package main

import (
	"fmt"
	"golang.org/x/net/context"
	"time"
)

type data struct {
	result string
}

func main() {
	deadLine := time.Now().Add(4 * time.Second)
	fmt.Println(deadLine)
	ctx, cancel := context.WithDeadline(context.Background(), deadLine)
	defer cancel()

	compute := func() <-chan data {
		ch := make(chan data)

		go func() {
			defer close(ch)
			time.Sleep(4 * time.Second)

			deadLine, ok := ctx.Deadline()
			fmt.Println(deadLine, "value dead line")
			fmt.Println(time.Now().Add(1*time.Second), "value sleep routine")
			fmt.Println("result time compare", deadLine.Sub(time.Now().Add(1*time.Second)).Seconds())
			if ok {
				if deadLine.Sub(time.Now().Add(1*time.Second)) < 0 {
					fmt.Println("not sufficient time given")
					return
				}
			}

			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				fmt.Println("terminate")
				return
			}

		}()

		return ch
	}

	chanel, ok := <-compute()

	if ok {
		fmt.Println("channel works data", chanel)
		return
	}

	fmt.Println("work is completed")
}
