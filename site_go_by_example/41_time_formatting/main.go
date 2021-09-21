package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	t := time.Now()

	p("RFC:", t.Format(time.RFC3339))
	p("template 1:", t.Format("3:04PM"))
	p("template 2:", t.Format("Mon Jan _2 15:04:05 2006"))
	p("template 3:", t.Format("2006-01-02T15:04:05.999999-07:00"))
	fmt.Printf("Printf: %d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p("parse 1:", t1, e)

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p("parse 2:", t2)

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p("error:", e)
}
