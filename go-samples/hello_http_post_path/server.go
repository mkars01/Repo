package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gorilla/mux"
)

// HelloHTTP is an HTTP handler that reads a path parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	if name == "" {
		fmt.Fprint(w, "Hello, World (no Name parm)!")
		return
	}

	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(name))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name:[a-zA-Z0-9]+}", HelloHTTP).Methods("GET")

	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Server error:", err)
	}
}
