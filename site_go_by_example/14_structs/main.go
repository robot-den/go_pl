package main

import (
	"fmt"
	"math/rand"
	"time"
)

type cat struct {
	Name         string
	levelOfCrazy int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newCat(name string) *cat {
	c := cat{
		Name:         name,
		levelOfCrazy: rand.Intn(10),
	}
	return &c
}

func main() {
	fmt.Println(newCat("Felix"))
	fmt.Println(cat{"Keks", 2})
	fmt.Println(cat{Name: "Zero"})

	cotyan := newCat("Kotyan")
	fmt.Println(cotyan.Name, cotyan.levelOfCrazy)
	borisP := &cat{Name: "Boris"}
	borisP.levelOfCrazy = 7
	fmt.Println(borisP.Name, borisP.levelOfCrazy)
}
