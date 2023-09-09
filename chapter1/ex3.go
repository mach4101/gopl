package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	sep = ""
	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i] + sep
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("low efficient version spent: ", time.Since(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println("strings.Join spent: ", time.Since(start))
}
