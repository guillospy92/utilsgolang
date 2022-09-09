package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	http.HandleFunc("/log", logHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error boom", err)
	}
}

// https://hackernoon.com/go-the-complete-guide-to-profiling-your-code-h51r3waz

func logHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)

	go func() {
		obj := make(map[string]float64)

		if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
			ch <- http.StatusBadRequest
			close(ch)
			return
		}

		time.Sleep(time.Duration(rand.Intn(400)) * time.Millisecond)
		ch <- http.StatusOK
		close(ch)
	}()

	select {
	case result := <-ch:
		w.WriteHeader(result)
	case <-time.After(200 * time.Second):
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
