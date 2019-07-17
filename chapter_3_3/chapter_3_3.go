package main

import (
	"go_pl/shared/mandelbrot"
	"net/http"
)

func main() {
	http.HandleFunc("/", drawSvg)
	http.ListenAndServe("localhost:8000", nil)
}

func drawSvg(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "image/svg+xml")
	mandelbrot.Draw(w)
}
