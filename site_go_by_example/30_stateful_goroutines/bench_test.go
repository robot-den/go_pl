package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
