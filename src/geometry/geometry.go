package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
	"unicode/utf8"
)

var _ = rectangle.Area

var rectlen, rectWidth float64 = 6, 7

func main()  {
	// var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometry shape properties")
	fmt.Printf(" area of rectangle %.2f\n", rectangle.Area(rectlen, rectWidth))
	fmt.Printf(" diagnal of rectangle %.2f\n", rectangle.Diagonal(rectlen, rectWidth))

	// 循环
	for i := 0; i < 10; i++ {
		// if i > 5 {
		// 	break
		// }
		if i % 2 == 0 {
			continue
		}
		fmt.Println("%d", i)
	}
	// i := 1
	// for i < 10 {
	// 	fmt.Printf("%d", i)
	// 	i += 2
	// }
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", i, no, no * i)
	}
	// for _, v := range v {
		
	// }
	finger := 8
	switch finger {
		case 1:
			fmt.Println("Thumb")
		case 2:
			fmt.Println("index")
		case 3:
			fmt.Println("Mid")
		case 4:
			fmt.Println("Ring")
		case 5:
			fmt.Println("Pinky")
		default: 
			fmt.Println("Invalid finger")
	}
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
			fmt.Println("vowel")
	default:
			fmt.Println("not a vowel")
	}

	switch num := number(); { // num is not a constant
	case num < 50:
			fmt.Printf("%d is lesser than 50\n", num)
			fallthrough
	case num < 100:
			fmt.Printf("%d is lesser than 100\n", num)
			fallthrough
	case num < 200:
			fmt.Printf("%d is lesser than 200", num)
	}

	arr := [3]int{1,2,3}
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 100
	}
	fmt.Println(arr)
	arrT()
	forRange()
	arrMult()
	slice()
	makeSlice()
	appendSlice()
	multSlice()
	copySlice()
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	nums := []int{89, 90, 95}
  find(89, nums...)

	runChange()

	createMap()
	printBytes()
	createString()
	createAdd()
	changeAddr()
}

func number() int {
	num := 15 * 5 
	return num
}

func arrT()  {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a // a copy of a is assigned to b
	b[0] = "Singapore"
	fmt.Println("a is ", a)
	fmt.Println("b is ", b) 
}

func forRange()  {
	a := [...]float64{0, 1, 2, 3, 4, 5}
	sum := float64(0)
	for i, v := range a {
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
}

func init()  {
	println("Initializing main program")
	if rectlen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
			for _, v2 := range v1 {
					fmt.Printf("%s ", v2)
			}
			fmt.Printf("\n")
	}
}

func arrMult()  {
	a := [3][2]string{
	{"lion", "tiger"},
	{"cat", "dog"},
	{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

func slice()  {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:3]
	// c := []int{1, 2, 3, 4}
	c := a[1:4]
	fmt.Println(b, c)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) // length of is 2 and capacity is 6
	fruitslice = fruitslice[:cap(fruitslice)] // re-slicing furitslice till its capacity
  fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
// make创建切片
func makeSlice()  {
	i := make([]int, 5, 5)
	fmt.Println(i)
}
// 切片增加
func appendSlice()  {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6

	var names []string //zero value of a slice is nil
	if names == nil {
			fmt.Println("slice is nil going to append")
			names = append(names, "John", "Sebastian", "Vinay")
			fmt.Println("names contents:",names, cap(names))
	}

	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food, cap(food))
}

func multSlice()  {
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
				fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}
// 内存优化
func copySlice() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

// 可变参数
func find(num int, nums ...int) {
	found := false
	for _, v := range nums {
		if v == num {
			found = true
		}
	}
	fmt.Println("found", found)
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func runChange()  {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

// map

func createMap()  {
	// var personSalary map[string]int
	// if personSalary == nil {
	// 	personSalary = make(map[string]int)
	// }
	personSalary := make(map[string]int)
	personSalary["ww"] = 15151515
	personSalary["tt"] = 121212
	
	customer := map[string]int {
		"ss": 1212121,
		"ssr": 1212121,
	}
	fmt.Println("personSalary", personSalary, customer)

	value, ok := customer["tt"]
	if ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("没找到")
	}
	for key, v := range personSalary {
		fmt.Println(key, v)
	}
	delete(personSalary, "tt")
	fmt.Println(personSalary, len(personSalary))
}


// stirng
func printBytes()  {
	s := "Señor"
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	// for index, rune := range s {
	// 	fmt.Printf("ss", rune, index)
	// }
	printCharsAndBytes(s)
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
			fmt.Printf("%c starts at byte %d\n", rune, index)
	}
	// for index, rune := range s {
	// 	fmt.Printf("ss", rune, index)
	// }
}

func createString()  {
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)

	fmt.Printf("length of %s is %d\n", str, utf8.RuneCountInString(str))
}

// 指针

func createAdd()  {
	b := 255
	var a *int = &b
	fmt.Println("address b is", a)
	fmt.Println("value b is", *a)
}

func changeAddr()  {
	a := 58
	b := &a
	fmt.Println(a)
	*b = 55
	fmt.Println(a)
}
// Go 不支持指针运算

// 结构体
