package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		defer close(ch1)
		ch1 <- "первый канал - 1"
		ch1 <- "первый канал - 2"
		ch1 <- "первый канал - 3"

	}()
	go func() {
		defer close(ch2)
		ch2 <- "второй канал - 1"
		ch2 <- "второй канал - 2"
		ch2 <- "второй канал - 3"
	}()
	go func() {
		defer close(ch3)
		ch3 <- "третий канал - 1"
		ch3 <- "третий канал - 2"
		ch3 <- "третий канал - 3"
	}()

	merge := mergeCh([]chan string{ch1, ch2, ch3})
	for v := range merge {
		fmt.Println(v)
	}
}

func mergeCh(chs []chan string) chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c chan string) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
