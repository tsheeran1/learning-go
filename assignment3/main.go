package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fileName := os.Args[1]
	fmt.Println(fileName)

	f, err := os.Open(fileName)
	if err != nil {
		println("Error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
