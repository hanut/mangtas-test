package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Struct defining the type of Top Ten Response objects
// in the response array from the service
type WordCountResult struct {
	Word  string
	Count uint16
}

func main() {

	// Check if the arguments have been passed in the command line
	if len(os.Args) != 2 {
		log.Fatalln("Invalid number of arguments provided. \nUsage: go run main.go \"<input text>\"")
	}
	// Get the input text from the command line arguments
	inputText := os.Args[1]

	// Make a Post request with the inputText
	resp, err := http.Post("http://localhost:3000/", "text", bytes.NewBufferString(inputText))

	if err != nil {
		log.Fatal(err)
	}

	// Parse the response json
	var res []WordCountResult
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Fatalln(err)
	}

	// Print the results :)
	fmt.Println("Response from service:")
	for i, r := range res {
		fmt.Printf("%d. %s - %d\n", i+1, r.Word, r.Count)
	}
}
