package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)

	s := []string{"fake1", "task2", "task3", "task4"}

	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println("done")

}

func doTask(context context.Context, task string) error {
	var t *time.Ticker

	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)

	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	case <-context.Done():
		fmt.Printf("%v cancelled %v\n", task, context.Err())
		return context.Err()

	case <-t.C:
		t.Stop()
		fmt.Printf("%v done\n", task)
	}
	return nil

}
