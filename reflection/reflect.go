package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
	Name string
}

type MyInt int

func main() {
	var f float64
	f = 35.2

	// reflect.TypeOf examples
	t1 := reflect.TypeOf(f)
	fmt.Println(t1.String())  // float64

	cat := Cat{
		Name: "Kitty",
	}
	t2 := reflect.TypeOf(cat)
	fmt.Println(t2.String()) // main.Cat

	// reflect.ValueOf examples
	v1 := reflect.ValueOf(f)
	fmt.Println(v1)  // 35.2
	fmt.Println(v1.String())  // <float64 Value>

	v2 := reflect.ValueOf(cat)
	fmt.Println(v2)  // {Kitty}
	fmt.Println(v2.String())  // <main.Cat Value>

	// reflect.TypeOf().Kind() -> Go has unlimited types, but kind() will find the similar type
	
	var i int
	var j MyInt
	fmt.Printf("%T\n", i) // equal to print ti, -> int
	fmt.Printf("%T\n", j) // main.MyInt

	fmt.Println("kind: ", reflect.TypeOf(i).Kind())  // int
	fmt.Println("kind: ", reflect.TypeOf(j).Kind())  // int
}