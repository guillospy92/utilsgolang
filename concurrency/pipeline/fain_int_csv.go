package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("error open file %v", err))
	}

	ch2, err := read("file2.csv")
	if err != nil {
		panic(fmt.Errorf("error open file %v", err))
	}

	for merge := range mergeChannel(ch1, ch2) {
		fmt.Println(merge)
	}

	fmt.Println("////////////////////////////////////////////////////// chanel synchronization")

	ch3, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("error open file %v", err))
	}

	ch4, err := read("file2.csv")
	if err != nil {
		panic(fmt.Errorf("error open file %v", err))
	}

	for merge := range mergeChannelSynchronization(ch3, ch4) {
		fmt.Println(merge)
	}
}

func read(file string) (<-chan []string, error) {
	ch1 := make(chan []string)

	// open file
	f, err := os.Open(file)

	// check error file open
	if err != nil {
		return nil, err
	}

	cr := csv.NewReader(f)

	// create new runtime
	go func() {
		for {
			// read rows of csv
			record, err := cr.Read()

			// terminate read file
			if err == io.EOF {
				close(ch1)
				break
			}

			ch1 <- record
		}
	}()

	return ch1, nil
}

func mergeChannelSynchronization(cs ...<-chan []string) <-chan []string {
	numChan := len(cs)

	wait := make(chan struct{}, numChan)

	out := make(chan []string)

	send := func(c <-chan []string) {

		for n := range c {
			out <- n
		}

		wait <- struct{}{}
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		for range wait {
			numChan--
			if numChan == 0 {
				break
			}
		}
		close(out)
	}()

	return out
}

func mergeChannel(cs ...<-chan []string) <-chan []string {
	var wg = new(sync.WaitGroup)

	out := make(chan []string)

	send := func(c <-chan []string, wg *sync.WaitGroup) {
		for n := range c {
			out <- n
		}
		defer wg.Done()
	}

	wg.Add(len(cs))

	// for all chan
	for _, forChanel := range cs {
		go send(forChanel, wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
