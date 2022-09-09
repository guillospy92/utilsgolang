package main

import (
	"fmt"
	"sync"
	"time"
)

type RecordCondition struct {
	sync.Mutex
	data string
	cond *sync.Cond
}

func NewRecordCondition() *RecordCondition {
	r := RecordCondition{}
	r.cond = sync.NewCond(&r)
	return &r
}

func main() {
	var wg sync.WaitGroup
	rec := NewRecordCondition()
	wg.Add(1)
	go func(re *RecordCondition) {
		defer wg.Done()
		re.Lock()
		re.cond.Wait()
		re.Unlock()
		fmt.Println("Data: ", re.data)
		return
	}(rec)

	time.Sleep(2 * time.Second)
	rec.Lock()
	rec.data = "gopher"
	rec.Unlock()

	rec.cond.Signal()

	wg.Wait()
}
