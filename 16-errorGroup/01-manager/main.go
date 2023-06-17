package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	s := []string{"task1", "fake1", "task2", "fake2", "task3"}

	for _, task := range s {
		task := task
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
	if task == "fake1" || task == "fake2" {

		return fmt.Errorf("%v failed", task)
	}

	fmt.Printf("task %v completed\n", task)

	return nil
}
