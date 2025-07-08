package main

import "fmt"

func nilPointer() {
	num := 15

	var ptr *int = &num
	var nilPtr *int
	fmt.Println("num", num)
	fmt.Println("ptr", ptr)
	fmt.Println("*ptr", *ptr)
	fmt.Println("nilPtr", nilPtr)
	if nilPtr != nil {
		fmt.Println("*nilPtr", *nilPtr)
	}
}

func foo(n *int) {
	*n++
}

func forDeferExamplePoint(n *int) {
	fmt.Println("defer Point", *n)
}

func forDeferExampleScalar(n int) {
	fmt.Println("defer Scalar", n)
}

func deferExample(pointer *int) {
	fmt.Println("function1", *pointer)
	defer forDeferExamplePoint(pointer)
	defer forDeferExampleScalar(*pointer)
	foo(pointer)
	fmt.Println("function2", *pointer)
}

func changeName(s *string) {
	*s = "Ivan"
}

func main() {
	number := 10
	var pointer *int = &number

	deferExample(pointer)

	name := "Petr"
	ptr := &name
	fmt.Println("Name", name)
	changeName(ptr)
	fmt.Println("changeName", name)
	nilPointer()
}
