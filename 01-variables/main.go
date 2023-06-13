package variables

import "fmt"

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
}
