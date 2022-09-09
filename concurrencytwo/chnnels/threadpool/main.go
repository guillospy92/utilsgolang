package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

var r = regexp.MustCompile(`\((\d*),(\d*)\)`)

var wg = sync.WaitGroup{}

func main() {
	t := time.Now()
	absPath, _ := filepath.Abs("./thread")

	inputChannel := make(chan string, 1000)

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go findArea(inputChannel)
	}

	dat, err := ioutil.ReadFile(filepath.Join(absPath, "polygons.txt"))

	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(dat)

	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)

	wg.Wait()

	fmt.Println(time.Since(t).Seconds())
}

func findArea(inputChannel <-chan string) {
	var points []Point2D

	for input := range inputChannel {
		for _, p := range r.FindAllStringSubmatch(input, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}

		area := 0.0

		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}

		fmt.Println(math.Abs(area) / 2.0)
	}

	wg.Done()
}
