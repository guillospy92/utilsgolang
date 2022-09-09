package main

import (
	"fmt"
	"sync"
	"time"
)

type Record struct {
	sync.Mutex
	data string
}

func main() {
	rec := &Record{}
	var wg sync.WaitGroup
	wg.Add(1)

	go func(rec *Record) {
		defer wg.Done()
		for {
			fmt.Println("1")
			rec.Lock()
			if rec.data != "" {
				fmt.Println("data", rec.data)
				rec.Unlock()
				return
			}
			rec.Unlock()
		}

	}(rec)

	time.Sleep(2 * time.Second)

	rec.Lock()
	rec.data = "gopher"
	rec.Unlock()

	wg.Wait()
}
