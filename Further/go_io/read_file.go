package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("read_file.go")
	if err != nil {
		log.Fatal(err)

	}

	data := make([]byte, 10)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
