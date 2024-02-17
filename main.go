package main

import (
	"fmt"
	"github.com/Notliam99/firstGo/file"
	"log"
)

func main() {
	data, err := file.JsonFile("./main.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	for key, val := range data {
		fmt.Printf("%s: %s,\n", key, val)
	}
}
