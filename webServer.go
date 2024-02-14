package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func home_response(response http.ResponseWriter, request *http.Request) {
	log.Printf("[INFO] Request From %s", request.URL.Path)

	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	data := map[string]interface{}{
		"Greatings": []string{"hello man You are cool", "Hello you Look Funny"},
	}

	json_data, err := json.Marshal(data)
	if err != nil {
		http.Error(response, "Json Error", 501)
		return
	}

	response.Header().Set("content-type", "application/json")
	response.Write([]byte(json_data))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_response)

	log.Println("Starting Server On :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
