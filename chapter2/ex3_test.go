package main

import (
	"WorkPlace/gopl/chapter2/popcount"
	"fmt"
	"testing"
	"time"
)

func main() {
	x := uint64(0x1234567890ABCDEF)
	start := time.Now()
	fmt.Println(popcount.PopCount(x))
	end := time.Since(start)

	fmt.Printf("elapsed: %v\n", end)

	start = time.Now()
	fmt.Println(popcount.RePopCount(x))
	end = time.Since(start)

	fmt.Printf("recount elapsed: %v\n", end)
}

// 注释掉main函数，然后在当前目录的命令行运行：go test -bench=.
func BenchmarkPopCount(b *testing.B) {
	x := uint64(0x1234567890ABCDEF)
	for i := 0; i < b.N; i++ {
		popcount.PopCount(x)
	}
}

func BenchmarkRePopCount(b *testing.B) {
	x := uint64(0x1234567890ABCDEF)
	for i := 0; i < b.N; i++ {
		popcount.RePopCount(x)
	}
}
