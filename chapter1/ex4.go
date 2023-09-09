package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filename := make(map[string][]string) // key: line，表示一行, value: filename，表示一行所对应的文件名，由于可能改行存在于多个文件中，所以使用slice
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, filename)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filename)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, filename[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filename map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// 如果没有添加过，那就添加，作为记录
		if !Contains(filename[input.Text()], f.Name()) {
			filename[input.Text()] = append(filename[input.Text()], f.Name())
		}
	}
}

// 避免重复添加到列表
func Contains(list []string, target string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}
