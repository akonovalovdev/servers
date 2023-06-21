package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"https://www.exzample.com",
		"https://www.google.com",
		"https://www.openai.com",
	}

	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error making the GET request to %s: %s\n", url, err.Error())
			continue
		}
		defer response.Body.Close()

		fmt.Printf("Response code for %s: %d\n", url, response.Header)
	}
}
