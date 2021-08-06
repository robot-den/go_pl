package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type customError struct {
	value int
}

func (c *customError) Error() string {
	return fmt.Sprintf("value is invalid: %d", c.value)
}

func main() {
	if result, err := func1(); err != nil {
		if customErr, ok := err.(*customError); ok {
			fmt.Println("This is a custom error:", customErr)
			return
		}
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func func1() (string, error) {
	if n := rand.Intn(10); n < 5 {
		return "", errors.New("can't work with ints less than 5")
	} else if n >= 8 {
		return "", &customError{value: n}
	} else {
		return fmt.Sprintf("value is valid: %d", n), nil
	}
}
