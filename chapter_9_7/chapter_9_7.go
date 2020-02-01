package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"go_pl/shared/memo_hash"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"http://pikabu.ru",
		"http://google.com",
		"http://google.com",
		"http://google.com",
		"http://pikabu.ru",
		"http://google.com",
		"http://pikabu.ru",
		"http://google.com",
	}
	loadAllUrls(urls)
}

func getUrl(url string) (interface{}, error) {
	fmt.Println("Load url - ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func loadAllUrls(urls []string) {
	cacheHash := memo_hash.New(getUrl)

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			body, _ := cacheHash.Get(url)
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(body.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}
