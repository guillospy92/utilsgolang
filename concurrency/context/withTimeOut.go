package main

import (
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://andcloud.io", nil)
	ctx, cancel := context.WithTimeout(req.Context(), 500*time.Millisecond)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("ERROR", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error close body")
			return
		}
	}(resp.Body)

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		log.Println("ERROR COPY", err)
		return
	}
}
