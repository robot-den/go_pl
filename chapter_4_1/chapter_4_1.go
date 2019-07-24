package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	str := []byte(os.Args[1])
	var flag string
	if len(os.Args) > 2 {
		flag = os.Args[2]
	}

	var arr [5]int
	brr := [5]int{1, 1}

	var crr [3]int = [...]int{1, 2, 3}
	const i = 3
	drr := [4]int{1: 9, 0: 8, i: 7}

	fmt.Println(arr)
	fmt.Println(brr)
	fmt.Println(crr)
	fmt.Println(drr)

	a := "X"
	b := "x"

	hA := sha256.Sum256([]byte(a))
	hB := sha256.Sum256([]byte(b))

	fmt.Println(hA)
	fmt.Println(hB)

	fmt.Println(hA == hB)

	counter := 0
	for i, byte := range hA {
		if byte != hB[i] {
			counter++
		}
	}

	fmt.Println(counter)

	fmt.Println("----------------")

	switch flag {
	case "SHA384":
		hash := sha512.Sum384(str)
		fmt.Printf("%x\n", hash)
	case "SHA512":
		hash := sha512.Sum512(str)
		fmt.Printf("%x\n", hash)
	default:
		hash := sha256.Sum256(str)
		fmt.Printf("%x\n", hash)
	}

}
