package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// contextはgoroutineのキャンセルやタイムアウトを管理するためのもの
	// WithCancel, WithDeadline, WithTimeout, WithValueの4つの関数がある
	// それぞれの関数はcontext.Contextとcancel関数を返す
	// 引数は親のgoroutineのcontext.Context

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	wg.Add(3)
	go subTask(ctx, &wg, "a")
	go subTask(ctx, &wg, "b")
	go subTask(ctx, &wg, "c")

	wg.Wait()
	fmt.Println("done")
}

func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
	defer wg.Done()
	t := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-t.C:
			t.Stop()
			fmt.Println("subtask", id, "done")
			return
		}
	}

}
