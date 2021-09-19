package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

type Plants struct {
	XMLName xml.Name `xml:"plants"`
	Plants  []*Plant `xml:"parent>child>plant"`
}

func (p Plant) String() string {
	str := fmt.Sprintf("Plant id=%d, name=%s, origin=%v", p.Id, p.Name, p.Origin)
	return str
}

func main() {
	coffee := &Plant{Id: 1, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopio", "Brazil"}

	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out), "\n")

	fmt.Println(xml.Header+string(out), "\n")

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p, "\n")

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	plants := &Plants{}
	plants.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(plants, " ", "  ")
	fmt.Println(string(out))
}
