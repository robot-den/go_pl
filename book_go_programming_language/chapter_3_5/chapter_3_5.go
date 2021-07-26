package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "Hello, 开维"
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(len(str))

	fmt.Println(string(999))
	fmt.Println(string(12345678))

	fmt.Println(strings.ToUpper(str))
	fmt.Println(strconv.Itoa(999))

	fmt.Println(unicode.IsDigit(55))

	path := "a/b/c/da.e.f"

	fmt.Println(basename(path))
	fmt.Println(basename2(path))

	fmt.Println(withDots("87654321.44"))

	fmt.Println(withDotsBuffer("87654321.33"))

	fmt.Println(isAn("foko", "oofk"))

	fmt.Println(math.Pi)
}

func isAn(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	splittedS1 := strings.Split(s1, "")
	splittedS2 := strings.Split(s2, "")
	sort.Strings(splittedS1)
	sort.Strings(splittedS2)

	return strings.Join(splittedS1, "") == strings.Join(splittedS2, "")
}

func less(r1, r2 rune) bool {
	return r1 < r2
}

func withDotsBuffer(s string) string {
	floatPartIndex := strings.Index(s, ".")
	floatPart := ""

	if floatPartIndex >= 0 {
		floatPart = s[floatPartIndex:]
		s = s[:floatPartIndex]
	}

	runes := []rune(s)
	buffer := bytes.Buffer{}
	from := len(runes) % 3

	fmt.Println(from)

	for i, r := range runes {
		v := i - from

		if (v >= 0) && v%3 == 0 {
			buffer.WriteRune(',')
		}

		buffer.WriteRune(r)
	}

	return buffer.String() + floatPart
}

func reverseSlice(sl []rune) []rune {
	reversed := []rune{}

	for i := len(sl) - 1; i >= 0; i-- {
		reversed = append(reversed, sl[i])
	}

	return reversed
}

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func basename2(s string) string {
	lastSlashIndex := strings.LastIndex(s, "/") // -1 if "/" is not found
	s = s[lastSlashIndex+1:]
	if lastDotIndex := strings.Index(s, "."); lastDotIndex >= 0 {
		s = s[:lastDotIndex]
	}

	return s
}

func withDots(s string) string {
	floatPartIndex := strings.Index(s, ".")
	floatPart := ""

	if floatPartIndex >= 0 {
		floatPart = s[floatPartIndex:]
		s = s[:floatPartIndex]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return withDots(s[:n-3]) + "," + s[n-3:] + floatPart
}
