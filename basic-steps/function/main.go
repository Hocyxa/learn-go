package main

import (
	"fmt"
)

func procedure(str string) {
	if len(str) < 5 {
		fmt.Println("Light light baby", str)
		return
	} else {
		fmt.Println("HAAARD ", str)
	}
	fmt.Println("HAAARD ", str)
	fmt.Println("HAAARD ", str)

}
func Sum(num1 float64, num2 float64) float64 {
	return num1 + num2
}
func foo(a *int) {
	fmt.Println(a)
	*a += 10
}

func main() {
	procedure("2222")
	sum := Sum(112346, 2)
	procedure(fmt.Sprint(sum))
	a := 1
	fmt.Println(&a)
	fmt.Println(a)
	foo(&a)
	fmt.Println(&a)
	fmt.Println(a)

}
