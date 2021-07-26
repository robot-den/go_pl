package main

import (
  "fmt"
)

func main() {
  fmt.Println(quoteAsString(nil))
  fmt.Println(quoteAsString(777))
  fmt.Println(quoteAsString("y"))
  fmt.Println(quoteAsString(true))
  fmt.Println(quoteAsString(6.7))
}

func quoteAsString(x interface{}) string {
  switch x := x.(type) {
  case nil:
    return "NULL"
  case int:
    return fmt.Sprintf("%d", x)
  case string:
    return fmt.Sprintf("quoted %s", x)
  case bool:
    if x {
      return "TRUE"
    }
    return "FALSE"
  default:
    return "I dunno what you gave me, bro"
  }
}
