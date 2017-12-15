package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args
	file, err := os.Open(filename[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
