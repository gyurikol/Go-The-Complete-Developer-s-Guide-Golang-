package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// get program args
	var fn string = os.Args[1]

	// open file
	var f *os.File
	var err error
	f, err = os.Open(fn)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// handle file output
	io.Copy(os.Stdout, f)
}
