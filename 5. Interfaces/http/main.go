package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://www.google.com/")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999) // the read function does not resize the slice - so we just assume the no of element - not realistic though
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body) // alternative to line 21-23

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
