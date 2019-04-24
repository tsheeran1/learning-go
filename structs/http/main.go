package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}
	len, err := io.Copy(lw, resp.Body)
	fmt.Println(len, err)

	// bs := make([]byte, 99999)
	// len, err := resp.Body.Read(bs)
	// fmt.Println(len, err)
	// fmt.Println(string(bs))

	// fmt.Println(resp.Header)
	// fmt.Println(resp.Body)
	// fmt.Println(resp.ContentLength)
	// fmt.Println(resp.TransferEncoding)
	// fmt.Println(resp.Close)
	// fmt.Println(resp.Uncompressed)
	// fmt.Println(resp.Trailer)
	// fmt.Println(resp.Request)
	// fmt.Println(resp.TLS)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote len:", len(bs))
	return len(bs), nil
}
