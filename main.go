package main

import (
	"strconv"
	"fmt"
	// "hello/prime/enum"
	// "hello/prime"
)

func main() {
	// println("hello world")
	// prime.PrintNumber(enum.Max)
	// defaultTypeVariable()
	// typeArray()
	// sliceArray()

	// variadic(1,3,5)
	// typeMap()
	// primetiveType()



	structJSON()
}

func defaultTypeVariable() {
	var a int
	var b bool
	var s string
	fmt.Printf("%T %d\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %q\n", s, s)
}

func typeArray() {
	var array = [...]int{1,3,5,7,9,11,13}

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(array[i])
	// }

	// LIKE for each  i=index, v=value
	// LIKE for each  _=ignore index, v=value
	for _, v := range array {
		fmt.Println(v)
	}
}

func sliceArray() {
	var slice []int
	
	if slice == nil {
		fmt.Println("it's nill")
	}

	// append function เติมของเข้า slice
	slice = append(slice, 1,3,5,7,9,11)

	// len = .lenght ใน java
	fmt.Println("length = ", len(slice))

	// for get by index
	if len(slice) > 3 {
		fmt.Println("access", slice[3])
	}

	for _, v := range slice[1:4] {
		fmt.Println(v)
	}

}


func variadic(x ...int) {
	for _, v := range x {
		fmt.Println(v)
	}
}

func typeMap() {
	m := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "coconut",
	}

	if m == nil {
		fmt.Println("it's nil")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	if c, ok := m["c"]; ok{
		fmt.Printf("%q\n", c)
	} else {
		fmt.Println("empty")
	}
}

// String type of string
type String string

func primetiveType() {
	var s String = "99"
	fmt.Printf("%T %v\n", s.toInt(), s.toInt())
}

func (s String) toInt() int {
	n, err := strconv.Atoi(string(s))
	if err != nil {
		fmt.Printf("Cannot convert %s to int\n", s)
	}
	return n
}


/* DAY 2 */

type Response struct {
	ID int64 `json: "id"`
	Name string `json: "firstName"`
	Title string `json: "title"`
}

var jsonString = `{
	"id" : 123,
	"firstName" : "Mos",
	"title" : "Mr."
}`

func structJSON() {
	var string int
	fmt.Println(string)

	// var response Response

}

type Rectangle struct {
	width float64
	length float64
}

func area(r Rectangle) float64 {
	return r.width * r.length
}

var rec Rectangle

func findArea() {
	
}



func closureFunc() {
	fn := series()

	fmt.Println(fn())
}

func series() (func() int, func()) {
	a, b := 0, 1

	return func() int {
		defer func() { a, b = b, a+b }()
		return a
	}
}
