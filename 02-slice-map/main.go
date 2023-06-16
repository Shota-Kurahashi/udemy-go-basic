package main

import "fmt"

func main() {
	var a1 [3]int
	var a2 = [3]int{1, 2, 3}
	a3 := [...]int{1, 2, 3}
	fmt.Printf("%v %v %v\n", a1, a2, a3)

	// len は配列の長さを返す
	// cap は配列の容量を返す

	fmt.Printf("len: %v %v %v\n", len(a1), len(a2), len(a3))
	fmt.Printf("cap: %v %v %v\n", cap(a1), cap(a2), cap(a3))

	// %T は型を表示する
	fmt.Printf("%T %T\n", a1, a2)

	// スライス

	var s1 []int

	s2 := []int{}

	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))

	// スライスの宣言方法によって、スライスの値が nil かどうかが異なる
	fmt.Printf("s1 == nil: %v\n", s1 == nil)
	fmt.Printf("s2 == nil: %v\n", s2 == nil)

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	s3 := []int{4, 5, 6}
	s1 = append(s1, s3...)

	// 配列... は配列の要素を展開して渡す JavaScript のスプレッド演算子と同じ
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))

	s4 := make([]int, 0, 2)
	// make([]int, 0, 2) は長さ 0 で容量 2 のスライスを作成する
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	s4 = append(s4, 1, 2, 3, 4)

	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	s5 := make([]int, 4, 6)
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	// [1:3] は 1 以上 3 未満の範囲を指定する
	s6 := s5[1:3]
	s6[1] = 10

	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	s6 = append(s6, 2)

	fmt.Printf("s6: %[1]T %[1]v %v %v\n", s6, len(s6), cap(s6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	//! スライスのコピーはメモリを共有する

	sc6 := make([]int, len(s5[1:3]))

	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	copy(sc6, s5[1:3])

	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	sc6[1] = 100

	fmt.Printf("sc6: %[1]T %[1]v %v %v\n", sc6, len(sc6), cap(sc6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	s5 = make([]int, 4, 6)

	// 最後の引数はメモリを共有する範囲を指定する
	// 3の場合は 3 未満の範囲を指定する
	fs6 := s5[1:3:3]

	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	fs6[0] = 6
	fs6[1] = 7

	fs6 = append(fs6, 8)

	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	s5[3] = 9

	fmt.Printf("fs6: %[1]T %[1]v %v %v\n", fs6, len(fs6), cap(fs6))
	fmt.Printf("s5: %[1]T %[1]v %v %v\n", s5, len(s5), cap(s5))

	var m1 map[string]int
	m2 := map[string]int{}
	// nilの違い
	fmt.Printf("m1: %v %v\n", m1, m1 == nil)
	fmt.Printf("m2: %v %v\n", m2, m2 == nil)

	m2["A"] = 1
	m2["B"] = 2
	m2["C"] = 3

	fmt.Printf("m2: %v %v %v\n", m2, len(m2), m2["A"])

	delete(m2, "A")

	fmt.Printf("m2: %v %v %v\n", m2, len(m2), m2["A"])

	// 範囲外の要素を指定すると 0 が返るそれの判定方法は以下の通り
	v, ok := m2["A"]
	fmt.Printf("v: %v %v\n", v, ok)

	for k, v := range m2 {
		//! for文の順番は保証されない
		fmt.Printf("%v: %v\n", k, v)
	}
}
