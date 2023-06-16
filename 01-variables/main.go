package variables

import (
	"fmt"
	"unsafe"
)

func basicVariables() {
	fmt.Println("\nbasicVariables")
	i := 1
	ui := uint16(2)
	fmt.Println(i)

	//? %T は型を表示する
	fmt.Printf("i: %v %T\n", i, i)

	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	f := 1.23456789
	s := "string"
	b := true
	fmt.Printf("f: %v %T s: %v %T b: %v %T\n", f, f, s, s, b, b)

	pi, title := 3.14, "pi"

	fmt.Printf("pi: %v %T title: %v %T\n", pi, pi, title, title)

	x := 10
	y := 1.23

	z := float64(x) + y

	fmt.Printf("z: %v %T\n", z, z)
}

func constVariables() {
	fmt.Println("\nconstVariables")
	const secret = "secret"
	type Os int

	//? iota は連番を生成する
	const (
		Mac Os = iota + 1
		Windows
		Linux
	)

	fmt.Printf("secret: %v %T\n", secret, secret)
	fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)

}

var (
	n int
	m string
	b bool
)

func Main() {
	basicVariables()
	constVariables()

	fmt.Printf("n: %v %T m: %v %T b: %v %T\n", n, n, m, m, b, b)

	var ui1 uint16
	fmt.Println(&ui1)

	var ui2 uint16

	fmt.Println(&ui2)

	var p1 *uint16

	fmt.Println(p1)

	p1 = &ui1

	fmt.Println(p1)
	fmt.Println(&p1)
	fmt.Println(*p1)
	fmt.Println(unsafe.Sizeof(p1))
	*p1 = 10
	fmt.Println(ui1)

	var pp1 **uint16 = &p1

	// &p1 は p1 のアドレスを返す
	// *pp1 は p1 のアドレスに格納されている値を返す
	fmt.Println(pp1)
	fmt.Println(&pp1)
	fmt.Println(unsafe.Sizeof(pp1))

	// *pp1 は p1 のアドレスが格納されているので、p1 のアドレスが返る

	fmt.Println(*pp1)
	fmt.Println(**pp1)

	ok, result := true, "A"
	if ok {
		// 同じスコープ内で result を再定義している
		// JavaScriptのletと同じ
		result := "B"
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)

	//? if の条件式で変数を宣言することができる

}
