package main

import (
	"fmt"
	"unsafe"
)

func main ()  {
	// fmt.Println("Starting")
	var age int = 29
	fmt.Println("age is ", age)
	age = 44

	var width, height = 100, 100
	fmt.Println("width is ", width, " height is ", height)

	var (
		name = "ww"
		ee = 123
		rr int
	)
	fmt.Println("name is ", name, ee, rr)

	a1, a2, a3 := "a1", "a2", 3
	fmt.Println("a1 is ", a1, " a2 is ", a2, a3)

	// 类型
	a := false
	b := 5555

	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))

	c := 5.6666

	fmt.Printf("type of c is %T, size of c is %d", c, unsafe.Sizeof(c))

	c1 := complex(5, 7)
	c2 := 8 + 27i
	fmt.Println(c1 + c2)
	cmul := c1 * c2
	fmt.Println(cmul)

	const contant = 123
	// contant = 333

	var intVar int = contant
	var int32Var int32 = contant
	var float64Var float64 = contant
	var complex64Var complex64 = contant
	fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

	var adiv = 5.9/8
  fmt.Printf("a's type %T value %v",adiv, adiv)

	price, no := 5, 4
	total := calc(price, no)
	fmt.Println(total)

	area, perimeter := retPC(10.26, 6.66)
	
	area1, _ := retPC(10.88, 9.88)
	fmt.Println(area, area1, perimeter)
}

func calc(price, no int) int {
	var total = price * no
	return total
}

func retP(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2

	return area, perimeter
}
func retPC(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2

	return
}