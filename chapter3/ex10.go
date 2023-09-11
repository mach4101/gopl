package main

import (
	"bytes"
	"fmt"
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

func main() {
	s := "1234567890"
	fmt.Println(commaNoRec(s))
}
