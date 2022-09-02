package test

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// web server
var mu sync.Mutex
var count int

func Server() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s  %s  %s\n", r.Method, r.URL, r.Proto)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, "Count: ", count)
	mu.Unlock()
}
