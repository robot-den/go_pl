package main

import (
  "fmt"
  "errors"
)

var messages = map[string]string{
  "1": "Error #1",
  "2": "Error #2",
  "3": "Error #3",
}

type Errstr string
func (e Errstr) Error() error  {
  return errors.New(messages[string(e)])
}

func main()  {
  err := Errstr("3")
  fmt.Println(err.Error())
}
