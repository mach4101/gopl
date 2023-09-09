package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch, 1) // 第一次执行
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch, 2) // 第一次执行
	}

	for i := 0; i < (len(os.Args)-1)*2; i++ {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs, elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, times int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("times: %d %.2fs %7d %s", times, secs, nbytes, url)
}
