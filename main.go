package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func json_file(file_name string) []byte {
	file, err := os.Open("main.json") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var maxSize int = 4

	var byteChunk = make([]byte, maxSize)

	fmt.Println(file.Stat())

	// var rawData []byte

	for {
		// readTotal, err := file.Read(byteChunk)
		_, err := file.Read(byteChunk)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Print(byteChunk)
	}

	return nil
}

func main() {
	json_file("./main.json")
}
