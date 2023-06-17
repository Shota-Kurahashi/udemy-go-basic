package main

import (
	"fmt"
)

func main() {
	a := -1

	if a == 0 {
		fmt.Println("a is 0")
	} else if a > 0 {
		fmt.Println("a is positive")
	} else {
		fmt.Println("a is negative")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println("infinite loop")
	// 	time.Sleep(1 * time.Second)
	// }

	// var i int

	// for {

	// 	if i > 3 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// 	i += 1
	// 	time.Sleep(1 * time.Second)
	// }

loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			continue
		case 3:
			continue
		case 8:
			break loop
		default:
			fmt.Println(i)
		}
	}

	items := []item{
		{price: 10},
		{price: 20},
		{price: 30},
	}

	for index := range items {
		//* 以下はコピーの値を変更しているだけなので、itemsの値は変わらない
		// item.price *= 1.1

		//* 以下はitemsの値を変更している
		items[index].price *= 1.1
	}

	fmt.Println(items)

}

type item struct {
	price float32
}
