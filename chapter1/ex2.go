package main

import (
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		println(index, arg)
	}
}
