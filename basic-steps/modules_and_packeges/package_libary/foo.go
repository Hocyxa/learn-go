package package_libary

import "fmt"

func getTen() int {
	return 10
}

func SayHello() {
	fmt.Println("Hello")
	fmt.Println(getTen())
}
