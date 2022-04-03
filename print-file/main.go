package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// test to check if extra argument was parsed through
	if len(os.Args) < 1 {
		fmt.Println("No file specified!")
		os.Exit(1)
	}
	fileName := os.Args[1]

	// load file into byte slice
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, file)

	// // read content from byte slice
	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
