package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaRec(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return commaRec(s[:n-3]) + "," + commaRec(s[n-3:])
}

func commaNoRec(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func commaFloatAndSymble(s string) string {
	var buf bytes.Buffer
	start := 0
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		start = 1
	}

	// n := len(s)

	index := strings.Index(s, ".")
	ints := s[start:index]
	floats := s[index:]

	return commaRec(ints) + floats
}

func main() {
	s := "11111112345.67890"
	fmt.Println(commaFloatAndSymble(s))
}
