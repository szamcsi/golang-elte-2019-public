package main

import (
	"fmt"
	"net/http"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/hello1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello 1!")
	})

	m.HandleFunc("/hello2", hello2)

	sh := &specialHandler{}
	m.Handle("/hello3", sh)

	m.HandleFunc("/hello4", hello4("hello world"))

	m.Handle("/hello5/", hello5("hello 5"))

	http.ListenAndServe(":8080", m)
}

// ------------
func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hallo 2!")
}

// ------------
type specialHandler struct {
}

func (*specialHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Helloka 3!")
}

// ------------
func hello4(msg string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}
}

// ------------
func hello5(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
	}
}
