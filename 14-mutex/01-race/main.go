package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int
	wg.Add(2)
	go func() {
		defer wg.Done()
		// 他のgoroutine中にロックされている場合はブロックされる
		mu.Lock()
		defer mu.Unlock()
		i = 1
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
		i = 2
	}()

	wg.Wait()

	// タイミングによっては同時にアクセスされるので1になる
	fmt.Printf("i = %d\n", i)

}
