package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloRequest struct {
	Name string `json:"name"`
}

func main() {
	// Define the request body
	requestBody, err := json.Marshal(HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Send POST request to local HelloHTTP function
	resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
