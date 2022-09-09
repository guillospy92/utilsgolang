package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("allocating new bytes.buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, debug string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:04:45"))
	b.WriteString(" : ")
	b.WriteString(debug)
	b.WriteString("\n")
	

	_, err := w.Write(b.Bytes())
	if err != nil {
		return
	}

	bufPool.Put(b)
}

func main() {
	log(os.Stdout, "debug-string1")
	log(os.Stdout, "debug-string-2")
}
