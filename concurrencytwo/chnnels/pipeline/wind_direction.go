package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	windRegex     = regexp.MustCompile(`\d* METAR.*EGLL \d*Z [A-Z ]*(\d{5}KT|VRB\d{2}KT).*=`)
	tafValidation = regexp.MustCompile(`.*TAF.*`)
	comment       = regexp.MustCompile(`\w*#.*`)
	metaClose     = regexp.MustCompile(`.*=`)
	variableWind  = regexp.MustCompile(`.*VRB\d{2}KT`)
	validWind     = regexp.MustCompile(`\d{5}KT`)
	windDirOnly   = regexp.MustCompile(`(\d{3})\d{2}KT`)
	windDist      [8]int
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()

	for {
		select {
		case c := <-ch:
			fmt.Println(c)
			break
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("delete cold")
		}
	}

	start := time.Now()
	absPath, err := filepath.Abs("./metarfiles")

	if err != nil {
		log.Printf("error recorring files %v", err)
		return
	}

	// get files
	files, err := ioutil.ReadDir(absPath)

	if err != nil {
		log.Printf("error files %v", err)
		return
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))

		if err != nil {
			log.Printf("error read file %v", err)
			return
		}

		text := string(data)
		_ = parseToArray(text)
		fmt.Println(file.Name())
		break
	}

	fmt.Println(time.Since(start).Seconds(), "seconds in run programs")
}

func parseToArray(text string) []string {
	lines := strings.Split(text, "\n")

	metaSlice := make([]string, 0, len(lines))
	metaStr := ""

	for _, line := range lines {
		if tafValidation.MatchString(line) {
			break
		}

		if !comment.MatchString(line) {
			metaStr += strings.Trim(line, " ")
		}

		if metaClose.MatchString(line) {
			metaSlice = append(metaSlice, metaStr)
			metaStr = ""
		}
	}

	return metaSlice
}
