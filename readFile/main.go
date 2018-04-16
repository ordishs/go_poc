package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify a filename")
		os.Exit(1)
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("err", err)
		os.Exit(2)
	}

	io.Copy(os.Stdout, file)
	file.Close()
}
