package main

import (
	"errors"
	"fmt"
	"image"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

// example run go run fileName ./img

const thumbnailNum = 5

type result struct {
	srcImgPath     string
	thumbnailImage *image.NRGBA
	error          error
}

func main() {
	// verify parameter
	if len(os.Args) < 2 {
		log.Fatal("need to send directory path of images")
	}

	start := time.Now()
	err := setupPipeLine(os.Args[1])

	if err != nil {
		log.Println("error", err)
	}

	fmt.Printf("Time Take %s\n", time.Since(start))
}

func setupPipeLine(root string) error {
	done := make(chan struct{})

	defer close(done)

	go func() {

		done <- struct{}{}
	}()

	// write channels
	pathsChanel, errChanel := walkFile(done, root)

	fmt.Println("no block")

	results := processImage(done, pathsChanel)

	fmt.Println("no block 2")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for r := range results {
			err := saveThumbnail(r.srcImgPath, r.thumbnailImage)

			if err != nil {
				fmt.Println("break error save img")
			}
		}
	}()

	wg.Wait()

	fmt.Println("no block 3")

	if err := <-errChanel; err != nil {
		return err
	}

	return nil
}

func processImage(done <-chan struct{}, paths <-chan string) <-chan *result {

	results := make(chan *result)

	// consumer
	thumbnail := func() {
		for path := range paths {
			srcImage, err := getPathSeparated(path)
			if err != nil {
				select {
				case results <- &result{srcImgPath: "", thumbnailImage: nil, error: err}:
				case <-done:
					fmt.Println(1111)
					return
				}
			} else {
				thumbnail := imaging.Thumbnail(srcImage, 100, 100, imaging.Lanczos)
				select {
				case results <- &result{srcImgPath: path, thumbnailImage: thumbnail}:
				case <-done:
					return
				}
			}
		}
	}

	// run num of routine producer
	var wg sync.WaitGroup

	wg.Add(thumbnailNum)
	// multiple producer
	for i := 0; i < thumbnailNum; i++ {
		go func() {
			thumbnail()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func walkFile(done <-chan struct{}, root string) (<-chan string, <-chan error) {

	paths := make(chan string)
	errChannel := make(chan error, 1)

	go func() {
		defer close(paths)
		defer close(errChannel)
		// list file in directory
		errChannel <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {

			// get error a list all file
			if err != nil {
				return err
			}

			// if file is regular
			if !info.Mode().IsRegular() {
				return nil
			}

			// check type img var path check all for example img/1.jpg
			contentType, err := getFileContentType(path)

			if err != nil {
				return err
			}

			if contentType != "image/jpeg" {
				return errors.New("image not compatible")
			}

			select {
			case paths <- path:
			case <-done:
				return fmt.Errorf("walk was cancelled")
			}

			return nil
		})
	}()

	return paths, errChannel
}

func saveThumbnail(srcImgPath string, thumbnailImage *image.NRGBA) error {

	// get part file last example img/28.jpg by 28.jpg
	fileName := filepath.Base(srcImgPath)
	dest := "thumbnail/thumbnail" + fileName

	err := imaging.Save(thumbnailImage, dest)

	if err != nil {
		return err
	}

	fmt.Printf("%s -> %s\n", srcImgPath, dest)

	return nil

}

func getFileContentType(file string) (string, error) {
	out, err := os.Open(file)

	if err != nil {
		return "", err
	}

	defer func(out *os.File) {
		err := out.Close()
		if err != nil {

		}
	}(out)

	buffer := make([]byte, 512)

	_, err = out.Read(buffer)

	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

func getPathSeparated(path string) (image.Image, error) {
	return imaging.Open(path)
}
