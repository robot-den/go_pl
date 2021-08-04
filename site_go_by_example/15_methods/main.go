package main

import (
	"fmt"
	"math"
)

type Orb struct {
	Radius float64
}

func (o *Orb) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(o.Radius, 3)
}

func (o Orb) SurfaceArea() float64 {
	return 4.0 * math.Pi * math.Pow(o.Radius, 2)
}

func newOrb(radius float64) *Orb {
	return &Orb{
		Radius: radius,
	}
}

func main() {
	orbP := newOrb(10)
	fmt.Println("Volume:", orbP.Volume())
	orb := *orbP
	fmt.Println("Surface area:", orb.SurfaceArea())
}
