package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	ch1 <- 10
	close(ch1)

	v, ok := <-ch1
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)

	// すでにcloseされているので、0, falseが返る
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)

	ch3 := generateCountStream()
	for v := range ch3 {

		// 書込み後すぐに呼ばれる。
		fmt.Println(v)
	}

	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-nCh
			fmt.Printf("goroutine %d start\n", i)
		}(i)
	}

	time.Sleep(2 * time.Second)
	close(nCh)
	fmt.Println("unblocked by manual close")

	wg.Wait()
	fmt.Println("finished")
}

func generateCountStream() <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for i := 0; i < 5; i++ {
			channel <- i
		}
	}()

	return channel
}
