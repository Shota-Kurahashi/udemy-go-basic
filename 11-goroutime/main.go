package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func tracer() {

	f, err := os.Create("trace.out")

	if err != nil {
		log.Fatalf("os.Create: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("f.Close: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("trace.Start: %v", err)
	}

	defer trace.Stop()

	ctx, t := trace.NewTask(context.Background(), "main")

	defer t.End()

	fmt.Println("Thee number of logical CPU Cores:", runtime.NumCPU())

	// task(ctx, "task1")
	// task(ctx, "task2")
	// task(ctx, "task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "task1")
	go cTask(ctx, &wg, "task2")
	go cTask(ctx, &wg, "task3")
	wg.Wait()

	s := []int{1, 2, 3}

	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main finished")
}

// func task(ctx context.Context, name string) {

// 	// Endのみ最後に行われる。
// 	defer trace.StartRegion(ctx, "task").End()
// 	time.Sleep(time.Second)
// 	fmt.Println((name))
// }

func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println(name)
}

func main() {

	var wg sync.WaitGroup
	// wg.Add(1) 1つのgoroutineを起動することを伝える
	wg.Add(1)
	go func() {
		// wg.Done() 1つのgoroutineが終了したことを伝える
		defer wg.Done()
		fmt.Println("Hello from goroutine")
	}()

	//* runtime.NumGoroutine() 起動しているgoroutineの数を返す 今回の場合はmainとgoroutineの2つ

	// wg.Wait() すべてのgoroutineが終了するまで待つ
	wg.Wait()
	fmt.Printf("num of working goroutines : %d\n", runtime.NumGoroutine())
	fmt.Println("main func finished")

	tracer()
}
