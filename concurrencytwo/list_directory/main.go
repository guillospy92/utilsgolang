package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var matches []string

var waitGroup sync.WaitGroup

var mutex sync.Mutex

var count int32

func main() {
	start := time.Now()
	waitGroup.Add(1)
	go fileSearch("/home/guillospy", "README.md")
	waitGroup.Wait()
	for _, file := range matches {
		fmt.Println("Matched", file)
	}
	fmt.Println("count total fond", count)
	fmt.Printf("Total Time: %v", time.Since(start).Seconds())
}

func fileSearch(root string, filename string) {
	fmt.Println("searching in", root)

	files, _ := ioutil.ReadDir(root)

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			mutex.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))

			atomic.AddInt32(&count, 1)

			mutex.Unlock()
		}

		if file.IsDir() {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}

	waitGroup.Done()
}
