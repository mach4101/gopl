package main

import (
	"WorkPlace/gopl/chapter2/popcount"
	"testing"
)

// 注释掉其他文件的main函数，然后在当前目录的命令行运行：go test -bench=.
func BenchmarkPopCountEx5(b *testing.B) {
	x := uint64(0x1234567890ABCDEF)
	for i := 0; i < b.N; i++ {
		popcount.PopCount(x)
	}
}

func BenchmarkRePopCountEx5(b *testing.B) {
	x := uint64(0x1234567890ABCDEF)
	for i := 0; i < b.N; i++ {
		popcount.RePopCountByBitOp(x)
	}
}
