// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	bodyBytes, err2 := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	fmt.Fprintf(w, bodyString, err2)
	fmt.Fprintln(w, "Name Parameter: "+d.Name)
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World (err)!"+err.Error())
		log.Println("json.Compact:", err)
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World (no Name parm)!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
