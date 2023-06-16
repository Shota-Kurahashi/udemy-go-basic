package main

import (
	"fmt"
	"unsafe"
)

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{
		Title:    "Learn Go",
		Estimate: 10,
	}

	task1.Title = "Learn Golang"

	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	// 実態は別のものになる
	var task2 Task = task1
	task2.Title = "Learn Golang Again"

	fmt.Printf("%[1]T %+[1]v %v\n", task2, task2.Title)
	fmt.Printf("%[1]T %+[1]v %v\n", task1, task1.Title)

	task1p := &Task{
		Title:    "Learn Go Pointer",
		Estimate: 10,
	}

	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))

	task1p.Title = "Learn Golang Pointer"

	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))

	var task2p *Task = task1p

	task2p.Title = "Learn Golang Pointer Again"

	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	fmt.Printf("task2p: %T %+v %v\n", task2p, *task2p, unsafe.Sizeof(task2p))

	task1.extendEstimate()
	fmt.Printf("task1: %T %+v %v\n", task1, task1, unsafe.Sizeof(task1))

	task1.extendEstimatePointer()
	fmt.Printf("task1: %T %+v %v\n", task1, task1, unsafe.Sizeof(task1))
}

func (task Task) extendEstimate() {
	// コピーが生成されてそれに対して操作が行われる
	task.Estimate += 10
}

func (task *Task) extendEstimatePointer() {
	// 実態に対して操作が行われる
	task.Estimate += 10
}
