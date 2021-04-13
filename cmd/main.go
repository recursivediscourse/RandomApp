package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("\n[Go project version 0.2]")
	fmt.Println("\nStarting server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is a test app")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})

	key := "d376008b03msh96a8f968c46ebb7p10a47ajsna2ea72779825"
	url := "https://brianiswu-cat-facts-v1.p.rapidapi.com/facts"
	client := &http.Client{}

	http.HandleFunc("/cats", func(w http.ResponseWriter, r *http.Request) {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("x-rapidapi-key", key)
		req.Header.Add("x-rapidapi-host", "brianiswu-cat-facts-v1.p.rapidapi.com")

		resp, err := client.Do(req)
		if err != nil {
			//handle
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error with reading body of request")
		} else {
			fmt.Fprintf(w, string(body))
			fmt.Println(string(body))
		}
	})

	// Add a page with a button to hit the cat API and then print the response.

	log.Fatal(http.ListenAndServe(":8081", nil))
}
