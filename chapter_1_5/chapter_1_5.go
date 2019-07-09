package main

import (
	"fmt"
	// "io/ioutil"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	paths := os.Args[1:]

	for _, path := range paths {
		if !strings.HasPrefix(path, "http://") {
			path = "http://" + path
		}

		resp, err := http.Get(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch (get url): %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()

		b, err := io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch (read body): %s - %v\n", path, err)
			os.Exit(1)
		}

		fmt.Printf("\n%v\n", b)
		fmt.Printf("%s\n", resp.Status)
	}
}
