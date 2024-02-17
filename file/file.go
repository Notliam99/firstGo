package file

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func JsonFile(file_name string) (jsonData map[string]interface{}, err error) {
	/* var jsonData map[string]interface{} */

	file, err := os.Open("main.json") // For read access.
	if err != nil {
		log.Fatal(err)
		return
	}

	rawData, err := io.ReadAll(file)

	err = json.Unmarshal(rawData, &jsonData)
	if err != nil {
		log.Println("json")
		return
	}

	return
}
