package main

import (
	"golang.org/x/net/context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func(ctx2 context.Context) {
		ctx2, cancel := context.WithTimeout(ctx2, 1*time.Second)
		defer cancel()

		func(ctx3 context.Context) (string, error) {
			select {
			case <-ctx.Done():
				return "", ctx3.Err()
			}
		}(ctx2)
	}(ctx)

}
