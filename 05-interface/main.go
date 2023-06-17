package main

import (
	"fmt"
	"unsafe"
)

// メソッドを含んだら自動でinterfaceになる
type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed int
		enginePower int
}

type bycycle struct {
	speed int
	humanPower int
}
	

func (v *vehicle) speedUp() int{
	v.speed += 10 * v.enginePower

	return v.speed
}

func (v *vehicle) speedDown() int{
	v.speed -= 5 * v.enginePower

	return v.speed
}

func (b *bycycle) speedUp() int{
	b.speed += 3 * b.humanPower

	return b.speed
}

func (b *bycycle) speedDown() int{
	b.speed -= 1 * b.humanPower

	return b.speed
}

// 標準のString()メソッドをオーバーライド
func (v vehicle) String() string{
 return fmt.Sprintf("speed: %v, enginePower: %v", v.speed, v.enginePower)
}

func speedUpAndDown(c controller)  {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}


func main() {
	v := &vehicle{0,5}
	b := &bycycle{0,5}

	speedUpAndDown(v)
	speedUpAndDown(b)
	
	fmt.Println(v)

	// 以下の２つは同じ
	var i1 interface {}
	var i2 any
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i1)
	i2 = 1
	checkType(i2)
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
}

func checkType (i any){
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
		case nil:
		fmt.Println("nil")
	default:
		fmt.Println("unknown")
	}
}