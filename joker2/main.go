package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Making a GET request to the API endpoint
	response, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Displaying the random joke in JSON format
	fmt.Printf("Random Joke: %s\n", string(body))
}
