package main

import (
	"fmt"
	"net/http"
)

func printHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List of HTTP Headers:")
	for key, values := range r.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
}

func main() {
	http.HandleFunc("/", printHeaders)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
