package main

import (
	"WorkPlace/gopl/chapter2/tempconv"
	"fmt"
)

func main() {
	var k tempconv.Kelvim
	k = 400

	fmt.Println(tempconv.KToC(k))
	fmt.Println(tempconv.KToF(k))
}
