package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Printf("error listener %v", err)
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}
		// response multiple client
		go handleCon(conn)

	}
}

func handleCon(c net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Println("error close connection")
		}
	}(c)
	for {
		_, err := io.WriteString(c, "response from server \n")
		if err != nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
}
