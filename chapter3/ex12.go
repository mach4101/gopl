package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "sfsadjfalksflasdfk"
	s2 := "sdfjalfksadjfksz"

	if checkIfShuffle(s1, s2) && checkIfShuffle(s2, s1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func checkIfShuffle(s1, s2 string) bool {
	for _, c := range s1 {
		if strings.Index(s2, string(c)) == -1 {
			return false
		}
	}
	return true

}
