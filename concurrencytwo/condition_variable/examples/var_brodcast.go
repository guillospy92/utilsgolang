package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type RecordBroadcast struct {
	sync.Mutex

	buf  string
	cond *sync.Cond

	writers []io.Writer
}

func NewRecordBroadcast(writers ...io.Writer) *RecordBroadcast {
	r := &RecordBroadcast{writers: writers}
	r.cond = sync.NewCond(r)
	return r
}

func main() {
	f, err := os.Create("cond.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := NewRecordBroadcast(f)
	r.Start()
	r.Prompt()
}

func (r *RecordBroadcast) Prompt() {
	for {

		fmt.Printf(":> ")
		var s string
		_, err := fmt.Scanf("%s", &s)
		if err != nil {
			return
		}

		r.Lock()
		r.buf = s
		r.Unlock()

		r.cond.Broadcast()
	}
}

func (r *RecordBroadcast) Start() error {
	f := func(w io.Writer) {
		for {
			r.Lock()
			r.cond.Wait()
			_, err := fmt.Fprintf(w, "%s\n command", r.buf)
			if err != nil {
				return
			}
			r.Unlock()
		}
	}
	for i := range r.writers {
		go f(r.writers[i])
	}
	return nil
}
