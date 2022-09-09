package controllers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadNormal(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello")
	file, handler, err := request.FormFile("file")

	if err != nil {
		log.Println("error file form file")
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println("error close file")
			return
		}
	}(file)

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// create file
	dst, err := os.Create(handler.Filename)
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {

		}
	}(dst)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(writer, "Successfully Uploaded File\n")
	if err != nil {
		return
	}

}
