package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func sum(val1 int, val2 int) (int, error) {
	result := val1 + val2
	return result, nil
}

//SumHandler handles a sum
type SumHandler struct {
	User string
}

func (s SumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	res, err := sum(1, 2)

	if err != nil {
		log.Print("There was an error", err)
	}

	fmt.Fprintf(w, "Sum result was %d", res)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	sumHandler := &SumHandler{
		User: "otto",
	}
	http.Handle("/sum", sumHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
