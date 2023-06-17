package main

import (
	"sync"
	"time"
)

const bufSize = 3

func main() {
	var wg sync.WaitGroup

	ch := make(chan string, bufSize)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1 * time.Second)
			ch <- "hello"
		}

	}()

	for i := 0; i < bufSize; i++ {
		//* selectはchannelのブロックを回避するための構文

		select {
		case m := <-ch:
			println(m)

		default:

			println("default")
		}

		time.Sleep(1500 * time.Millisecond)
	}

}
