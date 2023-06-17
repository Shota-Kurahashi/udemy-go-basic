package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var c int64

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				// 排他的にアクセスする つまり同時にアクセスされない
				// 第二引数は加算する値
				atomic.AddInt64(&c, 1)
			}
		}()
	}

	wg.Wait()
	println(c)
	fmt.Println("done")
}
