package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	err := json.NewDecoder(r.Body).Decode(&d)
	fmt.Fprintln(w, "d.Name: ", d.Name)
	if err != nil {
		fmt.Fprintln(w, "Hello, World (err)!")
		fmt.Fprintln(w, err.Error())
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World (no Name parm)!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}

func main() {
	http.HandleFunc("/", HelloHTTP)
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
