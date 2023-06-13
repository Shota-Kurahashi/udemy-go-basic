package modules

import (
	"fmt"
	"go-basics/00-modules/calculator"
)

func Main() {
	fmt.Println(calculator.Offset)
	fmt.Println(calculator.Sum(1, 2))
	fmt.Println(calculator.Multiply(1, 2))
}
