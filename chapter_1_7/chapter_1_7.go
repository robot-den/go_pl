package main

import (
	"fmt"
	"go_pl/shared/oscilloscope"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// http.HandleFunc("/", handler)
	// http.HandleFunc("/count", counter)
	http.HandleFunc("/", requestInfo)
	http.HandleFunc("/gif", gif)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func gif(w http.ResponseWriter, r *http.Request) {
	cycles, _ := strconv.Atoi(r.FormValue("cycles"))

	oscilloscope.Oscilloscope(w, float64(cycles))
}

func requestInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, %s %s\n\n", r.Method, r.URL, r.Proto)

	fmt.Fprint(w, "Headers:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "  %q => %q\n", k, v)
	}

	fmt.Fprintf(w, "\nRemote address: %s\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	fmt.Fprint(w, "Params:\n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "%q = %s\n", k, v)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
