package chapter_2

import (
	"fmt"
)

type Cels float64
type Fahr float64

const (
	AbsoluteZeroC Cels = -273.15
	FreezingC     Cels = 0.0
	BoilingC      Cels = 100.0
)

func (c Cels) String() string { return fmt.Sprintf("%g℃", c) }
func (f Fahr) String() string { return fmt.Sprintf("%g℉", f) }
