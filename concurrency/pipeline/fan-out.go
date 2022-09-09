package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	ch, err := readFanOut("file1-all.csv")

	if err != nil {
		panic("error boom")
	}

	br1 := breakup("1", ch)
	br2 := breakup("2", ch)
	br3 := breakup("3", ch)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		select {
		case _, ok := <-br1:
			if !ok {
				br1 = nil
			}
		case _, ok := <-br2:
			if !ok {
				br2 = nil
			}
		case _, ok := <-br3:
			if !ok {
				br3 = nil
			}
		}
	}

}

func readFanOut(file string) (<-chan []string, error) {
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

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})
	go func() {
		for v := range ch {
			fmt.Println(worker, v)
		}
		close(chE)
	}()

	return chE
}
