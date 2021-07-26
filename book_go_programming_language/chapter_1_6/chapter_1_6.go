package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	// "io"
	"io/ioutil"
	"net/url"
)

func main() {
	start := time.Now()
	channel := make(chan string)

	paths := os.Args[1:]

	for _, path := range paths {
		go fetch(path, channel)
	}

	for range paths {
		fmt.Println(<-channel)
	}

	fmt.Printf("Total %.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(path string, channel chan<- string) {
	start := time.Now()

	resp, err := http.Get(path)

	if err != nil {
		channel <- fmt.Sprint(err)
		return
	}

	// nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		channel <- fmt.Sprintf("while reading %s: %v", path, err)
		return
	}

	parsedUrl, _ := url.Parse(path)
	file, err := os.Create(fmt.Sprintf("./chapter_1_6/%s.html", parsedUrl.Hostname()))

	if err != nil {
		channel <- fmt.Sprintf("while file create %s: %v", path, err)
		return
	}

	nbytes, err := file.Write(body)

	if err != nil {
		channel <- fmt.Sprintf("while file write %s: %v", path, err)
		return
	}

	elapsed := time.Since(start).Seconds()
	channel <- fmt.Sprintf("%.2fs %7d %s", elapsed, nbytes, path)
}
