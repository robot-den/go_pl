package main

import (
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
)

func main() {
	sha1Example()
	base64Example()
}

func sha1Example() {
	fmt.Println("sha1:")
	s := "some string"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs)
}

func base64Example() {
	fmt.Println("\nbase64:")
	data := "abc123!?$*&()'-=@~"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
