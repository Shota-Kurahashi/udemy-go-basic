package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// 受信操作が行われるまでブロックされる
		ch <- 10
		time.Sleep(500 * time.Millisecond)
	}()

	fmt.Println(<-ch)
	wg.Wait()

	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()

	ch1 <- 10
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())

	// 最後の引数はバッファサイズ このサイズを超えるまではブロックしない

	ch2 := make(chan int, 1)

	ch2 <- 2
	// bufferが１なので、1つめが読み込まれるまでは実行できない
	ch2 <- 3
	fmt.Println(<-ch2)

}
