# tran-nam-golang-fresher-camp




package main

import (
	"fmt"
	"time"
)

func crawl(s int, ch chan int) {
	for i := range ch {

		ch <- i
		fmt.Printf("Crawling channel %v url %v\n", s, i)
	}
	close(ch)

}

func main() {
	ch := make(chan int, 1000)

	for i := 1; i <= 5; i++ {
		time.Sleep(10 * time.Millisecond)
		go crawl(i, ch)
	}

	for i := 1; i <= 1000; i++ {

		ch <- i
	}

}
