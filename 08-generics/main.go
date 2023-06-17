package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// * ~は型の制約を表す（独自の型を制約に含めることもできる）
type customConstraints interface {
	~int | int16 | float32 | float64 | string
}

type NewInt int

func add[T customConstraints](x, y T) T {
	return x + y
}

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func sumValues[K int | string, V constraints.Float | constraints.Integer](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("add(1, 2) = %d\n", add(1, 2))
	fmt.Printf("add(1.5, 2.5) = %f\n", add(1.5, 2.5))

	var i1, i2 NewInt = 1, 2

	fmt.Printf("add(i1, i2) = %d\n", add(i1, i2))

	fmt.Printf("min(1, 2) = %d\n", min(1, 2))

	m1 := map[string]uint{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	m2 := map[int]float32{
		1: 1.23,
		2: 4.56,
		3: 7.89,
	}

	fmt.Printf("sumValues(m1) = %d\n", sumValues(m1))
	fmt.Printf("sumValues(m2) = %f\n", sumValues(m2))

}
