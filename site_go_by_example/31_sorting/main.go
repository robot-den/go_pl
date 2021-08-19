package main

import (
	"31_sorting/pkg/person"
	"log"
	"sort"
)

func main() {
	builtInSorts()
	customSort()
}

func builtInSorts() {
	strs := []string{"b", "x", "a", "n"}
	sort.Strings(strs)
	log.Println(strs)

	ints := []int{3, 5, 1, 7}
	sort.Ints(ints)
	log.Println(ints)

	log.Println(sort.IntsAreSorted(ints))
}

func customSort() {
	persons := person.Persons{
		{
			Name:    "Fedor",
			Will:    5,
			Courage: 1,
		},
		{
			Name:    "Denis",
			Will:    9,
			Courage: 10,
		},
		{
			Name:    "Lev",
			Will:    9,
			Courage: 2,
		},
	}
	log.Printf("Order before: %v", persons)
	sort.Sort(persons)
	log.Printf("Order after: %v", persons)
}
