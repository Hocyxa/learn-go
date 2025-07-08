package main

import "fmt"

func foo() {
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		fmt.Println("defer2")
	}()
	defer func() {
		fmt.Println("defer3")
	}()
	fmt.Println("foo")
}

// TODO Посмотреть как дефер работает с слайсами
func main() {
	fmt.Println("Hello")
	foo()
	defer foo()
	fmt.Println("world")
}
