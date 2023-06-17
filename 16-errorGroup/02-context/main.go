package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	s := []string{"fake1", "fake1", "task2", "fake2", "task3"}

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
	case "fake1":
		t = time.NewTicker(500 * time.Millisecond)

	case "fake2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(200 * time.Millisecond)
	}

	select {
	case <-context.Done():
		fmt.Printf("%v cancelled %v\n", task, context.Err())
		return context.Err()

	case <-t.C:
		t.Stop()
		if task == "fake1" || task == "fake2" {
			return fmt.Errorf("error: %v", task)
		}

		fmt.Printf("%v done\n", task)
	}
	return nil

}
