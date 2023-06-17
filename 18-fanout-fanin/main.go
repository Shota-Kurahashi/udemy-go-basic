package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// channelを分解して、複数のgoroutineに処理を分散させることをfan-outと呼ぶ
	// fan-outした結果を、channelを結合して、一つのchannelにまとめることをfan-inと呼ぶ
	// これにより、複数のgoroutineで並行処理を行い、結果を一つのchannelにまとめることができる
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	outChs := make([]<-chan string, cores)

	inData := generate(ctx, nums...)

	for i := 0; i < cores; i++ {
		outChs[i] = fanOut(ctx, inData, i+1)
	}

	var i int

	flag := true

	for v := range fanIn(ctx, outChs...) {
		if i == 3 {
			cancel()
			flag = false
		}

		if flag {
			fmt.Println(v)
		}
		i++
	}

	fmt.Println("done")
}

func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case <-ctx.Done():
				return
			case out <- n:
			}
		}
	}()

	return out
}

func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		heavyWork := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result:%v (id:%v)", i*i, id)
		}

		for v := range in {
			select {
			case <-ctx.Done():
				return

			case out <- heavyWork(v, id):
			}
		}
	}()

	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
