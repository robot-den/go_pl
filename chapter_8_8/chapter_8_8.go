package main

import (
  "fmt"
  "io/ioutil"
  "path/filepath"
  "os"
  "flag"
  "sync"
)

func main() {
  flag.Parse()
  roots := flag.Args()
  if len(roots) == 0 {
    roots = []string{"."}
  }

  filesizes := make(chan int64)

  var wg sync.WaitGroup
  for _, root := range roots {
    wg.Add(1)
    go walkDir(&wg, root, filesizes)
  }

  go func () {
    wg.Wait()
    close(filesizes)
  }()

  var nfiles, nbytes int64
  for size := range filesizes {
    nfiles++
    nbytes += size
  }

  fmt.Printf("%d files take %d bytes\n", nfiles, nbytes)
}

func walkDir(wg *sync.WaitGroup, dir string, filesizes chan <-int64) {
  defer wg.Done()

  for _, entry := range dirents(dir) {
    if entry.IsDir() {
      wg.Add(1)
      subdir := filepath.Join(dir, entry.Name())
      go walkDir(wg, subdir, filesizes)
    } else {
      filesizes <- entry.Size()
    }
  }
}


var sema = make(chan struct{}, 40)
func dirents(dir string) []os.FileInfo {
  sema <- struct{}{}
  defer func () { <-sema }()

  entries, err := ioutil.ReadDir(dir)
  if err != nil {
    fmt.Fprintf(os.Stderr, "dir size: %v\n", err)
    return nil
  }
  return entries
}
