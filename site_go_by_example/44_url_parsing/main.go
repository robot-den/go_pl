package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://mike:pswrd@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("protocol:", u.Scheme)

	fmt.Println("user:", u.User)
	fmt.Println("user, name:", u.User.Username())
	pass, _ := u.User.Password()
	fmt.Println("user, pass:", pass)

	fmt.Println("full host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("host:", host)
	fmt.Println("port:", port)

	fmt.Println("path:", u.Path)
	fmt.Println("fragment after `#`:", u.Fragment)

	fmt.Println("raw query:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("parsed query:", m)
	fmt.Printf("the `k` param value (%T): %v\n", m["k"], m["k"])
}
