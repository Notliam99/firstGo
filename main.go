package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func json_file(file_name string) (jsonData map[string]interface{}, err error) {
	/* var jsonData map[string]interface{} */

	file, err := os.Open("main.json") // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}

	var maxSize int = 4

	var byteChunk = make([]byte, maxSize)

	// fmt.Println(file.Stat())

	var rawData []byte

	for {
		// readTotal, err := file.Read(byteChunk)
		_, err = file.Read(byteChunk)
		if err != nil {
			if err != io.EOF {
				return
			}
			break
		}

		fmt.Print("ChunkStart\n")
		fmt.Print(byteChunk)
		fmt.Print("\nChunkEND\n")
		if string(byteChunk) != string("]") {
			rawData = append(rawData, byteChunk...)
		}
	}

	println(string(rawData))

	err = json.Unmarshal(rawData, &jsonData)
	if err != nil {
		log.Println("json")
		return
	}

	return
}

func main() {
	data, err := json_file("./main.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
