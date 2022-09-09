package main

import (
	"fmt"
	"golang.org/x/net/context"
)

type database map[string]bool
type userIDKEY string

var db = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userID string) {
	ctx = context.WithValue(ctx, userIDKEY("newKey"), userID)
	ctx = context.WithValue(ctx, userIDKEY("userIDKEY"), userID)

	ch := checkMemberShip(ctx)
	status := <-ch
	fmt.Printf("membership status of userID : %s : %v\n", userID, status)
}

func checkMemberShip(newCtx context.Context) <-chan bool {
	ch := make(chan bool)
	userID := newCtx.Value(userIDKEY("userIDKEY")).(string)
	newKey := newCtx.Value(userIDKEY("newKey")).(string)
	fmt.Println(newKey)
	go func() {
		defer close(ch)
		status := db[userID]
		ch <- status
	}()

	return ch
}
